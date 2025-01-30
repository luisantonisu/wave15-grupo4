package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
)

func NewSectionService(rp repository.ISection) *SectionService {
	return &SectionService{rp: rp}
}

type SectionService struct {
	rp repository.ISection
}

func (h *SectionService) GetAll() (map[int]model.Section, error) {
	return h.rp.GetAll()
}

func (h *SectionService) GetByID(id int) (model.Section, error) {
	return h.rp.GetByID(id)
}

func (h *SectionService) Create(section model.Section) (model.Section, error) {
	if section.SectionNumber <= 0 ||
		section.CurrentTemperature <= 0 ||
		section.MinimumTemperature <= 0 ||
		section.CurrentCapacity <= 0 ||
		section.MinimumCapacity <= 0 ||
		section.MaximumCapacity <= 0 ||
		section.WarehouseID <= 0 ||
		section.ProductTypeID <= 0 ||
		len(section.ProductBatchID) == 0 {
		return model.Section{}, errors.New("invalid section data")
	}
	return h.rp.Create(section)
}
