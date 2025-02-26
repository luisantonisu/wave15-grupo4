package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryEm "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	repositoryWh "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	firstNameEmp   = "John"
	firstNameEmpB  = "Jane"
	lastNameEmp    = "Doe"
	cardNumberID   = 12345
	cardNumberB    = 54321
	warehouseID    = 1
	mockEmployeeA   = model.Employee{
		ID: 1,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: &cardNumberID,
			FirstName:    &firstNameEmp,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		},
	}
	employeeRequest = model.Employee{
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: nil,
			FirstName:    &firstNameEmp,
			LastName:     nil,
			WarehouseID:  nil,
		},
	}
	mockEmployeeB = model.Employee{
		ID: 3,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		},
	}
)

func TestEmployeeService_Create(t *testing.T) {
	t.Run("case 1: create employee success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Create", mockEmployeeA).Return(mockEmployeeA, nil)

		employee, err := service.Create(mockEmployeeA)
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 2: create employee bad request - missing fields", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Create", employeeRequest).Return(model.Employee{}, eh.GetErrInvalidData(eh.EMPLOYEE))

		employee, err := service.Create(employeeRequest)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 3: create employee error - warehouse foreign key not found", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Create", mockEmployeeA).Return(model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		employee, err := service.Create(mockEmployeeA)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("case 4: create employee conflict - card number id duplicated ", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Create", mockEmployeeA).Return(model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER))

		employee, err := service.Create(mockEmployeeA)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Get(t *testing.T) {
	t.Run("case 1: get employee success by id", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("GetByID", 1).Return(mockEmployeeA, nil)

		employee, err := service.GetByID(1)
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 2: get employee by id - not found", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("GetByID", 2).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		employee, err := service.GetByID(2)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 3: get all employees success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("GetAll").Return(map[int]model.Employee{0: mockEmployeeA, 1: mockEmployeeB}, nil)

		employees, err := service.GetAll()
		require.NoError(t, err)
		require.Equal(t, map[int]model.Employee{0: mockEmployeeA, 1: mockEmployeeB}, employees)
		mockRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Update(t *testing.T) {
	t.Run("case 1: update employee success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)

		mockEmployeeA.EmployeeAttributes.FirstName = &firstNameEmpB
		mockRepo.On("Update", mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes).Return(mockEmployeeA, nil)

		employee, err := service.Update(mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes)
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 2: update employee - id not found", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Update", 2, mockEmployeeA.EmployeeAttributes).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		employee, err := service.Update(2, mockEmployeeA.EmployeeAttributes)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 3: update employee error - warehouse foreign key not found", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Update", mockEmployeeB.ID, mockEmployeeB.EmployeeAttributes).Return(model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		employee, err := service.Update(mockEmployeeB.ID, mockEmployeeB.EmployeeAttributes)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("case 4: update employee conflict - card number id duplicated ", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Update", mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes).Return(model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER))

		employee, err := service.Update(mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes)
		require.Error(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Delete(t *testing.T) {
	t.Run("case 1: delete employee success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Delete", 1).Return(nil)

		err := service.Delete(1)
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 2: delete employee - id not found", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)
		mockRepo.On("Delete", 2).Return(eh.GetErrNotFound(eh.EMPLOYEE))

		err := service.Delete(2)
		require.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Report(t *testing.T) {
	t.Run("case 1: get all employees report success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)

		reportMap := map[int]model.InboundOrdersReport{
			0: {
				ID: 1,
				FirstName: "John",
				LastName: "Doe",
				CardNumberID: 12345,
				WarehouseID: 1,
			},
			1: {
				ID: 2,
				FirstName: "Jane",
				LastName: "Wick",
				CardNumberID: 54321,
				WarehouseID: 1,
				InboundOrdersCount: 3,
			},
		}
		mockRepo.On("Report", -1).Return(reportMap, nil)

		report, err := service.Report(-1)
		require.NoError(t, err)
		require.Equal(t, reportMap, report)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 2: get by employee id, report success", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)

		reportMap := map[int]model.InboundOrdersReport{
			0: {
				ID: 1,
				FirstName: "John",
				LastName: "Doe",
				CardNumberID: 12345,
				WarehouseID: 1,
			},
		}
		mockRepo.On("Report", 1).Return(reportMap, nil)

		report, err := service.Report(1)
		require.NoError(t, err)
		require.Equal(t, reportMap, report)
		mockRepo.AssertExpectations(t)
	})

	t.Run("case 3: get by employee id report - id not found ", func(t *testing.T) {
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepository()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)

		mockRepo.On("Report", 10).Return(map[int]model.InboundOrdersReport{}, eh.GetErrNotFound(eh.EMPLOYEE))

		report, err := service.Report(-1)
		require.NoError(t, err)
		require.Equal(t, map[int]model.InboundOrdersReport{}, report)
		mockRepo.AssertExpectations(t)
	})
}