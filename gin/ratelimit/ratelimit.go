package ratelimit

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	rate "github.com/nazwa/go-rate"
)

type LimiterStruct struct {
	sync.Mutex
	list map[string]*rate.RateLimiter
}

func (l *LimiterStruct) Init() {
	l.list = make(map[string]*rate.RateLimiter)
}

func (l *LimiterStruct) Try(ip string, limit int, duration time.Duration) bool {
	l.Lock()
	defer l.Unlock()

	var timer *rate.RateLimiter
	var ok bool

	if timer, ok = l.list[ip]; !ok {
		timer = rate.New(limit, duration)
		l.list[ip] = timer
	}

	ok, _ = timer.Try()
	return ok
}

// This method checks if enpoints are logged in
func RateLimitMiddleware(limit int, duration time.Duration) gin.HandlerFunc {

	limiter := LimiterStruct{}
	limiter.Init()

	return func(c *gin.Context) {
		// Return 429 Too Many Requests is we are rate limited
		if !limiter.Try(c.ClientIP(), limit, duration) {
			c.AbortWithStatus(429)
			return
		}
	}
}
