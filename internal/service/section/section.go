package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
)

func NewSectionService(rp repository.ISection) *SectionService {
	return &SectionService{rp: rp}
}

type SectionService struct {
	rp repository.ISection
}