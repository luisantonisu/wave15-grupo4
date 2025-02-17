package repository

import (
	"database/sql"
	"strconv"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewLocalityRepository(db *sql.DB) *LocalityRepository {
	return &LocalityRepository{
		db: db,
	}
}

type LocalityRepository struct {
	db *sql.DB
}

func (r *LocalityRepository) Create(locality model.LocalityDBModel) (model.LocalityDBModel, error) {
	hasIdAlreadyExist := r.localityIDExist(locality.Id)
	if hasIdAlreadyExist {
		return model.LocalityDBModel{}, eh.GetErrAlreadyExists(eh.LOCALITY_ID)
	}

	row, err := r.db.Exec("INSERT INTO localities (id, locality_name, province_id) VALUES (?,?,?)",
		locality.Id, locality.LocalityName, locality.ProvinceID)
	if err != nil {
		return model.LocalityDBModel{}, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.LocalityDBModel{}, err
	}

	var newLocality model.LocalityDBModel
	newLocality.Id = int(id)
	newLocality.LocalityName = locality.LocalityName
	newLocality.ProvinceID = locality.ProvinceID

	// return new seller
	return newLocality, nil
}

func (r *LocalityRepository) localityIDExist(localityID int) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM localities WHERE ID = ?)", localityID).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}

func (r *LocalityRepository) Report(id int) (map[int]model.CarriersByLocalityReport, error) {
	query := `SELECT l.id, l.locality_name, COUNT(c.id) as carriers_count 
		FROM localities l 
		INNER JOIN carriers c ON l.id = c.locality_id 
		GROUP BY l.id, l.locality_name`

	if id != -1 {
		if !r.localityIDExist(id) {
			return nil, eh.GetErrNotFound(eh.LOCALITY)
		}
		query += " HAVING l.id = " + strconv.Itoa(id)
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, eh.GetErrGettingData(eh.LOCALITY)
	}

	report := make(map[int]model.CarriersByLocalityReport)
	for rows.Next() {
		var record model.CarriersByLocalityReport
		err := rows.Scan(&record.LocalityID, &record.LocalityName, &record.CarriersCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.LOCALITY)
		}
		report[record.LocalityID] = record
	}
	return report, nil
}
