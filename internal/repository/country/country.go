package repository

import (
	"database/sql"

	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

type CountryRepository struct {
	db *sql.DB
}

func (r *CountryRepository) GetCountryIDByCountryName(countryName string) (int, error) {
	//verify if country exist
	hasCountryAlreadyExist := r.countryNameExist(countryName)
	if !hasCountryAlreadyExist {
		return 0, eh.GetErrNotFound(eh.COUNTRY_NAME)
		
	}

	//search countryId
	var countryId int
	err := r.db.QueryRow("SELECT id FROM countries WHERE country_name = ?", countryName).Scan(&countryId)
	if err != nil {
		return 0, eh.GetErrNotFound(eh.COUNTRY_NAME)
	}

	return countryId, nil
}

func (r *CountryRepository) countryNameExist(countryName string) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM countries WHERE country_name = ?)", countryName).Scan(&exist)
	if err != nil {
		return false
	}

	return exist

}
