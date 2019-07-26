package model

import (
	"github.com/jinzhu/gorm"
)

// AddressEvent structure.
type AddressEvent struct {
	gorm.Model
	TxID       string
	Type       string
	Symbol     string
	Amount     string
	Info       string
	CostAmount string
	CostSymbol string
	AddressID  uint    `gorm:"type:bigint REFERENCES addresses(id)" json:"address_id"`
	Address    Address `json:"-"`
}

// NewAddressEvent return a AddressEvent instance.
func NewAddressEvent(addressID uint, eventType, symbol, amount, txid, info, costAmount, costSymbol string) *AddressEvent {
	return &AddressEvent{
		AddressID:  addressID,
		Amount:     amount,
		Type:       eventType,
		Symbol:     symbol,
		Info:       info,
		TxID:       txid,
		CostAmount: costAmount,
		CostSymbol: costSymbol,
	}
}
