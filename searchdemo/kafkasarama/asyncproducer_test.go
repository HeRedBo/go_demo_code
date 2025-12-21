package kafkasarama

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	//怎么知道发送是否成功
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(addrs, cfg)
	require.NoError(t, err)
	messages := producer.Input()
	go func() {
		for {
			messages <- &sarama.ProducerMessage{
				Topic: "test_topic",
				//分区依据
				Key: sarama.StringEncoder("user_123"), // 🔑 这里是分区依据
				//消息数据本体
				Value: sarama.StringEncoder("hello world ,这是一条使用kafka的消息"),
				//会在生产者和消费者之间传递
				Headers: []sarama.RecordHeader{
					{
						Key:   []byte("trace_id"),
						Value: []byte("123456"),
					},
				},
				//只作用于发送过程
				Metadata: "这是metadata",
			}
		}
	}()

	errCh := producer.Errors()
	succCh := producer.Successes()
	for {
		//两个都不满足就会阻塞
		select {
		case err := <-errCh:
			t.Log("发送出了问题", err.Err)
		case <-succCh:
			t.Log("发送成功")
		}
	}
}
