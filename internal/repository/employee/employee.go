package repository

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewEmployeeRepository(db map[int]model.Employee) *EmployeeRepository {
	defaultDb := make(map[int]model.Employee)
	if db != nil {
		defaultDb = db
	}
	return &EmployeeRepository{db: defaultDb}
}

type EmployeeRepository struct {
	db map[int]model.Employee
}

func (r *EmployeeRepository) employeeExists(id int) bool {
	_, exists := r.db[id]
	return exists
}

func (r *EmployeeRepository) cardNumberIdExists(cardNumberId int) bool {
	for _, employee := range r.db {
		if employee.CardNumberID == cardNumberId {
			return true
		}
	}

	return false
}

func (r *EmployeeRepository) GetAll() (map[int]model.Employee, error) {
	return r.db, nil
}

func (r *EmployeeRepository) GetByID(id int) (model.Employee, error) {
	if !r.employeeExists(id) {
		return model.Employee{}, errors.New("Employee not found")
	}

	return r.db[id], nil
}

func (r *EmployeeRepository) Create(employee model.Employee) (model.Employee, error) {

	if r.cardNumberIdExists(employee.CardNumberID) {
		return model.Employee{}, errors.New("Card number ID already exists")
	}

	lastId := r.db[len(r.db)].ID + 1
	employee.ID = lastId
	r.db[lastId] = employee

	return r.db[lastId], nil
}

func (r *EmployeeRepository) Update(id int, employee model.Employee) (model.Employee, error) {
	if !r.employeeExists(id) {
		return model.Employee{}, errors.New("Employee not found")
	}

	if r.cardNumberIdExists(employee.CardNumberID) {
		return model.Employee{}, errors.New("Card number ID already exists")
	}

	emp := r.db[id]

	if employee.FirstName != "" {
		emp.FirstName = employee.FirstName
	}

	if employee.LastName != "" {
		emp.LastName = employee.LastName
	}

	if employee.CardNumberID > 0 {
		emp.CardNumberID = employee.CardNumberID
	}

	if employee.WarehouseID > 0 {
		emp.WarehouseID = employee.WarehouseID
	}

	r.db[id] = emp
	return r.db[id], nil
}

func (r *EmployeeRepository) Delete(id int) error {
	if !r.employeeExists(id) {
		return errors.New("Employee not found")
	}

	delete(r.db, id)
	return nil
}
