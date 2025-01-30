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

func (s *SectionRepository) sectionNumberExist(sectionNumber int) bool {
	for _, section := range s.db {
		if section.SectionNumber == sectionNumber {
			return true
		}
	}
	return false
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

func (s *SectionRepository) Create(section model.Section) (model.Section, error) {
	if s.sectionNumberExist(section.SectionNumber) {
		return model.Section{}, errors.New("section number already exists")
	}
	lastId := s.db[len(s.db)].ID + 1
	section.ID = lastId
	s.db[section.ID] = section

	return s.db[lastId], nil
}

func (s *SectionRepository) Delete(id int) error {
	_, exists := s.db[id]
	if !exists {
		return errors.New("section not found")
	}
	delete(s.db, id)
	return nil
}
