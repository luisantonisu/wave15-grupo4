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

