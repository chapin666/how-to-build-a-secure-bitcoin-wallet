package queue

type Queue interface {
	Listen()
	Execute(command Command) error
	ExecuteAndWait(command Command) (chan map[string]interface{}, error)
	SetTransactionHanlder(func(*Transaction) error)
	SetConfirmationHandler(func(*Transaction) error)
}

type Command struct {
	CommandTopic string                 `json:"-"`
	Command      string                 `json:"command"`
	Data         map[string]interface{} `json:"data"`
	Meta         map[string]interface{} `json:"meta"`
	ReplyTopic   string                 `json:"reply_topic,omitempty"`
}

type Transaction struct {
	Symbol string
	TxID   string
	Value  string
	To     string
}
