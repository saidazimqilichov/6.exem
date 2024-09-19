package middleware

import (
	"gateway/storage/redisdb"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	RedisClient *redisdb.RedisClient
	RateLimit   int
	Window      time.Duration
}

func (rl *RateLimiter) Limit(next http.HandlerFunc) http.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
