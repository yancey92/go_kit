package dbkit

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

type RedisPool struct {
	Host        string
	Password    string
	MaxIdle     int
	IdleTimeOut time.Duration
	pool        *redis.Pool
	init        bool
	once        sync.Once
}

func (this *RedisPool) New() *RedisPool {
	this.init = true
	if this.Host == "" {
		panic("please provide host")
	}
	if this.MaxIdle <= 0 {
		this.MaxIdle = 4
	}
	if this.IdleTimeOut <= 0 {
		this.IdleTimeOut = 8 * time.Minute
	}
	this.once.Do(func() {
		this.pool = &redis.Pool{
			MaxIdle:     this.MaxIdle,
			IdleTimeout: this.IdleTimeOut,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", this.Host)
				if err != nil {
					panic(err)
				}
				if this.Password != "" {
					if _, err := c.Do("AUTH", this.Password); err != nil {
						panic(err)
					}
				}

				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	})
	return this
}
func (this *RedisPool) GetRedis() redis.Conn {
	if this.init == false {
		panic("please call New before")
	}
	return this.pool.Get()
}
