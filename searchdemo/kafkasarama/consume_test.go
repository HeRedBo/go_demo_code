package kafkasarama

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestConsumer(t *testing.T) {
	cfg := sarama.NewConfig()
	//正常来说，一个消费者都是归属一个消费者组的
	//消费者就是你的业务
	consumerGroup, err := sarama.NewConsumerGroup(addrs, "test_group", cfg)

	assert.NoError(t, err)
	err = consumerGroup.Consume(context.Background(), []string{"test_topic"}, testConsumerGroupHandler{})
	//你消费结束，就会到这里
	t.Log(err)
}

type testConsumerGroupHandler struct {
}

func (t testConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("Setup session:", session)
	dump.P("Setup session:", session)
	return nil
}

func (t testConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("Cleanup session:", session)
	return nil
}

func (t testConsumerGroupHandler) ConsumeClaim(
	//代表的是你和Kafka的会话（从建立连接到连接彻底断掉的那一段时间）
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	for msg := range msgs {
		//var bizMsg MyBizMsg
		//err := json.Unmarshal(msg.Value, &bizMsg)
		//if err != nil {
		//	//这就是消费消息出错
		//	//大多数时候就是重试
		//	//记录日志
		//	continue
		//}
		log.Println(string(msg.Value))
		session.MarkMessage(msg, "")
	}
	//什么情况下会到这里
	//msg被人关了，也就是要退出消费逻辑
	return nil
}

type MyBizMsg struct {
	Name string
}
