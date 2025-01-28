package loader

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

type DB struct {
	Buyers     map[int]model.Buyer
	Employees  map[int]model.Employee
	Products   map[int]model.Product
	Sections   map[int]model.Section
	Sellers    map[int]model.Seller
	Warehouses map[int]model.Warehouse
}

func Load() (*DB, error) {
	// load employees
	employeesDb, err := LoadEmployees()

	if err != nil {
		return nil, err
	}
    
  productDb, err := LoadProducts()
	if err != nil {
		return nil, err
	}

	return &DB{
		Buyers:     map[int]model.Buyer{},
		Employees:  employeesDb,
		Products:   productDb,
		Sections:   map[int]model.Section{},
		Sellers:    map[int]model.Seller{},
		Warehouses: map[int]model.Warehouse{},
	}, nil
}

func LoadEmployees() (map[int]model.Employee, error) {
	// open file
	file, err := os.Open("./infrastructure/json/employees.json") //TODO static path
	if err != nil {
		return nil, errors.New("Error opening Employees file")
	}
	defer file.Close()

	// decode file
	var employeesJSON []dto.EmployeeDTO
	err = json.NewDecoder(file).Decode(&employeesJSON)
	if err != nil {
		return nil, errors.New("Error decoding Employees file")
	}

	// serialize Employees
	e := make(map[int]model.Employee)
	for _, emp := range employeesJSON {
		e[emp.Id] = model.Employee{
			Id: emp.Id,
			EmployeeAttributes: model.EmployeeAttributes{
				CardNumberId: emp.CardNumberId,
				FirstName:    emp.FirstName,
				LastName:     emp.LastName,
				WarehouseId:  emp.WarehouseId,
			},
		}
	}

	return e, nil
}
   
func LoadProducts() (p map[int]model.Product, err error) {
	// open file
	file, err := os.Open("./infrastructure/json/product.json")
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var productsJSON []dto.ProductDTO
	err = json.NewDecoder(file).Decode(&productsJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	p = make(map[int]model.Product)
	for _, pr := range productsJSON {
		p[pr.ID] = model.Product{
			ID: pr.ID,
			ProductAtrributes: model.ProductAtrributes{
				ProductCode:                    pr.ProductCode,
				Description:                    pr.Description,
				Width:                          pr.Width,
				Height:                         pr.Height,
				Length:                         pr.Length,
				NetWeight:                      pr.NetWeight,
				ExpirationRate:                 pr.ExpirationRate,
				RecommendedFreezingTemperature: pr.RecommendedFreezingTemperature,
				FreezingRate:                   pr.FreezingRate,
				ProductTypeId:                  pr.ProductTypeId,
				SellerId:                       pr.SellerId,
			},
		}
	}

	return
}

