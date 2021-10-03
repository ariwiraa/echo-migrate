package entity

import (
	"github.com/gofrs/uuid"
)

const (
	TransactionTableName = "Transactions"
)

type Transactions struct {
	TransactionID uuid.UUID `gorm:"type:uuid;primary_key" json:"transactionID"`
	Customers     Customers `gorm:"foreignKey:CustomerID" json:"customers"`
	Items         []Item    `gorm:"foreignKey:ItemID" json:"items"`
	Date          string    `gorm:"type:date;not_null" json:"date"`
	Qty           int       `gorm:"type:int;not_null" json:"qty"`
	TotalPrice    int       `gorm:"type:int;not_null" json:"totalPrice"`
}

func NewTransaction(transactionID uuid.UUID, customers Customers, items []Item, date string, qty int, totalPrice int) *Transactions {
	return &Transactions{
		TransactionID: transactionID,
		Customers:     customers,
		Items:         items,
		Date:          date,
		Qty:           qty,
		TotalPrice:    totalPrice,
	}
}

func (model *Transactions) TableName() string {
	return TransactionTableName
}
