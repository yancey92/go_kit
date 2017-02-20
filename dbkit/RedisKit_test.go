package dbkit

import (
	"github.com/garyburd/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedisPool(t *testing.T) {
	r := (&RedisPool{Host: "127.0.0.1:6379"}).New().GetRedis()
	Convey("TestRedisPool", t, func() {
		reply, err := r.Do("ping")
		repl, _ := redis.String(reply, err)
		So(err, ShouldBeNil)
		So(repl, ShouldEqual, "PONG")
	})
}
