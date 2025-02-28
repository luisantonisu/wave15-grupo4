package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryProduct "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	sectionRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewSectionService(sectionRepository sectionRepository.ISection, repositoryProduct repositoryProduct.IProduct, warehouseRepository warehouseRepository.IWarehouse) *SectionService {
	return &SectionService{
		sectionRp:   sectionRepository,
		productRp:   repositoryProduct,
		warehouseRp: warehouseRepository,
	}
}

type SectionService struct {
	sectionRp   sectionRepository.ISection
	productRp   repositoryProduct.IProduct
	warehouseRp warehouseRepository.IWarehouse
}

func (h *SectionService) GetAll() ([]model.Section, error) {
	return h.sectionRp.GetAll()
}

func (h *SectionService) GetByID(id int) (model.Section, error) {
	section, err := h.sectionRp.GetByID(id)
	if err != nil {
		return model.Section{}, err
	}
	return section, nil
}

func (h *SectionService) Create(section model.SectionAttributes) (model.Section, error) {
	_, err := h.warehouseRp.GetByID(*section.WarehouseID)
	if err != nil {
		return model.Section{}, eh.GetErrForeignKey(eh.WAREHOUSE)
	}

	_, err = h.productRp.GetProductByID(*section.ProductTypeID)
	if err != nil {
		return model.Section{}, eh.GetErrForeignKey(eh.PRODUCT)
	}

	newSection, err := h.sectionRp.Create(section)
	if err != nil {
		return model.Section{}, err
	}
	return newSection, nil
}

func (h *SectionService) Patch(id int, section model.SectionAttributes) (model.Section, error) {
	if section.WarehouseID != nil {
		_, err := h.warehouseRp.GetByID(*section.WarehouseID)
		if err != nil {
			return model.Section{}, eh.GetErrForeignKey(eh.WAREHOUSE)
		}
	}

	if section.ProductTypeID != nil {
		_, err := h.productRp.GetProductByID(*section.ProductTypeID)
		if err != nil {
			return model.Section{}, eh.GetErrForeignKey(eh.PRODUCT)
		}
	}

	updateSection, err := h.sectionRp.Patch(id, section)
	if err != nil {
		return model.Section{}, err
	}
	return updateSection, nil
}

func (h *SectionService) Delete(id int) error {
	err := h.sectionRp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (h *SectionService) Report(id *int) ([]model.ReportProductsBatches, error) {
	report, err := h.sectionRp.Report(id)
	if err != nil {
		return []model.ReportProductsBatches{}, err
	}
	return report, nil
}
