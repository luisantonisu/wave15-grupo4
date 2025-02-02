package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
		return model.Section{}, eh.GetErrInvalidData(eh.SECTION)
	}
	return h.rp.Create(section)
}

func (h *SectionService) Patch(id int, section model.SectionAttributesPtr) (model.Section, error) {
	return h.rp.Patch(id, section)
}

func (h *SectionService) Delete(id int) error {
	return h.rp.Delete(id)
}
