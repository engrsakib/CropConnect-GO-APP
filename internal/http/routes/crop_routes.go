package routes

import (
    crophandler "github.com/engrsakib/erp-system/internal/http/handlers/crop"
    croprepository "github.com/engrsakib/erp-system/internal/repository/crop"
    cropservice "github.com/engrsakib/erp-system/internal/services/crop"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterCropRoutes(r *gin.RouterGroup, db *gorm.DB) {

    // Initialize repo → service → handler
    repo := croprepository.NewCropRepository(db)
    service := cropservice.NewCropService(repo)
    handler := crophandler.NewCropHandler(service)

    // Route group
    crop := r.Group("/crops")
    {
        crop.POST("/", handler.CreateCrop)
        crop.PUT("/:id", handler.UpdateCrop)
        crop.DELETE("/:id", handler.DeleteCrop)
        crop.GET("/", handler.GetAllCrops)
        crop.GET("/:id", handler.GetCropByID)
    }
}
