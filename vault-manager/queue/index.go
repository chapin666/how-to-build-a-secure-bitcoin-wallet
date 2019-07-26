package queue

import (
	"encoding/json"
	"how-to-build-a-secure-bitcoin-wallet/vault-manager/config"
	"how-to-build-a-secure-bitcoin-wallet/vault-manager/net"
	"log"

	"github.com/Shopify/sarama"
)

type queue struct {
	config             map[string]config.ChainConfig
	producer           net.KafkaProducer
	consumer           net.KafkaConsumer
	pending            map[string]chan map[string]interface{}
	transactionHandler func(*Transaction) error
	confirmTransaction func(*Transaction) error
}

func NewQueue(producer net.KafkaProducer, consumer net.KafkaConsumer, cfg map[string]config.ChainConfig) Queue {
	return &queue{
		producer: producer,
		consumer: consumer,
		config:   cfg,
		pending:  make(map[string]chan map[string]interface{}),
	}
}

func (q *queue) SetTransactionHanlder(transactionHandler func(*Transaction) error) {
	q.transactionHandler = transactionHandler
}

func (q *queue) ProcessTransaction(symbol string, msg *sarama.ConsumerMessage) {
	data := make(map[string]interface{})
	err := json.Unmarshal(msg.Value, &data)
	if err != nil {
		log.Println("Error parsing transaction", err, msg)
		return
	}
	transaction := &Transaction{
		Symbol: data["symbol"].(string),
		TxID:   data["txid"].(string),
		Value:  data["value"].(string),
		To:     data["to"].(string),
	}

	go q.transactionHandler(transaction)
}

func (q *queue) Listen() {

}

func (q *queue) Execute(command Command) error {
	return nil
}

func (q *queue) ExecuteAndWait(command Command) (chan map[string]interface{}, error) {
	return nil, nil
}

func (q *queue) SetConfirmationHandler(func(*Transaction) error) {
}
