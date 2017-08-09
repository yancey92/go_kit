package dbkit

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)

//redis 命令文档 http://doc.redisfans.com/

func TestInitRedis(t *testing.T) {
	InitRedis("traderedisdev.redis.cache.chinacloudapi.cn:6379", "mOuUcyvHCUtvEkakSIqthQIoXQhUc8JDyHA12G/VzkM=", 0)
}

func TestRedisSetGetItem(t *testing.T) {
	err := RedisSet("key0", "key0value")
	if err != nil {
		t.Logf("插入失败 %v\n", err)
		t.Fail()
	}

	stats, _ := RedisGetPoolStats()
	//查看连接池状态
	t.Logf("连接池状态 %#v\n", stats)

	result, err := RedisGet("key0")
	if err != nil {
		t.Logf("读取失败%v\n", err)
		t.Fail()
	} else {
		t.Logf("读取成功 %s\n", result)
	}
}

func TestRedisSetEXItem(t *testing.T) {
	err := RedisSetWithExpire("key1", "key1value", 10 * time.Second)
	if err != nil {
		t.Logf("插入失败 %v\n", err)
		t.Fail()
	}

	t.Log("休息5秒")
	time.Sleep(5 * time.Second)

	stats, _ := RedisGetPoolStats()
	//查看连接池状态
	t.Logf("连接池状态 %#v\n", stats)

	result, err := RedisGet("key1")
	if err != nil {
		t.Logf("读取失败%v\n", err)
		t.Fail()
	} else {
		t.Logf("读取成功 %s\n", result)
	}

	t.Log("再休息7秒")
	time.Sleep(5 * time.Second)

	result, err = RedisGet("key1")
	if err != nil {
		t.Logf("读取失败%v\n", err)
		t.Fail()
	} else {
		t.Logf("读取成功 %s\n", result)
	}
}

func TestBeathRedisSetEXItem(t *testing.T) {
	InitRedis("127.0.0.1:6379", "", 0)

	for i := 0; i < 20; i++ {
		keyName := "key" + strconv.Itoa(i)
		keyValue := keyName + "value"
		go func() {
			t.Logf("插入 %s=%s\n", keyName, keyValue)
			RedisSet(keyName, keyValue)
		}()
	}

	t.Log("休息5秒")
	time.Sleep(5 * time.Second)

	stats, _ := RedisGetPoolStats()
	t.Logf("连接池状态 %#v\n", stats)

	for i := 0; i < 20; i++ {
		keyName := "key" + strconv.Itoa(i)
		result, _ := RedisGet(keyName)
		t.Logf("读取成功 %s=%s\n", keyName, result)
	}
}

func TestRedisExists(t *testing.T)  {
	InitRedis("traderedisdev.redis.cache.chinacloudapi.cn:6379", "mOuUcyvHCUtvEkakSIqthQIoXQhUc8JDyHA12G/VzkM=", 0)

	result, err := RedisKeyExists("key10")
	if err != nil {
		t.Fail()
	}
	t.Logf("判断结果 %v\n", result)
}

func TestAutoRetryConn(t *testing.T)  {
	InitRedis("127.0.0.1:6379", "", 0)
	result, _ := RedisKeyExists("key10")
	fmt.Printf("判断结果 %v\n", result)

	fmt.Println("休息10秒,关掉Redis")
	time.Sleep(10 * time.Second)

	result, err := RedisKeyExists("key10")
	if err != nil {
		fmt.Printf("断开连接 %v\n", err)
	}
	fmt.Printf("判断结果 %v\n", result)

	fmt.Println("再休息10秒,启动Redis")
	time.Sleep(10 * time.Second)

	result, err = RedisKeyExists("key10")
	if err != nil {
		fmt.Printf("断开连接 %v\n", err)
	}
	fmt.Printf("判断结果 %v\n", result)
}

func TestRedisSetMapWithExpire(t *testing.T) {
	InitRedis("127.0.0.1:6379", "", 0)
	testMap := make(map[string]interface{},0)
	testMap["name"] = "张三"
	testMap["age"] = "20"
	err := RedisSetMap("key0", testMap)
	if err != nil {
		t.Logf("插入失败 %v\n", err)
		t.Fail()
	}

	stats, _ := RedisGetPoolStats()
	//查看连接池状态
	t.Logf("连接池状态 %#v\n", stats)

	result, err := RedisGetMap("key0")
	if err != nil {
		t.Logf("读取失败%v\n", err)
		t.Fail()
	} else {
		t.Logf("读取成功result %s\n", result)
	}

	resultVal, err := RedisGetMapVal("key0","name","age")
	if err != nil {
		t.Logf("读取失败%v\n", err)
		t.Fail()
	} else {
		t.Logf("读取成功resultVal %s\n", resultVal)
	}
}