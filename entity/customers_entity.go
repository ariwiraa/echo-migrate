package entity

import (
	"github.com/gofrs/uuid"
)

const (
	CustomersTableName = "customers"
)

type Customers struct {
	CustomerID      uuid.UUID `gorm:"type:uuid;primary_key" json:"customerID"`
	CustomerName    string    `gorm:"type:varchar(200);not_null" json:"customerName"`
	CustomerAddress string    `gorm:"type:varchar(200);not_null" json:"customersName"`
}

func NewCustomers(customerID uuid.UUID, customerName, customerAddress string) *Customers {
	return &Customers{
		CustomerID:      customerID,
		CustomerName:    customerName,
		CustomerAddress: customerAddress,
	}
}

func (model *Customers) TableName() string {
	return CustomersTableName
}
