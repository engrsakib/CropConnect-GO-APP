package cropservice

import ("github.com/engrsakib/erp-system/internal/models"

	dto "github.com/engrsakib/erp-system/internal/dto/crop"
)

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