package routes

import (
    orderhandler "github.com/engrsakib/erp-system/internal/http/handlers/order"
    croprepository "github.com/engrsakib/erp-system/internal/repository/crop"
    orderrepository "github.com/engrsakib/erp-system/internal/repository/order"
    orderservice "github.com/engrsakib/erp-system/internal/services/order"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterOrderRoutes(r *gin.RouterGroup, db *gorm.DB) {
    orderRepo := orderrepository.NewOrderRepository(db)
    cropRepo := croprepository.NewCropRepository(db)
    orderService := orderservice.NewOrderService(orderRepo, cropRepo)
    orderHandler := orderhandler.NewOrderHandler(orderService)

    order := r.Group("/orders")
    {
        order.POST("/", orderHandler.CreateOrder)
        order.GET("/", orderHandler.GetAllOrders)
        order.GET("/:id", orderHandler.GetOrderByID)
        order.PUT("/:id/status", orderHandler.UpdateOrderStatus)
    }
}
