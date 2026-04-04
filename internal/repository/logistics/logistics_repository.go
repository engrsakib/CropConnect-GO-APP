package logisticsrepository

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type LogisticsRepository interface {
    Create(log *models.Logistics) error
    Update(log *models.Logistics) error
    FindByOrderID(orderID uint) (*models.Logistics, error)
}

type logisticsRepository struct {
    db *gorm.DB
}

func NewLogisticsRepository(db *gorm.DB) LogisticsRepository {
    return &logisticsRepository{db}
}

func (r *logisticsRepository) Create(log *models.Logistics) error {
    return r.db.Create(log).Error
}

func (r *logisticsRepository) Update(log *models.Logistics) error {
    return r.db.Save(log).Error
}

func (r *logisticsRepository) FindByOrderID(orderID uint) (*models.Logistics, error) {
    var log models.Logistics
    err := r.db.Where("order_id = ?", orderID).First(&log).Error
    return &log, err
}
