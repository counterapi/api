package config

import (
	"fmt"
	"os"
	"time"

	"github.com/counterapi/api/pkg"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
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
		DefaultCacheTime: pkg.DefaultCacheTime,
	}
}
