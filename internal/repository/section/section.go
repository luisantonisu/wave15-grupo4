package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
	if !s.sectionExists(id) {
		return model.Section{}, eh.GetErrNotFound(eh.SECTION)
	}
	return s.db[id], nil
}

func (s *SectionRepository) Create(section model.Section) (model.Section, error) {
	if s.sectionNumberExist(section.SectionNumber) {
		return model.Section{}, eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
	}
	lastId := s.db[len(s.db)].ID + 1
	section.ID = lastId
	s.db[section.ID] = section

	return s.db[lastId], nil
}

func (s *SectionRepository) Patch(id int, section model.SectionAttributesPtr) (model.Section, error) {
	if !s.sectionExists(id) {
		return model.Section{}, eh.GetErrNotFound(eh.SECTION)
	}

	if section.SectionNumber != nil && s.sectionNumberExist(*section.SectionNumber) {
		return model.Section{}, eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
	}

	sec := s.db[id]

	if section.SectionNumber != nil {
		sec.SectionNumber = *section.SectionNumber
	}
	if section.CurrentTemperature != nil {
		sec.CurrentTemperature = *section.CurrentTemperature
	}
	if section.MinimumTemperature != nil {
		sec.MinimumTemperature = *section.MinimumTemperature
	}
	if section.CurrentCapacity != nil {
		sec.CurrentCapacity = *section.CurrentCapacity
	}
	if section.MinimumCapacity != nil {
		sec.MinimumCapacity = *section.MinimumCapacity
	}
	if section.MaximumCapacity != nil {
		sec.MaximumCapacity = *section.MaximumCapacity
	}
	if section.WarehouseID != nil {
		sec.WarehouseID = *section.WarehouseID
	}
	if section.ProductTypeID != nil {
		sec.ProductTypeID = *section.ProductTypeID
	}
	if section.ProductBatchID != nil {
		sec.ProductBatchID = *section.ProductBatchID
	}

	s.db[id] = sec
	return sec, nil
}

func (s *SectionRepository) Delete(id int) error {
	if !s.sectionExists(id) {
		return eh.GetErrNotFound(eh.SECTION)
	}
	delete(s.db, id)
	return nil
}
