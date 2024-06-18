package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

// NewClient 实例化客户端
func NewClient(config Config) (redis.UniversalClient, error) {

	// redis客户端实例
	var rdb redis.UniversalClient

	// 判断客户端类型
	switch config.Type {

	// 集群
	case "cluster":
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    config.Addrs,
			Password: config.Password,
		})

	// 单机
	default:
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Addrs[0],
			Password: config.Password,
			DB:       config.DB,
		})
	}

	// 测试redis客户端是否可用
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return rdb, err
	}

	return rdb, nil
}

// Client
// 获取指定的Redis客户端实例
func Client(name string) redis.UniversalClient {
	return clients[name]
}

// ClientAndErr
// 获取指定的Redis客户端实例不存在返回error
func ClientAndErr(name string) (redis.UniversalClient, error) {
	if client, ok := clients[name]; ok {
		return client, nil
	}
	return nil, errors.New("redis client not exists")
}
