package entity

import (
	"github.com/gofrs/uuid"
)

const (
	SupplierTableName = "supplier"
)

type Supplier struct {
	SupplierID      uuid.UUID `gorm:"type:uuid;primary_key" json:"supplierID"`
	SupplierName    string    `gorm:"type:varchar(200);not_null" json:"supplierName"`
	SupplierAddress string    `gorm:"type:text;not_null" json:"supplierAddress"`
}

func NewSupplier(supplierID uuid.UUID, supplierName, supplierAddress string) *Supplier {
	return &Supplier{
		SupplierID:      supplierID,
		SupplierName:    supplierName,
		SupplierAddress: supplierAddress,
	}
}

func (model *Supplier) TableName() string {
	return SupplierTableName
}
