package crophandler

import (
	"net/http"
	"strconv"

	dto "github.com/engrsakib/erp-system/internal/dto/crop"
	service "github.com/engrsakib/erp-system/internal/services/crop"

	"github.com/gin-gonic/gin"
)

type CropHandler struct {
    service service.CropService
}

func NewCropHandler(s service.CropService) *CropHandler {
    return &CropHandler{service: s}
}

// CreateCrop godoc
// @Summary Create a new crop
// @Description Farmer can create a crop listing
// @Tags Crops
// @Accept json
// @Produce json
// @Param crop body dto.CreateCropRequest true "Crop Data"
// @Success 201 {object} models.Crop
// @Failure 400 {object} map[string]string
// @Router /crops/ [post]
func (h *CropHandler) CreateCrop(c *gin.Context) {
    var req dto.CreateCropRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    farmerID := c.GetUint("user_id")
    crop, err := h.service.CreateCrop(farmerID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, crop)
}


// UpdateCrop godoc
// @Summary Update crop information
// @Tags Crops
// @Accept json
// @Produce json
// @Param id path int true "Crop ID"
// @Param crop body dto.CreateCropRequest true "Crop Data"
// @Success 200 {object} models.Crop
// @Router /crops/{id} [put]
func (h *CropHandler) UpdateCrop(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    var req dto.CreateCropRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    crop, err := h.service.UpdateCrop(uint(id), req)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, crop)
}


// DeleteCrop godoc
// @Summary Delete a crop
// @Tags Crops
// @Param id path int true "Crop ID"
// @Success 200 {object} map[string]string
// @Router /crops/{id} [delete]
func (h *CropHandler) DeleteCrop(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    if err := h.service.DeleteCrop(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "crop deleted"})
}


// GetCropByID godoc
// @Summary Get crop by ID
// @Tags Crops
// @Param id path int true "Crop ID"
// @Success 200 {object} models.Crop
// @Router /crops/{id} [get]
func (h *CropHandler) GetCropByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    crop, err := h.service.GetCropByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "crop not found"})
        return
    }

    c.JSON(http.StatusOK, crop)
}


// GetAllCrops godoc
// @Summary Get all crops
// @Tags Crops
// @Success 200 {array} models.Crop
// @Router /crops/ [get]
func (h *CropHandler) GetAllCrops(c *gin.Context) {
    crops, err := h.service.GetAllCrops()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, crops)
}
