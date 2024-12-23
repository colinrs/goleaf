package config

import (
	"github.com/colinrs/goleaf/internal/infra"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DataBase              *infra.DBConfig `json:"Database" yaml:"Database"`
	Redis                 redis.RedisConf `json:"Redis" yaml:"Redis"`
	Etcd                  EtcdConfig
	LocalCacheRefreshTime string // default 5m
}

type EtcdConfig struct {
	Endpoints []string `json:"Endpoints" yaml:"Endpoints"`
}
