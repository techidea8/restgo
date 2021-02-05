package redis

import (
	"time"

	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
)

func NewRedisPool(redisCfg Conf) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     redisCfg.MaxIdle,
		IdleTimeout: redisCfg.IdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(
				redisCfg.Addr,
				redis.DialDatabase(redisCfg.DB),
				redis.DialPassword(redisCfg.Password),
			)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return log.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
