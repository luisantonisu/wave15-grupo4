package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

type EmployeeRepository struct {
	db *sql.DB
}

func (r *EmployeeRepository) employeeExists(id int) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE id = ?) ", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *EmployeeRepository) cardNumberIdExists(cardNumberId int, id int) bool {
	var exists bool
	var err error
	if id < 0 {
		err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE card_number_id = ?) ", cardNumberId).Scan(&exists)
	} else {
		err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE card_number_id = ?) AND id != ?", cardNumberId, id).Scan(&exists)
	}
	
	if err != nil {
		return false
	}
	return exists
}

func (r *EmployeeRepository) GetAll() (map[int]model.Employee, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, card_number_id, warehouse_id FROM employees")
	if err != nil {
		return nil, eh.GetErrGettingData(eh.EMPLOYEE)
	}

	employees := make(map[int]model.Employee)
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.EMPLOYEE)
		}
		employees[employee.ID] = employee
	}

	return employees, nil
}

func (r *EmployeeRepository) GetByID(id int) (employee model.Employee, err error) {
	err = r.db.QueryRow("SELECT id, first_name, last_name, card_number_id, warehouse_id FROM employees WHERE id = ?", id).Scan(
		&employee.ID, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID)
	if err != nil {
		return model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE)
	}

	return employee, nil
}

func (r *EmployeeRepository) Create(employee model.EmployeeAttributes) (model.Employee, error) {

	if r.cardNumberIdExists(*employee.CardNumberID, -1) {
		return model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER)
	}

	row, err := r.db.Exec("INSERT INTO employees (first_name, last_name, card_number_id, warehouse_id) VALUES (?, ?, ?, ?)",
		employee.FirstName, employee.LastName, employee.CardNumberID, employee.WarehouseID)
	if err != nil {
		return model.Employee{}, eh.GetErrInvalidData(eh.EMPLOYEE)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Employee{}, eh.GetErrDatabase(eh.EMPLOYEE)
	}

	var emp model.Employee
	emp.ID = int(id)
	emp.EmployeeAttributes = employee

	return emp, nil
}

func (r *EmployeeRepository) Update(id int, employee model.EmployeeAttributes) (model.Employee, error) {
	if !r.employeeExists(id) {
		return model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE)
	}

	if employee.CardNumberID != nil && r.cardNumberIdExists(*employee.CardNumberID, id) {
		return model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER)
	}

	var emp model.Employee
	err := r.db.QueryRow("SELECT id, first_name, last_name, card_number_id, warehouse_id FROM employees WHERE id = ?", id).Scan(
		&emp.ID, &emp.EmployeeAttributes.FirstName, &emp.EmployeeAttributes.LastName, &emp.EmployeeAttributes.CardNumberID, &emp.EmployeeAttributes.WarehouseID)

	if err != nil {
		return model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE)
	}

	if employee.FirstName != nil {
		emp.EmployeeAttributes.FirstName = employee.FirstName
	}

	if employee.LastName != nil {
		emp.EmployeeAttributes.LastName = employee.LastName
	}

	if employee.CardNumberID != nil {
		emp.EmployeeAttributes.CardNumberID = employee.CardNumberID
	}

	if employee.WarehouseID != nil {
		emp.EmployeeAttributes.WarehouseID = employee.WarehouseID
	}

	_, err = r.db.Exec("UPDATE employees SET first_name = ?, last_name = ?, card_number_id = ?, warehouse_id = ? WHERE id = ?",
		emp.EmployeeAttributes.FirstName, emp.EmployeeAttributes.LastName, emp.EmployeeAttributes.CardNumberID, emp.EmployeeAttributes.WarehouseID, id)
	if err != nil {
		return model.Employee{}, eh.GetErrInvalidData(eh.EMPLOYEE)
	}

	return emp, nil
}

func (r *EmployeeRepository) Delete(id int) error {
	if !r.employeeExists(id) {
		return eh.GetErrNotFound(eh.EMPLOYEE)
	}

	_, err := r.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return eh.GetErrDatabase(eh.EMPLOYEE)
	}

	return nil
}

func (r *EmployeeRepository) Report(id int) (map[int]model.InboundOrdersReport, error) {
	var rows *sql.Rows
	var err error
	if id == -1 {
		rows, err = r.db.Query("SELECT em.id, em.first_name, em.last_name, em.card_number_id, em.warehouse_id, COUNT(*) as inbound_orders_count FROM employees em INNER JOIN inbound_orders ib ON em.id = ib.employee_id GROUP BY em.id")
		if err != nil {
			return nil, eh.GetErrGettingData(eh.EMPLOYEE)
		}
	} else {
		if !r.employeeExists(id) {
			return nil, eh.GetErrNotFound(eh.EMPLOYEE)
		}

		rows, err = r.db.Query("SELECT em.id, em.first_name, em.last_name, em.card_number_id, em.warehouse_id, COUNT(*) as inbound_orders_count FROM employees em INNER JOIN inbound_orders ib ON em.id = ib.employee_id GROUP BY em.id HAVING em.id = ?", id)
		if err != nil {
			return nil, eh.GetErrGettingData(eh.EMPLOYEE)
		}
	}

	employees := make(map[int]model.InboundOrdersReport)
	for rows.Next() {
		var employee model.InboundOrdersReport
		err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.CardNumberID, &employee.WarehouseID, &employee.InboundOrdersCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.EMPLOYEE)
		}
		employees[employee.ID] = employee
	}

	return employees, nil
}
