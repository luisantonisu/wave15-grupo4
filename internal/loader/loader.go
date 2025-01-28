package loader

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type DB struct {
	Buyers   map[int]model.Buyer
	Employees map[int]model.Employee
	Products map[int]model.Product
	Sections map[int]model.Section
	Sellers map[int]model.Seller
	Warehouses map[int]model.Warehouse
}

func Load() (*DB, error) {
	return &DB{
		Buyers: map[int]model.Buyer{},
		Employees: map[int]model.Employee{},
		Products: map[int]model.Product{},
		Sections: map[int]model.Section{},
		Sellers: map[int]model.Seller{},
		Warehouses: map[int]model.Warehouse{},
	}, nil
}
