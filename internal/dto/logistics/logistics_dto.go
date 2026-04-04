package logisticsdto

type UpdateLogisticsStatusRequest struct {
    Status   string `json:"status" binding:"required"`
    Location string `json:"location" binding:"required"`
}
