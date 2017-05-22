package dbkit

import (
	"github.com/go-redis/redis"
	"time"
	"strings"
	"fmt"
	"errors"
	"strconv"
)

var (
	redisClient *redis.Client
	inited bool
)

// 初始化Redis连接池
// @addr 	数据库地址
// @passad 	数据库密码,如果没有密码,填空
// @dbNum 	数据库名称,只能选择0~16之间
func InitRedis(addr string, passwd string, dbNum int) {
	if inited || redisClient != nil {
		return
	}
	if addr == ""{
		panic("redis addr is empty!")
	}
	if dbNum < 0 || dbNum > 16 {
		panic("redis dbNum is error!")
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:		addr,
		Password: 	passwd, //如果没有密码,默认为空
		DB: 		dbNum, //默认选择0数据库
		MaxRetries: 	3, //连接失败后重试3次
		DialTimeout: 	10 * time.Second, //拨号超时
		WriteTimeout: 	5 * time.Second, //写超时
		PoolSize: 	10, //最大连接数
		IdleTimeout: 	200 * time.Second,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis ping is error %v", err))
	}
	if strings.ToUpper(pong) != "PONG" {
		panic("redis conn return is not pong")
	}
	inited = true
}

// 向数据库中添加键值对内容,无过期时间
// @key 	主键
// @value 	内容
func RedisSet(key string, value string) error {
	return RedisSetWithExpire(key, value, 0)
}

// 向数据库中添加键值对内容,带过期时间
// @key 	主键
// @value 	内容
// @sec		过期时间,单位秒,0:永不过期
func RedisSetWithExpire(key string, value string, sec time.Duration) error {
	if key == "" || value == "" {
		return errors.New("redis set key or value is empty")
	}
	client := getRedisClient()
	if client == nil {
		return errors.New("redis client is nil")
	}
	err := client.Set(key, value, sec).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从数据库获取对应键内容
// @key 	主键
func RedisGet(key string) (string, error) {
	if key == "" {
		return "", errors.New("redis set key or value is empty")
	}
	client := getRedisClient()
	if client == nil {
		return "", errors.New("redis client is nil")
	}
	result, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("redis get key does not exists")
	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

// 获取Redis连接池状态
func RedisGetPoolStats() (map[string]string, error) {
	client := getRedisClient()
	if client == nil {
		return nil, errors.New("redis client is nil")
	}
	stats := make(map[string]string)
	poolStats := client.PoolStats()
	//共发起操作命令总数
	stats["Requests"] = strconv.FormatUint(uint64(poolStats.Requests), 10)
	//发起的所有操作命令中,从连接池拿到连接的总次数
	stats["Hits"] = strconv.FormatUint(uint64(poolStats.Hits), 10)
	//写操作超时总次数
	stats["Timeouts"] = strconv.FormatUint(uint64(poolStats.Timeouts), 10)
	//连接池中连接总数
	stats["TotalConns"] = strconv.FormatUint(uint64(poolStats.TotalConns), 10)
	//连接池中空闲连接总数
	stats["FreeConns"] = strconv.FormatUint(uint64(poolStats.FreeConns), 10)
	return stats, nil
}

func getRedisClient() *redis.Client {
	if !inited || redisClient == nil {
		return nil
	}
	return redisClient
}

//type RedisPool struct {
//	Host        string
//	Password    string
//	MaxIdle     int
//	IdleTimeOut time.Duration
//	pool        *redis.Pool
//	init        bool
//	once        sync.Once
//}
//
//func (this *RedisPool) New() *RedisPool {
//	this.init = true
//	if this.Host == "" {
//		panic("please provide host")
//	}
//	if this.MaxIdle <= 0 {
//		this.MaxIdle = 4
//	}
//	if this.IdleTimeOut <= 0 {
//		this.IdleTimeOut = 8 * time.Minute
//	}
//	this.once.Do(func() {
//		this.pool = &redis.Pool{
//			MaxIdle:     this.MaxIdle,
//			IdleTimeout: this.IdleTimeOut,
//			Dial: func() (redis.Conn, error) {
//				c, err := redis.Dial("tcp", this.Host)
//				if err != nil {
//					panic(err)
//				}
//				if this.Password != "" {
//					if _, err := c.Do("AUTH", this.Password); err != nil {
//						panic(err)
//					}
//				}
//
//				return c, nil
//			},
//			TestOnBorrow: func(c redis.Conn, t time.Time) error {
//				_, err := c.Do("PING")
//				return err
//			},
//		}
//	})
//	return this
//}
//func (this *RedisPool) GetRedis() redis.Conn {
//	if this.init == false {
//		panic("please call New before")
//	}
//	return this.pool.Get()
//}
