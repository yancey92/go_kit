package nsqkit

import (
	"testing"
	"github.com/nsqio/go-nsq"
	"fmt"
	"sync"
)

func TestInitNsqClient(t *testing.T) {
	nsqClient := &NsqClient{}
	nsqConfig := &NsqConfig{}
	nsqClient.Init("123.206.79.254:4761", nsqConfig)
}

func TestSendMessage(t *testing.T) {
	nsqClient := &NsqClient{}
	nsqConfig := &NsqConfig{}
	nsqClient.Init("123.206.79.254:4761", nsqConfig)
	nsqClient.CreateProducer()
	for  {
		nsqClient.SendMessage("write_test", "test")
	}
}

func TestCreateConsumer(t *testing.T)  {
	count := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	nsqClient := &NsqClient{}
	nsqConfig := &NsqConfig{}
	nsqClient.Init("123.206.79.254:4761", nsqConfig)
	nsqClient.CreateConsumer("write_test", "test", func(message *nsq.Message) error {
		count++
		fmt.Printf("%v-%v\n", count, string(message.Body))
		return nil
	})
	wg.Wait()
}

func HandleJsonMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
