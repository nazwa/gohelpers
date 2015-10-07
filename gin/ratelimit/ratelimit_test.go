package ratelimit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimiting(t *testing.T) {

	limiter := LimiterStruct{}
	limiter.Init()

	count := 1
	duration := time.Second

	assert.True(t, limiter.Try("ip ip ip", count, duration))
	assert.True(t, limiter.Try("ip2 ip2 ip2", count, duration))
	assert.True(t, limiter.Try("ip3 ip3 ip3", count*2, duration))
	assert.False(t, limiter.Try("ip ip ip", count, duration))
	assert.False(t, limiter.Try("ip2 ip2 ip2", count, duration))
	assert.True(t, limiter.Try("ip3 ip3 ip3", count*2, duration))
	assert.False(t, limiter.Try("ip ip ip", count, duration))
	assert.False(t, limiter.Try("ip2 ip2 ip2", count, duration))
	assert.False(t, limiter.Try("ip3 ip3 ip3", count*2, duration))
	time.Sleep(duration)
	assert.True(t, limiter.Try("ip ip ip", count, duration))
	assert.True(t, limiter.Try("ip2 ip2 ip2", count, duration))
	assert.True(t, limiter.Try("ip3 ip3 ip3", count*2, duration))

}
