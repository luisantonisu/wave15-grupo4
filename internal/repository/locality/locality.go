package repository

import (
	"database/sql"

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

func (r *LocalityRepository) GetByID(id int) (model.LocalityDBModel, error) {
	var locality model.LocalityDBModel
	err := r.db.QueryRow("SELECT id, locality_name, province_id FROM localities WHERE id = ?", id).Scan(
		&locality.Id, &locality.LocalityName, &locality.ProvinceID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY)
		}
		return model.LocalityDBModel{}, eh.GetErrParsingData(eh.LOCALITY)
	}

	return locality, nil
}

func (r *LocalityRepository) Create(locality model.LocalityDBModel) (model.LocalityDBModel, error) {
	hasIdAlreadyExist := r.localityIDExist(locality.Id)
	if hasIdAlreadyExist {
		return model.LocalityDBModel{}, eh.GetErrAlreadyExists(eh.LOCALITY_ID)
	}

	hasNameAlreadyExist := r.localityNameExist(locality.LocalityName)
	if hasNameAlreadyExist {
		return model.LocalityDBModel{}, eh.GetErrAlreadyExists(eh.LOCALITY_NAME)
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

func (r *LocalityRepository) SellersReport(id *int) ([]model.LocalityReport, error) {

	rows, err := r.buildingReportQuery(id)
	if err != nil {
		return nil, err
	}

	var localities []model.LocalityReport
	for rows.Next() {
		var locality model.LocalityReport
		err := rows.Scan(&locality.Id, &locality.LocalityName, &locality.SellerCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.LOCALITY)
		}
		localities = append(localities, locality)
	}

	return localities, nil
}

func (r *LocalityRepository) CarriersReport(id *int) ([]model.CarriersReport, error) {
	var rows *sql.Rows
	var err error
	query := `SELECT l.id, l.locality_name, COUNT(c.id) as carriers_count 
		FROM localities l 
		INNER JOIN carriers c ON l.id = c.locality_id 
		GROUP BY l.id, l.locality_name`

	if id != nil {
		query += " HAVING l.id = ?"
		rows, err = r.db.Query(query, id)
	} else {
		rows, err = r.db.Query(query)
	}
	if err != nil {
		return nil, eh.GetErrGettingData(eh.LOCALITY)
	}

	var report []model.CarriersReport
	for rows.Next() {
		var record model.CarriersReport
		err := rows.Scan(&record.LocalityID, &record.LocalityName, &record.CarriersCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.LOCALITY)
		}
		report = append(report, record)
	}
	return report, nil
}

func (r *LocalityRepository) localityIDExist(localityID int) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM localities WHERE ID = ?)", localityID).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}

func (r *LocalityRepository) localityNameExist(localityName string) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM localities WHERE locality_name = ?)", localityName).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}

func (r *LocalityRepository) buildingReportQuery(id *int) (rows *sql.Rows, err error) {

	if id != nil {
		//validate if locality exist
		hasIdAlreadyExist := r.localityIDExist(*id)
		if !hasIdAlreadyExist {
			return nil, eh.GetErrNotFound(eh.LOCALITY_ID)
		}

		//make query
		rows, err = r.db.Query(`SELECT l.id, l.locality_name, Count(*) FROM localities l 
						INNER JOIN sellers s ON l.id = s.locality_id GROUP BY l.id HAVING l.id = ?`, *id)
		if err != nil {
			return nil, eh.GetErrGettingData(eh.LOCALITY_ID)
		}

		return rows, nil
	}

	rows, err = r.db.Query(`SELECT l.id, l.locality_name, Count(*) FROM localities l 
						INNER JOIN sellers s ON l.id = s.locality_id GROUP BY l.id`)
	if err != nil {
		return nil, eh.GetErrGettingData(eh.LOCALITY_ID)
	}

	return rows, nil
}
