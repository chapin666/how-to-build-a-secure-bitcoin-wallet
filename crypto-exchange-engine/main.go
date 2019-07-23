package main

import (
	"crypto-exchange-engine/engine"
	"log"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func main() {

	// create the consumer and listen for new order messages
	consumer := createConsumer()

	// create the producer of trade message
	producer := createProducer()

	// create the order book
	book := engine.OrderBook{
		BuyOrders:  make([]engine.Order, 0, 100),
		SellOrders: make([]engine.Order, 0, 100),
	}

	// create a signal channel to know when we are done
	done := make(chan bool)

	// start processing orders
	go func() {
		for msg := range consumer.Messages() {
			var order engine.Order
			// decode the message
			order.FromJSON(msg.Value)
			// process the order
			trades := book.Process(order)
			// send trades to message queue
			for _, trade := range trades {
				rawTrade := trade.ToJSON()
				producer.Input() <- &sarama.ProducerMessage{
					Topic: "trades",
					Value: sarama.ByteEncoder(rawTrade),
				}
			}

			// mark the message as processed
			consumer.MarkOffset(msg, "")
		}
		done <- true
	}()

	// wait until we are done
	<-done
}

func createConsumer() *cluster.Consumer {
	// define our configuration to the cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Group.Return.Notifications = false

	// create the consumer
	consumer, err := cluster.NewConsumer([]string{"127.0.0.1:9092"}, "myconsumer", []string{"orders"}, config)
	if err != nil {
		log.Fatal("Unable to connect consumer to kafa cluster")
	}

	go handleErrors(consumer)
	go handleNotifictions(consumer)

	return consumer
}

func handleNotifictions(consumer *cluster.Consumer) {
	for ntf := range consumer.Notifications() {
		log.Printf("Rebalanced: %+v\n", ntf)
	}
}

func handleErrors(consumer *cluster.Consumer) {
	for err := range consumer.Errors() {
		log.Printf("Error: %s\n", err.Error())
	}
}

func createProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Fatal("Unable to connect producer to kafa server")
	}
	return producer
}
