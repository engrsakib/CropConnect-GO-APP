package cropservice

import (
	"errors"

	"github.com/engrsakib/erp-system/internal/models"

	dto "github.com/engrsakib/erp-system/internal/dto/crop"
)

func (s *cropService) UpdateCrop(id uint, req dto.CreateCropRequest) (*models.Crop, error) {
	crop, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("crop not found")
	}

	crop.Name = req.Name
	crop.Category = req.Category
	crop.Description = req.Description
	crop.ImageURL = req.ImageURL
	crop.PricePerKg = req.PricePerKg
	crop.StockKg = req.StockKg

	return crop, s.repo.Update(crop)
}