package dbkit

import (
	"fmt"
	"git.gumpcome.com/go_kit/logiccode"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

var (
	redisClient *redis.Client
	redisInited bool
)

// 初始化Redis连接池
// @addr 	数据库地址
// @passad 	数据库密码,如果没有密码,填空
// @dbNum 	数据库名称,只能选择0~16之间
func InitRedis(addr string, passwd string, dbNum int) {
	if redisInited || redisClient != nil {
		return
	}
	if addr == "" {
		panic("redis addr is empty!")
	}
	if dbNum < 0 || dbNum > 16 {
		panic("redis dbNum is error!")
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     passwd,           //如果没有密码,默认为空
		DB:           dbNum,            //默认选择0数据库
		MaxRetries:   3,                //连接失败后重试3次
		DialTimeout:  10 * time.Second, //拨号超时
		WriteTimeout: 5 * time.Second,  //写超时
		PoolSize:     10,               //最大连接数
		IdleTimeout:  200 * time.Second,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis ping is error %v", err))
	}
	if strings.ToUpper(pong) != "PONG" {
		panic("redis conn return is not pong")
	}
	redisInited = true
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
		return logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return logiccode.RedisClientErrorCode()
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
		return "", logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return "", logiccode.RedisClientErrorCode()
	}
	result, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", logiccode.RedisKeyErrorCode()
	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

// 向数据库中添加键值对内容,值是一组KV集合,无过期时间
// @key 	主键
// @fields 	内容
func RedisSetMap(key string, fields map[string]interface{}) error {
	return RedisSetMapWithExpire(key, fields, 0)
}

// 向数据库中添加键值对内容,,值是一组KV集合,带过期时间
// @key 	主键
// @fields 	内容
// @sec		过期时间,单位秒,0:永不过期
func RedisSetMapWithExpire(key string, fields map[string]interface{}, sec time.Duration) error {
	if key == "" || fields == nil || len(fields) == 0 {
		return logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return logiccode.RedisClientErrorCode()
	}
	err := client.HMSet(key, fields).Err()
	if err != nil {
		return err
	}
	if sec > 0 { //设置KEY过期时间
		client.Expire(key, sec)
	}
	return nil
}

// 从数据库获取对应键内容
// @key 	主键
func RedisGetMap(key string) (map[string]string, error) {
	result := make(map[string]string)
	if key == "" {
		return result, logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return result, logiccode.RedisClientErrorCode()
	}
	result, err := client.HGetAll(key).Result()

	if err == redis.Nil {
		return result, logiccode.RedisKeyErrorCode()
	} else if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

// 从数据库获取对应键内容
// @key 	主键
func RedisGetMapVal(key string,value ...string) ([]interface{}, error) {
	result := make([]interface{},0)
	if key == "" {
		return result, logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return result, logiccode.RedisClientErrorCode()
	}
	result, err := client.HMGet(key,value...).Result()

	if err == redis.Nil {
		return result, logiccode.RedisKeyErrorCode()
	} else if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

// 判断数据库是否存在该键
// @key 	主键
func RedisKeyExists(key string) (bool, error) {
	if key == "" {
		return false, logiccode.RedisParamsErrorCode()
	}
	client := getRedisClient()
	if client == nil {
		return false, logiccode.RedisClientErrorCode()
	}
	result, err := client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	if result == 0 {
		return false, nil
	}
	return true, nil
}

// 获取Redis连接池状态
func RedisGetPoolStats() (map[string]string, error) {
	client := getRedisClient()
	if client == nil {
		return nil, logiccode.RedisClientErrorCode()
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
	if !redisInited || redisClient == nil {
		return nil
	}
	return redisClient
}
