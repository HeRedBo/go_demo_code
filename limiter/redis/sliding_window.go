package redis

import (
	"context"
	"helloGO/limiter"

	"github.com/redis/go-redis/v9"

	"time"
)

type SlidingWindow struct {
	client redis.Cmdable
	cfg    limiter.Config
	script *redis.Script
}

func init() {
	limiter.RegisterRedis(limiter.SlidingWindow, func(cfg limiter.Config, client redis.Cmdable) limiter.Limiter {
		return NewSlidingWindow(client, cfg)
	})
}

func NewSlidingWindow(client redis.Cmdable, cfg limiter.Config) *SlidingWindow {
	script := redis.NewScript(`
		local key = KEYS[1]
		local limit = tonumber(ARGV[1])
		local window_ms = tonumber(ARGV[2])
		local now = tonumber(ARGV[3])
		local expire = tonumber(ARGV[4])

		redis.call("ZREMRANGEBYSCORE", key, 0, now - window_ms)
		local count = redis.call("ZCARD", key)

		if count >= limit then
			return 0
		end

		redis.call("ZADD", key, now, now)
		redis.call("EXPIRE", key, expire)
		return 1
	`)

	return &SlidingWindow{
		client: client,
		cfg:    cfg,
		script: script,
	}
}

func (r *SlidingWindow) Allow() (bool, error) {
	ctx := context.Background()
	now := time.Now().UnixMilli()

	resp, err := r.script.Run(ctx, r.client,
		[]string{r.cfg.Key},
		r.cfg.Rate,
		1000,
		now,
		int64(r.cfg.Expire.Seconds()), // 过期时间必须传！
	).Result()

	if err != nil {
		return false, err
	}
	return resp.(int64) == 1, nil
}
