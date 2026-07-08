package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type IPRateLimiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	r        rate.Limit
	b        int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		visitors: make(map[string]*visitor),
		r:        r,
		b:        b,
	}
	go i.cleanupVisitors()
	return i
}

func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, exists := i.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(i.r, i.b)
		i.visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

func (i *IPRateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		i.mu.Lock()
		for ip, v := range i.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(i.visitors, ip)
			}
		}
		i.mu.Unlock()
	}
}

func RateLimitMiddleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		l := limiter.getLimiter(ip)

		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, please try again later",
			})
			return
		}
		c.Next()
	}
}
