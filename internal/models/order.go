package models

import "time"

type Order struct {
    ID         uint       `gorm:"primaryKey" json:"id"`
    BuyerID    uint       `json:"buyer_id"`
    FarmerID   uint       `json:"farmer_id"`
    TotalPrice float64    `json:"total_price"`
    Status     string     `json:"status"` // pending, paid, shipped, delivered
    CreatedAt  time.Time  `json:"created_at"`
    UpdatedAt  time.Time  `json:"updated_at"`
    Items      []OrderItem `json:"items"`
}

type OrderItem struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    OrderID   uint    `json:"order_id"`
    CropID    uint    `json:"crop_id"`
    Quantity  float64 `json:"quantity"`
    UnitPrice float64 `json:"unit_price"`
    Subtotal  float64 `json:"subtotal"`
}
