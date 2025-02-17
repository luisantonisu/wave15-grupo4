package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewSectionRepository(db *sql.DB) *SectionRepository {
	return &SectionRepository{
		db: db,
	}
}

type SectionRepository struct {
	db *sql.DB
}

func (s *SectionRepository) sectionExists(id int) bool {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (s *SectionRepository) sectionNumberExist(sectionNumber int) bool {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE section_number = ?)", sectionNumber).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (s *SectionRepository) GetAll() (map[int]model.Section, error) {
	rows, err := s.db.Query("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections")
	if err != nil {
		return nil, eh.GetErrGettingData(eh.SECTION)
	}
	sections := make(map[int]model.Section)
	for rows.Next() {
		var section model.Section
		err := rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.SECTION)
		}
		sections[section.ID] = section
	}
	return sections, nil
}

func (s *SectionRepository) GetByID(id int) (section model.Section, err error) {
	err = s.db.QueryRow("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?", id).Scan(
		&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		return model.Section{}, eh.GetErrNotFound(eh.SECTION)
	}
	return section, nil
}

func (s *SectionRepository) Create(section model.SectionAttributes) (model.Section, error) {
	if s.sectionNumberExist(section.SectionNumber) {
		return model.Section{}, eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
	}

	row, err := s.db.Exec("INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseID, section.ProductTypeID)
	if err != nil {
		return model.Section{}, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Section{}, err
	}

	var sect model.Section
	sect.ID = int(id)
	sect.SectionAttributes = section

	return sect, nil
}

func (s *SectionRepository) Patch(id int, section model.SectionAttributesPtr) (model.Section, error) {
	if !s.sectionExists(id) {
		return model.Section{}, eh.GetErrNotFound(eh.SECTION)
	}

	if section.SectionNumber != nil && s.sectionNumberExist(*section.SectionNumber) {
		return model.Section{}, eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
	}

	var sec model.Section
	err := s.db.QueryRow("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?", id).Scan(
		&sec.ID, &sec.SectionNumber, &sec.CurrentTemperature, &sec.MinimumTemperature, &sec.CurrentCapacity, &sec.MinimumCapacity, &sec.MaximumCapacity, &sec.WarehouseID, &sec.ProductTypeID)

	if err != nil {
		return model.Section{}, eh.GetErrNotFound(eh.SECTION)
	}

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
	_, err = s.db.Exec("UPDATE sections SET section_number = ?, current_temperature = ?, minimum_temperature = ?, current_capacity = ?, minimum_capacity = ?, maximum_capacity = ?, warehouse_id = ?, product_type_id = ?, WHERE id = ?",
		sec.SectionNumber, sec.CurrentTemperature, sec.MinimumTemperature, sec.CurrentCapacity, sec.MinimumCapacity, sec.MaximumCapacity, sec.WarehouseID, sec.ProductTypeID, id)

	if err != nil {
		return model.Section{}, eh.GetErrInvalidData(eh.SECTION)
	}

	return sec, nil
}

func (s *SectionRepository) Delete(id int) error {
	if !s.sectionExists(id) {
		return eh.GetErrNotFound(eh.SECTION)
	}

	_, err := s.db.Exec("DELETE FROM sections WHERE id = ?", id)
	if err != nil {
		return eh.GetErrNotFound(eh.SECTION)
	}

	return nil
}

func (s *SectionRepository) Report(id int) (map[int]model.ReportProductsBatches, error) {

	var rows *sql.Rows
	var err error

	query := "SELECT sections.id, sections.section_number, COUNT(*) AS products_count FROM sections JOIN product_batches ON sections.id = product_batches.section_id "

	if id != -1 {
		if !s.sectionExists(id) {
			return nil, eh.GetErrNotFound(eh.SECTION)
		}
		query += "WHERE sections.id = ? GROUP BY sections.id"
		rows, err = s.db.Query(query, id)
	} else {
		query += "GROUP BY sections.id"
		rows, err = s.db.Query(query)
	}

	if err != nil {
		return nil, eh.GetErrGettingData(eh.SECTION)
	}

	productsBatches := make(map[int]model.ReportProductsBatches)
	for rows.Next() {
		var productBatch model.ReportProductsBatches
		err := rows.Scan(&productBatch.SectionID, &productBatch.SectionNumber, &productBatch.ProductsCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.SECTION)
		}
		productsBatches[productBatch.SectionID] = productBatch
	}
	return productsBatches, nil
}
