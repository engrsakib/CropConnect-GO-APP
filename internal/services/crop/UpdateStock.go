package cropservice

import (
	"errors"
	dto "github.com/engrsakib/erp-system/internal/dto/crop"
	"github.com/engrsakib/erp-system/internal/models"
)

func (s *cropService) UpdateStock(id uint, req dto.UpdateStockRequest) (*models.Crop, error) {
	crop, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("crop not found")
	}

	crop.StockKg = req.StockKg
	return crop, s.repo.Update(crop)
}