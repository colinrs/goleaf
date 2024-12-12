package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/colinrs/goleaf/internal/config"
	"github.com/colinrs/goleaf/internal/infra"
	"github.com/colinrs/goleaf/internal/model"
	"github.com/colinrs/goleaf/pkg/cache"
	"github.com/colinrs/goleaf/pkg/client"
	"github.com/colinrs/goleaf/pkg/codec"
	"github.com/colinrs/goleaf/pkg/snowflake"
	"github.com/colinrs/goleaf/pkg/utils"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"sync"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	RedisClient    *redis.Redis
	EtcdClient     client.EtcdClient
	LocalCache     cache.Cache
	LocalCacheLock sync.RWMutex

	snowflake snowflake.Snowflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config:      c,
		DB:          initDB(c),
		RedisClient: initRedis(c),
		EtcdClient:  initEtcdClient(c),
	}
	s.snowflake = s.initSnowflake(c)
	s.InitLocalCache()
	s.SyncLocalCache()
	return s
}

func initDB(c config.Config) *gorm.DB {
	db, err := infra.Database(c.DataBase)
	logx.Must(err)
	return db
}

func initRedis(c config.Config) *redis.Redis {
	return redis.MustNewRedis(c.Redis)
}

func initEtcdClient(c config.Config) client.EtcdClient {
	etcdClient, err := client.NewEtcdClient(c.Etcd.Endpoints)
	logx.Must(err)
	return etcdClient
}

func (s *ServiceContext) initSnowflake(c config.Config) snowflake.Snowflake {
	ipAddr := utils.GetOutboundIP()
	ipKey := fmt.Sprintf("%s/%s/%d", config.GetSnowflakeNodeIDKey(), ipAddr.String(), c.RestConf.Port)
	value, err := s.EtcdClient.Get(context.Background(), ipKey)
	if err != nil {
		logx.Error(err)
		value = 0
	}
	valueByte, ok := value.([]byte)
	if !ok {
		logx.Must(errors.New("invalid value type"))
	}
	intValue := cast.ToInt64(string(valueByte))
	nodeID := int64(0)
	if intValue > 0 {
		nodeID = intValue
		logx.Infof("initSnowflake get befor key:%s,nodeID: %d", ipKey, nodeID)
		return snowflake.NewSnowflake(nodeID)
	}
	nodeID, err = s.EtcdClient.Incr(context.Background(), config.GetSnowflakeNodeIDKey())
	logx.Must(err)
	err = s.EtcdClient.Set(context.Background(), ipKey, nodeID)
	logx.Must(err)
	logx.Infof("initSnowflake get new key:%s,nodeID: %d", ipKey, nodeID)
	return snowflake.NewSnowflake(nodeID)

}

func (s *ServiceContext) GetDB(ctx context.Context) *gorm.DB {
	return s.DB.WithContext(ctx)

}

func (s *ServiceContext) GetSnowflake() snowflake.Snowflake {
	return s.snowflake
}

func (s *ServiceContext) InitLocalCache() {
	memCache, err := newLocalCache()
	logx.Must(err)
	s.LocalCache = memCache
	return
}

func newLocalCache() (cache.Cache, error) {
	memCache, err := cache.NewRistrettoCache(cache.RistrettoCacheConfig{
		NumCounters: 100000,
		Capacity:    1000000,
		CostFunc:    func(value interface{}) int64 { return 1 },
	}, codec.NewSonicCodec())
	if err != nil {
		return nil, err
	}
	return memCache, nil
}

func (s *ServiceContext) SyncLocalCache() {
	var bizTags []*model.LeafAlloc
	db := s.DB.WithContext(context.Background())
	syncCache := func(bizTags []*model.LeafAlloc) {
		memCache, err := newLocalCache()
		if err != nil {
			logx.Errorf("newLocalCache err:%s", err.Error())
			return
		}
		count := 0
		_ = memCache.Flush(context.Background())
		for _, bizTag := range bizTags {
			err := memCache.Set(context.Background(), bizTag.BizTag.String, bizTag, 0)
			if err != nil {
				logx.Errorf("bizTag:%s,sync local cache err:%s", bizTag.BizTag.String, err.Error())
				continue
			}
			logx.Infof("bizTag:%s,sync local cache success", bizTag.BizTag.String)
			count++
		}
		oldCache := s.LocalCache
		s.LocalCacheLock.Lock()
		s.LocalCache = memCache
		s.LocalCacheLock.Unlock()
		_ = oldCache.Flush(context.Background())
		logx.Infof("sync local cache success count:%d, all:%d", count, len(bizTags))
	}
	err := db.Find(&bizTags).Error
	logx.Must(err)
	syncCache(bizTags)
	localCron := cron.New()
	entryID, err := localCron.AddFunc("@every 5m", func() {
		logx.Infof("sync local cache")
		db = s.DB.WithContext(context.Background())
		bizTags = []*model.LeafAlloc{}
		err = db.Find(&bizTags).Error
		if err != nil {
			logx.Errorf("sync local cache err:%s", err.Error())
			return
		}
		syncCache(bizTags)
	})
	logx.Must(err)
	localCron.Start()
	logx.Infof("entryID:%d", entryID)
	return
}

func (s *ServiceContext) GetLocalCache() cache.Cache {
	s.LocalCacheLock.RLock()
	localCache := s.LocalCache
	s.LocalCacheLock.RUnlock()
	return localCache
}
