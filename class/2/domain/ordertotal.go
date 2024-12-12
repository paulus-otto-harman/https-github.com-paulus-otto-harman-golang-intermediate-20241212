package domain

import "time"

type OrderTotal struct {
	ID              uint                `json:"id"`
	CustomerName    string              `json:"customer_name"`
	CustomerAddress string              `json:"customer_address"`
	PaymentMethod   string              `json:"payment_method"`
	Total           float32             `json:"total"`
	Status          string              `json:"status"`
	CreatedAt       time.Time           `json:"order_date"`
	Items           []OrderItemSubtotal `gorm:"foreignKey:OrderID" json:"items"`
}
