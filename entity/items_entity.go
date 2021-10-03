package entity

import (
	"github.com/gofrs/uuid"
)

const (
	ItemsTableName = "items"
)

type Item struct {
	ItemID   uuid.UUID `gorm:"type:uuid;primary_key" json:"itemID"`
	ItemName string    `gorm:"type:varchar(200);not_null" json:"itemName"`
	Price    float64   `gorm:"type:float64;not_null" json:"price"`
	Supplier Supplier  `gorm:"foreignKey:SupplierID" json:"supplier"`
}

func NewItems(itemID uuid.UUID, itemName string, price float64, supplier Supplier) *Item {
	return &Item{
		ItemID:   itemID,
		ItemName: itemName,
		Price:    price,
		Supplier: supplier,
	}
}

func (model *Item) TableName() string {
	return ItemsTableName
}
