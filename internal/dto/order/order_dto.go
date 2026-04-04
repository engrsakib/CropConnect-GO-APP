package orderdto

type OrderItemRequest struct {
    CropID   uint    `json:"crop_id" binding:"required"`
    Quantity float64 `json:"quantity" binding:"required"`
}

type CreateOrderRequest struct {
    Items []OrderItemRequest `json:"items" binding:"required"`
}

type UpdateOrderStatusRequest struct {
    Status string `json:"status" binding:"required"`
}
