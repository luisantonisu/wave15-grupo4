package service

import (
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
