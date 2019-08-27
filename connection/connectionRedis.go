package connection

import (
	"github.com/gomodule/redigo/redis"
)

// NewRedis is connection for redis
func NewRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err)
			}
			return c, err
		},
	}
}
