package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Producer(m string) {

	brokerList := viper.GetStringSlice("kafka.brokerList")
	maxRetry := viper.GetInt("kafka.maxRetry")
	topic := viper.GetString("kafka.topic")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetry
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(m),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
