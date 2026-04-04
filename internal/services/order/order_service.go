package orderservice

import (
    "errors"

    orderdto "github.com/engrsakib/erp-system/internal/dto/order"
    "github.com/engrsakib/erp-system/internal/models"
    croprepository "github.com/engrsakib/erp-system/internal/repository/crop"
    orderrepository "github.com/engrsakib/erp-system/internal/repository/order"
)

type OrderService interface {
    CreateOrder(buyerID uint, req orderdto.CreateOrderRequest) (*models.Order, error)
    GetOrderByID(id uint) (*models.Order, error)
    GetAllOrders() ([]models.Order, error)
    UpdateOrderStatus(id uint, req orderdto.UpdateOrderStatusRequest) (*models.Order, error)
}

type orderService struct {
    orderRepo orderrepository.OrderRepository
    cropRepo  croprepository.CropRepository
}

func NewOrderService(o orderrepository.OrderRepository, c croprepository.CropRepository) OrderService {
    return &orderService{o, c}
}

func (s *orderService) CreateOrder(buyerID uint, req orderdto.CreateOrderRequest) (*models.Order, error) {
    order := &models.Order{
        BuyerID: buyerID,
        Status:  "pending",
    }

    var total float64
    var farmerID uint

    for _, item := range req.Items {
        crop, err := s.cropRepo.FindByID(item.CropID)
        if err != nil {
            return nil, errors.New("crop not found")
        }

        if crop.StockKg < item.Quantity {
            return nil, errors.New("not enough stock")
        }

        subtotal := crop.PricePerKg * item.Quantity
        total += subtotal
        farmerID = crop.FarmerID

        order.Items = append(order.Items, models.OrderItem{
            CropID:    crop.ID,
            Quantity:  item.Quantity,
            UnitPrice: crop.PricePerKg,
            Subtotal:  subtotal,
        })

        crop.StockKg -= item.Quantity
        s.cropRepo.Update(crop)
    }

    order.TotalPrice = total
    order.FarmerID = farmerID

    err := s.orderRepo.Create(order)
    return order, err
}

func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
    return s.orderRepo.FindByID(id)
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
    return s.orderRepo.FindAll()
}

func (s *orderService) UpdateOrderStatus(id uint, req orderdto.UpdateOrderStatusRequest) (*models.Order, error) {
    order, err := s.orderRepo.FindByID(id)
    if err != nil {
        return nil, errors.New("order not found")
    }

    order.Status = req.Status
    return order, s.orderRepo.Update(order)
}
