package entity

import (
	"github.com/gofrs/uuid"
)

const (
	PaymentsTableName = "payments"
)

type Payment struct {
	PaymentID   uuid.UUID    `gorm:"type:uuid;primary_key" json:"paymentID"`
	PaymentDate string       `gorm:"type:date;not_null" json:"paymentDate"`
	Pay         Transactions `gorm:"ForeignKey:TransactionID" json:"supplierName"`
}

func NewPayment(paymentID uuid.UUID, paymentDate string, pay Transactions) *Payment {
	return &Payment{
		PaymentID:   paymentID,
		PaymentDate: paymentDate,
		Pay:         pay,
	}
}

func (model *Payment) TableName() string {
	return PaymentsTableName
}
