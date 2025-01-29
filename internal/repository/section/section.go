package repository

import (
	"errors"

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

func (s *SectionRepository) GetAll() (map[int]model.Section, error) {
	return s.db, nil
}

func (s *SectionRepository) GetByID(id int) (model.Section, error) {
	section, exists := s.db[id]
	if !exists {
		return model.Section{}, errors.New("not exist")
	}
	return section, nil
}
