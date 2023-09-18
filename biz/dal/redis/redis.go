package redis

import (
	"context"
	"sync"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	once   sync.Once
)

func GetClient(ctx context.Context) *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // redis地址
			Password: "",               // 密码
			DB:       0,                // 使用默认数据库
		})
	})
	return client
}
