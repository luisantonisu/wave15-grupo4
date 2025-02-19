package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	sectionRp "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
)

func NewSectionService(sectionRp sectionRp.ISection) *SectionService {
	return &SectionService{
		sectionRp: sectionRp,
	}
}

type SectionService struct {
	sectionRp sectionRp.ISection
	// productRepo repositoryProduct.IProduct
}

func (h *SectionService) GetAll() ([]model.Section, error) {
	allSections, err := h.sectionRp.GetAll()
	if err != nil {
		return nil, err
	}
	return allSections, nil
}

func (h *SectionService) GetByID(id int) (model.Section, error) {
	section, err := h.sectionRp.GetByID(id)
	if err != nil {
		return model.Section{}, err
	}
	return section, nil
}

func (h *SectionService) Create(section model.SectionAttributes) (model.Section, error) {
	newSection, err := h.sectionRp.Create(section)
	if err != nil {
		return model.Section{}, err
	}
	return newSection, nil
}

func (h *SectionService) Patch(id int, section model.SectionAttributes) (model.Section, error) {
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
		return nil, err
	}
	return report, nil
}
