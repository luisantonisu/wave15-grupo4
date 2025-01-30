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

func (s *SectionRepository) Patch(id int, section model.Section) (model.Section, error) {
	existingSection, exists := s.db[id]
	if !exists {
		return model.Section{}, errors.New("section not found")
	}

	if section.SectionNumber != 0 {
		existingSection.SectionNumber = section.SectionNumber
	}
	if section.CurrentTemperature != 0 {
		existingSection.CurrentTemperature = section.CurrentTemperature
	}
	if section.MinimumTemperature != 0 {
		existingSection.MinimumTemperature = section.MinimumTemperature
	}
	if section.CurrentCapacity != 0 {
		existingSection.CurrentCapacity = section.CurrentCapacity
	}
	if section.MinimumCapacity != 0 {
		existingSection.MinimumCapacity = section.MinimumCapacity
	}
	if section.MaximumCapacity != 0 {
		existingSection.MaximumCapacity = section.MaximumCapacity
	}
	if section.WarehouseID != 0 {
		existingSection.WarehouseID = section.WarehouseID
	}
	if section.ProductTypeID != 0 {
		existingSection.ProductTypeID = section.ProductTypeID
	}
	if len(section.ProductBatchID) != 0 {
		existingSection.ProductBatchID = section.ProductBatchID
	}

	s.db[id] = existingSection
	return existingSection, nil
}

func (s *SectionRepository) Delete(id int) error {
	_, exists := s.db[id]
	if !exists {
		return errors.New("section not found")
	}
	delete(s.db, id)
	return nil
}
