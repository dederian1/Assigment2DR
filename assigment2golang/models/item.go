// models/item.go
package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Item represents an item entity
type Item struct {
	ID          uint      `gorm:"primary_key" json:"item_id"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate performs validation on item fields
func (i *Item) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}
