package net

import "github.com/Shopify/sarama"

type KafkaProducer interface {
	Start() error
	Input() chan<- *sarama.ProducerMessage
	Errors() <-chan *sarama.ProducerError
	Close() error
}

type KafkaConsumer interface {
	Start() error
	GetMessageChan() <-chan *sarama.ConsumerMessage
	MarkOffset(msg *sarama.ConsumerMessage, meta string)
	ResetOffset(topic string, partition int32, offset int64, meta string) error
	Close() error
}
