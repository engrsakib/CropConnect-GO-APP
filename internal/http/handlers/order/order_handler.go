package orderhandler

import (
    "net/http"
    "strconv"

    orderdto "github.com/engrsakib/erp-system/internal/dto/order"
    orderservice "github.com/engrsakib/erp-system/internal/services/order"

    "github.com/gin-gonic/gin"
)

type OrderHandler struct {
    service orderservice.OrderService
}

func NewOrderHandler(s orderservice.OrderService) *OrderHandler {
    return &OrderHandler{service: s}
}

// @Summary Create Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body orderdto.CreateOrderRequest true "Order Data"
// @Success 201 {object} models.Order
// @Router /orders/ [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var req orderdto.CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    buyerID := c.GetUint("user_id")
    order, err := h.service.CreateOrder(buyerID, req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

// @Summary Get Order by ID
// @Tags Orders
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    order, err := h.service.GetOrderByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

// @Summary Get All Orders
// @Tags Orders
// @Success 200 {array} models.Order
// @Router /orders/ [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
    orders, err := h.service.GetAllOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}

// @Summary Update Order Status
// @Tags Orders
// @Param id path int true "Order ID"
// @Param status body orderdto.UpdateOrderStatusRequest true "Status"
// @Success 200 {object} models.Order
// @Router /orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    var req orderdto.UpdateOrderStatusRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order, err := h.service.UpdateOrderStatus(uint(id), req)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}
