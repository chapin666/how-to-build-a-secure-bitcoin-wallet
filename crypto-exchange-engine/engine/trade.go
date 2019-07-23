package engine

import "encoding/json"

// Trade  - 交易.
type Trade struct {
	TakerOrderID string `json:"taker_order_id"`
	MakerOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        uint64 `json:"price"`
}

// FromJSON - json unmarshal
func (trade *Trade) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, trade)
}

// ToJSON - json marshal
func (trade *Trade) ToJSON() []byte {
	str, _ := json.Marshal(trade)
	return str
}
