package kafkasarama

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
	"testing"
)

var addrs = []string{"127.0.0.1:9092"}

func TestSyncProducer(t *testing.T) {
	//创建一个 Sarama 的配置对象。
	cfg := sarama.NewConfig()
	//表示生产者要等待 Kafka 确认消息成功写入后再返回（同步模式）。如果不设置这个，SyncProducer.SendMessage 会一直失败。
	cfg.Producer.Return.Successes = true //同步的Producer一定要设置
	//创建一个同步的生产者实例
	producer, err := sarama.NewSyncProducer(addrs, cfg)
	assert.NoError(t, err)
	//构建消息并发送
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test_topic",
		//消息数据本体
		Value: sarama.StringEncoder("hello world ,这是一条使用kafka的消息"),
		//会在生产者和消费者之间传递，消息头，可传递自定义键值对，比如 trace_id 用于链路追踪。
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("trace_id"),
				Value: []byte("123456"),
			},
		},
		//只作用于发送过程。元信息，在发送过程中使用，可以用来传递额外信息，发送完成后会原样返回（不会传给消费者）。
		Metadata: "这是metadata",
	})

	// 打印消息
	dump.P(fmt.Sprintf("消息发送成功! 主题: %s, 分区: %d, 偏移: %d", "test_topic", partition, offset))
	assert.NoError(t, err)
}
