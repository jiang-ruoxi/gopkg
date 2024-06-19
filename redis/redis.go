package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var defaultName string
var clients map[string]redis.UniversalClient
var clientDefault map[string]*redis.Client

type Config struct {
	Type     string   `json:"type" mapstructure:"type"` // cluster, failover,single-node , default is single-node
	Addrs    []string `json:"addrs" mapstructure:"addrs"`
	Password string   `json:"password" mapstructure:"password"`
	DB       int      `json:"db" mapstructure:"db"`
}

func InitFromViper() {
	defaultName = viper.GetString("redis.default")
	var cfg map[string]Config
	err := viper.UnmarshalKey("redis.clients", &cfg)
	if err != nil {
		panic(err)
	}
	clients = make(map[string]redis.UniversalClient)
	for k := range cfg {
		if clients[k], err = NewClient(cfg[k]); err != nil {
			panic(err)
		}
	}
}

func InitFromViperDefault() {
	defaultName = viper.GetString("redis.default")
	var cfg map[string]Config
	err := viper.UnmarshalKey("redis.clients", &cfg)
	if err != nil {
		panic(err)
	}
	clientDefault = make(map[string]*redis.Client)
	for k := range cfg {
		if clientDefault[k], err = NewClientDefault(cfg[k]); err != nil {
			panic(err)
		}
	}
}
