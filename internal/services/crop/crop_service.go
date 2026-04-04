package cropservice

import (
	"errors"

	dto "github.com/engrsakib/erp-system/internal/dto/crop"
	"github.com/engrsakib/erp-system/internal/models"
	repo "github.com/engrsakib/erp-system/internal/repository/crop"
)

type CropService interface {
    CreateCrop(farmerID uint, req dto.CreateCropRequest) (*models.Crop, error)
    UpdateStock(cropID uint, req dto.UpdateStockRequest) (*models.Crop, error)
}

type cropService struct {
    repo repo.CropRepository
}

func NewCropService(r repo.CropRepository) CropService {
    return &cropService{r}
}

func (s *cropService) CreateCrop(farmerID uint, req dto.CreateCropRequest) (*models.Crop, error) {
    crop := &models.Crop{
        FarmerID:    farmerID,
        Name:        req.Name,
        Category:    req.Category,
        Description: req.Description,
        ImageURL:    req.ImageURL,
        PricePerKg:  req.PricePerKg,
        StockKg:     req.StockKg,
    }

    err := s.repo.Create(crop)
    return crop, err
}

func (s *cropService) UpdateStock(cropID uint, req dto.UpdateStockRequest) (*models.Crop, error) {
    crop, err := s.repo.FindByID(cropID)
    if err != nil {
        return nil, errors.New("crop not found")
    }

    crop.StockKg = req.StockKg
    err = s.repo.Update(crop)
    return crop, err
}
