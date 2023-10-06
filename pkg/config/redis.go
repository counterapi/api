package config

import (
	"fmt"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

// RedisCache is a config struct for redis cache.
type RedisCache struct {
	Store            *persist.RedisStore
	DefaultCacheTime time.Duration
}

// SetupRedisCache sets the redis up.
func SetupRedisCache() *RedisCache {
	return &RedisCache{
		Store: persist.NewRedisStore(redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr: fmt.Sprintf(
				"%s:%s",
				os.Getenv("REDIS_HOST"),
				os.Getenv("REDIS_PORT"),
			),
		})),
		DefaultCacheTime: 10 * time.Second,
	}
}
