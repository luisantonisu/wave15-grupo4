package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewEmployeeService(employeeRp employeeRepository.IEmployee, warehouseRp warehouseRepository.IWarehouse) *EmployeeService {
	return &EmployeeService{
		employeeRp:  employeeRp,
		warehouseRp: warehouseRp,
	}
}

type EmployeeService struct {
	employeeRp  employeeRepository.IEmployee
	warehouseRp warehouseRepository.IWarehouse
}

func (h *EmployeeService) GetAll() (map[int]model.Employee, error) {
	return h.employeeRp.GetAll()
}

func (h *EmployeeService) GetByID(id int) (model.Employee, error) {
	return h.employeeRp.GetByID(id)
}

func (h *EmployeeService) Create(employee model.Employee) (model.Employee, error) {
	if employee.FirstName == nil || employee.LastName == nil || employee.CardNumberID == nil || employee.WarehouseID == nil {
		return model.Employee{}, eh.GetErrInvalidData(eh.EMPLOYEE)
	}

	_, err := h.warehouseRp.GetByID(*employee.WarehouseID)
	if err != nil {
		return model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE)
	}

	return h.employeeRp.Create(employee.EmployeeAttributes)
}

func (h *EmployeeService) Update(id int, employee model.EmployeeAttributes) (model.Employee, error) {
	if employee.WarehouseID != nil {
		_, err := h.warehouseRp.GetByID(*employee.WarehouseID)
		if err != nil {
			return model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE)
		}
	}

	return h.employeeRp.Update(id, employee)
}

func (h *EmployeeService) Delete(id int) error {
	return h.employeeRp.Delete(id)
}

func (h *EmployeeService) Report(id int) (map[int]model.InboundOrdersReport, error) {
	return h.employeeRp.Report(id)
}
