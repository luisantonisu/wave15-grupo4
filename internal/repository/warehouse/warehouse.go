package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewWarehouseRepository(db map[int]model.Warehouse) *WarehouseRepository {
	defaultDb := make(map[int]model.Warehouse)
	if db != nil {
		defaultDb = db
	}
	return &WarehouseRepository{db: defaultDb}
}

type WarehouseRepository struct {
	db map[int]model.Warehouse
}
