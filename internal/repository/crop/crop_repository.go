package croprepository

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type CropRepository interface {
    Create(crop *models.Crop) error
    Update(crop *models.Crop) error
    FindByID(id uint) (*models.Crop, error)
    FindByCategory(category string) ([]models.Crop, error)
}

type cropRepository struct {
    db *gorm.DB
}

func NewCropRepository(db *gorm.DB) CropRepository {
    return &cropRepository{db}
}

func (r *cropRepository) Create(crop *models.Crop) error {
    return r.db.Create(crop).Error
}

func (r *cropRepository) Update(crop *models.Crop) error {
    return r.db.Save(crop).Error
}

func (r *cropRepository) FindByID(id uint) (*models.Crop, error) {
    var crop models.Crop
    err := r.db.First(&crop, id).Error
    return &crop, err
}

func (r *cropRepository) FindByCategory(category string) ([]models.Crop, error) {
    var crops []models.Crop
    err := r.db.Where("category = ?", category).Find(&crops).Error
    return crops, err
}
