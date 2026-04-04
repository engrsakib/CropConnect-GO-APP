package models

import "time"

type Logistics struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    OrderID   uint      `json:"order_id"`
    Status    string    `json:"status"` // pending, shipped, out_for_delivery, delivered
    Location  string    `json:"location"`
    UpdatedAt time.Time `json:"updated_at"`
}
