package cropdto

type CreateCropRequest struct {
    Name        string  `json:"name" binding:"required"`
    Category    string  `json:"category" binding:"required"`
    Description string  `json:"description"`
    ImageURL    string  `json:"image_url"`
    PricePerKg  float64 `json:"price_per_kg" binding:"required"`
    StockKg     float64 `json:"stock_kg" binding:"required"`
}

type UpdateStockRequest struct {
    StockKg float64 `json:"stock_kg" binding:"required"`
}

type CropResponse struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Category    string  `json:"category"`
    Description string  `json:"description"`
    ImageURL    string  `json:"image_url"`
    PricePerKg  float64 `json:"price_per_kg"`
    StockKg     float64 `json:"stock_kg"`
    FarmerID    uint    `json:"farmer_id"`
}
