package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewSectionRepository(db map[int]model.Section) *SectionRepository {
	defaultDb := make(map[int]model.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionRepository{db: defaultDb}
}

type SectionRepository struct {
	db map[int]model.Section
}

func (s *SectionRepository) sectionExists(id int) bool {
	_, exists := s.db[id]
	return exists
}

func (s *SectionRepository) GetAll() (map[int]model.Section, error) {
	return s.db, nil
}
