// models/order.go
package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Order represents an order entity
type Order struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CustomerName string     `json:"customer_name"`
	OrderedAt    *time.Time `json:"ordered_at"`
	Items        []Item     `json:"items"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// Validate performs validation on order fields
func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
