package logisticsservice

import (
    "errors"
    logisticsdto "github.com/engrsakib/erp-system/internal/dto/logistics"
    "github.com/engrsakib/erp-system/internal/models"
    logisticsrepository "github.com/engrsakib/erp-system/internal/repository/logistics"
    orderrepository "github.com/engrsakib/erp-system/internal/repository/order"
)

type LogisticsService interface {
    UpdateStatus(orderID uint, req logisticsdto.UpdateLogisticsStatusRequest) (*models.Logistics, error)
    GetTracking(orderID uint) (*models.Logistics, error)
}

type logisticsService struct {
    logRepo   logisticsrepository.LogisticsRepository
    orderRepo orderrepository.OrderRepository
}

func NewLogisticsService(l logisticsrepository.LogisticsRepository, o orderrepository.OrderRepository) LogisticsService {
    return &logisticsService{l, o}
}

func (s *logisticsService) UpdateStatus(orderID uint, req logisticsdto.UpdateLogisticsStatusRequest) (*models.Logistics, error) {
    _, err := s.orderRepo.FindByID(orderID)
    if err != nil {
        return nil, errors.New("order not found")
    }

    log, err := s.logRepo.FindByOrderID(orderID)
    if err != nil {
        log = &models.Logistics{
            OrderID:  orderID,
            Status:   req.Status,
            Location: req.Location,
        }
        s.logRepo.Create(log)
    } else {
        log.Status = req.Status
        log.Location = req.Location
        s.logRepo.Update(log)
    }

    return log, nil
}

func (s *logisticsService) GetTracking(orderID uint) (*models.Logistics, error) {
    return s.logRepo.FindByOrderID(orderID)
}
