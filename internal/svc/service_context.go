package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/colinrs/goleaf/internal/config"
	"github.com/colinrs/goleaf/internal/infra"
	"github.com/colinrs/goleaf/pkg/client"
	"github.com/colinrs/goleaf/pkg/snowflake"
	"github.com/colinrs/goleaf/pkg/utils"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Redis
	EtcdClient  client.EtcdClient

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
