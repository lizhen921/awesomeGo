package kafka

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Consumer() {
	brokerList := viper.GetStringSlice("kafka.brokerList")
	topic := viper.GetString("kafka.topic")
	messageCountStart := viper.GetInt32("kafka.messageCountStart")

	brokerList = []string{"127.0.0.1:9092"}
	topic = "test_topics"
	messageCountStart = 0

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := brokerList
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)

	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				messageCountStart++
				fmt.Println("Received messages", msg.Offset, string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", messageCountStart, "messages")
}
