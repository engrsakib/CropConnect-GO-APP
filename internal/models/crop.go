package models

import (
    "time"

    "gorm.io/gorm"
)

type Crop struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    FarmerID    uint           `json:"farmer_id"`
    Name        string         `json:"name"`
    Category    string         `json:"category"` // vegetable, fruit, grain
    Description string         `json:"description"`
    ImageURL    string         `json:"image_url"`
    PricePerKg  float64        `json:"price_per_kg"`
    StockKg     float64        `json:"stock_kg"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
