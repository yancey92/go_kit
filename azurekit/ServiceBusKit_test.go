package azurekit

import (
	"testing"
	"fmt"
	"time"
)

func TestSendMessageToQueue(t *testing.T)  {
	InitServiceBus("dev")

	uri := "https://dyltest.servicebus.chinacloudapi.cn/dyltest-queue"
	keyName := "RootManageSharedAccessKey"
	key := "iEZvPvwcSFe+U4gOFyZx7F4s2tIrznyNvMzn/4tpuAI="
	message := "你好, Azure ServiceBus"

	err := SendMessageToQueue(uri, keyName, key, message)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("消息发送到队列成功...")
}

func TestStartReceiveMessageFromQueueServer(t *testing.T)  {
	InitServiceBus("dev")

	uri := "https://dyltest.servicebus.chinacloudapi.cn/dyltest-queue"
	keyName := "RootManageSharedAccessKey"
	key := "iEZvPvwcSFe+U4gOFyZx7F4s2tIrznyNvMzn/4tpuAI="

	msgProcessor := func(msg string) bool {
		fmt.Println("开始处理消息")
		time.Sleep(2 * time.Second)
		fmt.Println("msgProcessor " + msg)
		fmt.Println("消息处理结束")
		return true
	}

	StartReceiveMessageFromQueueServer(uri, keyName, key, msgProcessor)
	time.Sleep(5 * time.Minute)
}