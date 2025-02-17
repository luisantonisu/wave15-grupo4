package repository

import (
	"database/sql"

	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProvinceRepository(db *sql.DB) *ProvinceRepository {
	return &ProvinceRepository{
		db: db,
	}
}

type ProvinceRepository struct {
	db *sql.DB
}

func (r *ProvinceRepository) GetProvinceID(countryId int, provinceName string) (int, error) {
	//verify if province exist
	hasProvinceAlreadyExist := r.provinceNameExist(provinceName)
	if !hasProvinceAlreadyExist {
		return 0, eh.GetErrNotFound(eh.PROVINCE_NAME)
		
	}

	//search provinceId and verify if have relation with country
	var provinceId int
	err := r.db.QueryRow("SELECT id FROM provinces WHERE country_id = ? AND province_name = ?", countryId, provinceName).Scan(&provinceId)
	if err != nil {
		return 0, eh.GetErrNotFound(eh.PROVINCE_ID)
	}

	return countryId, nil
}

func (r *ProvinceRepository) provinceNameExist(provinceName string) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM provinces WHERE province_name = ?)", provinceName).Scan(&exist)
	if err != nil {
		return false
	}

	return exist

}
