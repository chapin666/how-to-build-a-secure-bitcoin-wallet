package engine

import (
	"encoding/json"
)

// Order struct.
type Order struct {
	ID     string `json:"id"`
	Amount uint64 `json:"amount"`
	Price  uint64 `json:"price"`
	Side   int8   `json:"side"`
}

// FromJSON - json unmarshal
func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

// ToJSON - json marshal
func (order *Order) ToJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}
