package cropservice

import (
    dto "github.com/engrsakib/erp-system/internal/dto/crop"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/crop"
)

type CropService interface {
    CreateCrop(farmerID uint, req dto.CreateCropRequest) (*models.Crop, error)
    UpdateCrop(id uint, req dto.CreateCropRequest) (*models.Crop, error)
    UpdateStock(id uint, req dto.UpdateStockRequest) (*models.Crop, error)
    DeleteCrop(id uint) error
    GetCropByID(id uint) (*models.Crop, error)
    GetAllCrops() ([]models.Crop, error)
}

type cropService struct {
    repo repo.CropRepository
}

func NewCropService(r repo.CropRepository) CropService {
    return &cropService{repo: r}
}







func (s *cropService) DeleteCrop(id uint) error {
    return s.repo.Delete(id)
}

func (s *cropService) GetCropByID(id uint) (*models.Crop, error) {
    return s.repo.FindByID(id)
}

func (s *cropService) GetAllCrops() ([]models.Crop, error) {
    return s.repo.FindAll()
}
