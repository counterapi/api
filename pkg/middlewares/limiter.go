package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/counterapi/counterapi/pkg"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

const (
	limitInterval = time.Second
	limitCapacity = 10
)

// RateKeyFunc is a function for rate key.
type RateKeyFunc func(ctx *gin.Context) (string, error)

// RateLimiterMiddleware is middleware for Gin.
type RateLimiterMiddleware struct {
	fillInterval time.Duration
	capacity     int64
	ratekeygen   RateKeyFunc
	limiters     map[string]*ratelimit.Bucket
}

// get gets limiter bucket.
func (r *RateLimiterMiddleware) get(ctx *gin.Context) (*ratelimit.Bucket, error) {
	key, err := r.ratekeygen(ctx)
	if err != nil {
		return nil, err
	}

	if limiter, existed := r.limiters[key]; existed {
		return limiter, nil
	}

	limiter := ratelimit.NewBucketWithQuantum(r.fillInterval, r.capacity, r.capacity)
	r.limiters[key] = limiter

	return limiter, nil
}

// Middleware returns Gin middleware.
func (r *RateLimiterMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter, err := r.get(ctx)
		if err != nil || limiter.TakeAvailable(1) == 0 {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    "429",
				"message": fmt.Sprintf(pkg.ErrorMessageFormat, "too many requests"),
			})
		} else {
			ctx.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
			ctx.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
			ctx.Next()
		}
	}
}

// NewRateLimiter creates new middleware.
func NewRateLimiter(keyGen RateKeyFunc) *RateLimiterMiddleware {
	limiters := make(map[string]*ratelimit.Bucket)

	return &RateLimiterMiddleware{
		limitInterval,
		limitCapacity,
		keyGen,
		limiters,
	}
}
