package logisticshandler

import (
    "net/http"
    "strconv"

    logisticsdto "github.com/engrsakib/erp-system/internal/dto/logistics"
    logisticsservice "github.com/engrsakib/erp-system/internal/services/logistics"

    "github.com/gin-gonic/gin"
)

type LogisticsHandler struct {
    service logisticsservice.LogisticsService
}

func NewLogisticsHandler(s logisticsservice.LogisticsService) *LogisticsHandler {
    return &LogisticsHandler{service: s}
}

// @Summary Update logistics status
// @Tags Logistics
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param data body logisticsdto.UpdateLogisticsStatusRequest true "Status + Location"
// @Success 200 {object} models.Logistics
// @Router /logistics/{id}/status [put]
func (h *LogisticsHandler) UpdateStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    var req logisticsdto.UpdateLogisticsStatusRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log, err := h.service.UpdateStatus(uint(id), req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, log)
}

// @Summary Get logistics tracking
// @Tags Logistics
// @Param id path int true "Order ID"
// @Success 200 {object} models.Logistics
// @Router /logistics/{id} [get]
func (h *LogisticsHandler) GetTracking(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    log, err := h.service.GetTracking(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tracking not found"})
        return
    }

    c.JSON(http.StatusOK, log)
}
