package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/counterapi/api/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Default rate limit values
const (
	defaultLimitInterval = time.Minute
	defaultLimitCapacity = 30
	defaultKeyPrefix     = "ratelimit" // Removed the colon
)

// RateKeyFunc is a function for rate key.
type RateKeyFunc func(ctx *gin.Context) (string, error)

// RedisRateLimiter is middleware for rate limiting using Redis.
type RedisRateLimiter struct {
	fillInterval time.Duration
	capacity     int64
	ratekeygen   RateKeyFunc
	redisClient  *redis.Client
	keyPrefix    string
}

// Middleware returns Gin middleware.
func (r *RedisRateLimiter) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key, err := r.ratekeygen(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    "500",
				"message": fmt.Sprintf(pkg.ErrorMessageFormat, "rate limiter error"),
			})
			return
		}

		redisKey := r.keyPrefix + ":" + key

		// Use Redis INCR to count requests and EXPIRE to set the time window
		val, err := r.redisClient.Incr(context.Background(), redisKey).Result()
		if err != nil {
			fmt.Printf("Redis error: %v\n", err)
			// Let the request through if Redis fails
			ctx.Next()
			return
		}

		// If this is the first request in this window, set expiry
		if val == 1 {
			r.redisClient.Expire(context.Background(), redisKey, r.fillInterval)
		}

		// Check if rate limit is exceeded
		if val > r.capacity {
			fmt.Printf("Rate limit exceeded for key %s\n", key)
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    "429",
				"message": fmt.Sprintf(pkg.ErrorMessageFormat, "too many requests, please use v2 endpoints for higher rate limits"),
			})
			return
		}

		// Set rate limit headers
		remaining := r.capacity - val
		fmt.Printf("Request allowed for key %s, remaining=%d\n", key, remaining)
		ctx.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		ctx.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", r.capacity))
		ctx.Next()
	}
}

// NewRedisRateLimiter creates new Redis-based rate limiter middleware.
func NewRedisRateLimiter(keyGen RateKeyFunc) *RedisRateLimiter {
	// Get Redis connection details from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	// Get Redis DB index, default to 0 if not specified
	redisDB := 0
	redisDBStr := os.Getenv("REDIS_DB")
	if redisDBStr != "" {
		var err error
		redisDB, err = strconv.Atoi(redisDBStr)
		if err != nil {
			fmt.Printf("Warning: Invalid REDIS_DB value '%s', using default DB 0\n", redisDBStr)
			redisDB = 0
		}
	}

	// Get rate limit settings from environment variables
	limitInterval := defaultLimitInterval
	limitIntervalStr := os.Getenv("RATE_LIMIT_INTERVAL_SECONDS")
	if limitIntervalStr != "" {
		seconds, err := strconv.Atoi(limitIntervalStr)
		if err != nil {
			fmt.Printf("Warning: Invalid RATE_LIMIT_INTERVAL_SECONDS value '%s', using default %v\n",
				limitIntervalStr, defaultLimitInterval)
		} else {
			limitInterval = time.Duration(seconds) * time.Second
			fmt.Printf("Using rate limit interval of %v from environment\n", limitInterval)
		}
	}

	limitCapacity := int64(defaultLimitCapacity)
	limitCapacityStr := os.Getenv("RATE_LIMIT_CAPACITY")
	if limitCapacityStr != "" {
		capacity, err := strconv.ParseInt(limitCapacityStr, 10, 64)
		if err != nil {
			fmt.Printf("Warning: Invalid RATE_LIMIT_CAPACITY value '%s', using default %d\n",
				limitCapacityStr, defaultLimitCapacity)
		} else {
			limitCapacity = capacity
			fmt.Printf("Using rate limit capacity of %d from environment\n", limitCapacity)
		}
	}

	// Get Redis key prefix, default to "ratelimit:" if not specified
	keyPrefix := defaultKeyPrefix
	keyPrefixEnv := os.Getenv("RATE_LIMIT_KEY_PREFIX")
	if keyPrefixEnv != "" {
		keyPrefix = keyPrefixEnv
		fmt.Printf("Using rate limit key prefix '%s' from environment\n", keyPrefix)
	}

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	fmt.Printf("Connecting to Redis at %s (DB: %d) for rate limiting\n", redisAddr, redisDB)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   redisDB,
	})

	// Test the connection
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Warning: Could not connect to Redis: %v\n", err)
	}

	return &RedisRateLimiter{
		fillInterval: limitInterval,
		capacity:     limitCapacity,
		ratekeygen:   keyGen,
		redisClient:  redisClient,
		keyPrefix:    keyPrefix,
	}
}
