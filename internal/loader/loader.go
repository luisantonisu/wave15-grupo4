package loader

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
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

	// load products
	productDb, err := LoadProducts()
	if err != nil {
		return nil, err
	}

	// load products
	sellerDB, err := LoadSellers()
	if err != nil {
		return nil, err
	}

	// load buyers
	buyersDb, err := LoadBuyers()
	if err != nil {
		return nil, err
	}

	sectionDb, err := LoadSections()
	if err != nil {
		return nil, err
	}

	// load warehouses
	warehousesDb, err := LoadWarehouses()
	if err != nil {
		return nil, err
	}

	return &DB{
		Buyers:     buyersDb,
		Employees:  employeesDb,
		Products:   productDb,
		Sections:   sectionDb,
		Sellers:    sellerDB,
		Warehouses: warehousesDb,
	}, nil
}

func LoadEmployees() (e map[int]model.Employee, err error) {
	// open file
	file, err := os.Open("./infrastructure/json/employees.json") //TODO static path
	if err != nil {
		return nil, errors.New("Error opening Employees file")
	}
	defer file.Close()

	// decode file
	var employeesJSON []dto.EmployeeResponseDTO
	err = json.NewDecoder(file).Decode(&employeesJSON)
	if err != nil {
		return nil, errors.New("Error decoding Employees file")
	}

	e = make(map[int]model.Employee)
	// serialize Employees
	for _, emp := range employeesJSON {
		e[emp.ID] = helper.EmployeeResponseDTOToEmployee(emp)
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
	var productsJSON []dto.ProductResponseDTO
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
				ProductTypeID:                  pr.ProductTypeId,
				SellerID:                       pr.SellerId,
			},
		}
	}

	return
}

func LoadBuyers() (b map[int]model.Buyer, err error) {
	// open file
	file, err := os.Open("./infrastructure/json/buyers.json")
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var buyersJSON []dto.BuyerDTO
	err = json.NewDecoder(file).Decode(&buyersJSON)
	if err != nil {
		return
	}

	// serialize buyers
	b = make(map[int]model.Buyer)
	for _, buyer := range buyersJSON {
		b[buyer.Id] = helper.BuyerDtoToBuyer(buyer)
	}
	return
}

func LoadSections() (s map[int]model.Section, err error) {
	// open file
	file, err := os.Open("./infrastructure/json/section.json")
	if err != nil {
		return nil, errors.New("Error opening Sections file")
	}
	defer file.Close()

	// decode file
	var sectionsJSON []dto.SectionRequestDTO
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return nil, errors.New("Error decoding Sections file")
	}

	// serialize sections
	s = make(map[int]model.Section)
	for key, value := range sectionsJSON {
		section := helper.SectionRequestDTOToSection(value)
		section.Id = key + 1
		s[key+1] = section
	}

	return
}

func LoadWarehouses() (w map[int]model.Warehouse, err error) {
	// open file
	file, err := os.Open("./infrastructure/json/warehouses.json")
	if err != nil {
		return
	}
	defer file.Close()

	var warehousesJSON []dto.WarehouseRequestDTO
	err = json.NewDecoder(file).Decode(&warehousesJSON)
	if err != nil {
		return
	}

	// serialize warehouses
	w = make(map[int]model.Warehouse)
	for key, value := range warehousesJSON {
		warehouse := helper.WarehouseRequestDTOToWarehouse(value)
		warehouse.Id = key + 1
		w[key+1] = warehouse
	}
	return
}

func LoadSellers() (map[int]model.Seller, error) {
	// open file
	file, err := os.Open("./infrastructure/json/sellers.json")
	if err != nil {
		return nil, err
	}
	var sellersJSON []dto.SellerRequestDTO
	err = json.NewDecoder(file).Decode(&sellersJSON)
	if err != nil {
		return nil, err
	}

	// serialize sellers
	data := make(map[int]model.Seller)
	for key, value := range sellersJSON {
		seller := helper.SellerRequestDTOToSeller(value)
		seller.ID = key + 1
		data[key+1] = seller
	}

	return data, nil
}
