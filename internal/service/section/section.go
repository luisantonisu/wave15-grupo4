package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositorySection "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewSectionService(repositorySection repositorySection.ISection) *SectionService {
	return &SectionService{
		sectionRepo: repositorySection,
	}
}

type SectionService struct {
	sectionRepo repositorySection.ISection
	// productRepo repositoryProduct.IProduct
}

func (h *SectionService) GetAll() (map[int]model.Section, error) {
	return h.sectionRepo.GetAll()
}

func (h *SectionService) GetByID(id int) (model.Section, error) {
	return h.sectionRepo.GetByID(id)
}

func (h *SectionService) Create(section model.Section) (model.Section, error) {
	if section.SectionNumber <= 0 ||
		section.CurrentTemperature <= 0 ||
		section.MinimumTemperature <= 0 ||
		section.CurrentCapacity <= 0 ||
		section.MinimumCapacity <= 0 ||
		section.MaximumCapacity <= 0 ||
		section.WarehouseID <= 0 ||
		section.ProductTypeID <= 0 {
		return model.Section{}, eh.GetErrInvalidData(eh.SECTION)
	}

	// _, err := h.productRepo.GetProductByID(section.ProductTypeID)
	// if err != nil {
	// 	return model.Section{}, eh.GetErrNotFound(eh.PRODUCT)
	// }

	return h.sectionRepo.Create(section.SectionAttributes)
}

func (h *SectionService) Patch(id int, section model.SectionAttributesPtr) (model.Section, error) {
	return h.sectionRepo.Patch(id, section)
}

func (h *SectionService) Delete(id int) error {
	return h.sectionRepo.Delete(id)
}
