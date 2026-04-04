package routes

import (
    logisticshandler "github.com/engrsakib/erp-system/internal/http/handlers/logistics"
    logisticsrepository "github.com/engrsakib/erp-system/internal/repository/logistics"
    orderrepository "github.com/engrsakib/erp-system/internal/repository/order"
    logisticsservice "github.com/engrsakib/erp-system/internal/services/logistics"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterLogisticsRoutes(r *gin.RouterGroup, db *gorm.DB) {
    logRepo := logisticsrepository.NewLogisticsRepository(db)
    orderRepo := orderrepository.NewOrderRepository(db)
    service := logisticsservice.NewLogisticsService(logRepo, orderRepo)
    handler := logisticshandler.NewLogisticsHandler(service)

    log := r.Group("/logistics")
    {
        log.PUT("/:id/status", handler.UpdateStatus)
        log.GET("/:id", handler.GetTracking)
    }
}
