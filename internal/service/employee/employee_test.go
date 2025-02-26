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
	firstNameEmpA = "John"
	firstNameEmpB = "Jane"
	lastNameEmp   = "Doe"
	cardNumberA   = 12345
	cardNumberB   = 54321
	warehouseID   = 1
	mockEmployeeA = model.Employee{
		ID: 1,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: &cardNumberA,
			FirstName:    &firstNameEmpA,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		},
	}
	employeeRequest = model.Employee{
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: nil,
			FirstName:    &firstNameEmpA,
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
	t.Run("case 1: create employee successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: warehouseID}, nil)
		mockEmployeeRepo.On("Create", mockEmployeeA.EmployeeAttributes).Return(mockEmployeeA, nil)

		//Act
		employee, err := service.Create(mockEmployeeA)

		//Assert
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: invalid data - employee missing fields", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		//Act
		result, err := service.Create(employeeRequest)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, eh.GetErrInvalidData(eh.EMPLOYEE), err)
		require.Equal(t, model.Employee{}, result)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 3: foreign key - warehouse id not found", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE))

		//Act
		employee, err := service.Create(mockEmployeeA)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.WAREHOUSE), err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 4: conflict - card number id duplicated ", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: warehouseID}, nil)
		mockEmployeeRepo.On("Create", mockEmployeeA.EmployeeAttributes).Return(model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER))

		//Act
		employee, err := service.Create(mockEmployeeA)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER), err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Get(t *testing.T) {
	t.Run("case 1: get employee successfully by id", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		mockEmployeeRepo.On("GetByID", 1).Return(mockEmployeeA, nil)

		//Act
		employee, err := service.GetByID(1)

		//Assert
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: employee not found", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		mockEmployeeRepo.On("GetByID", 2).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		//Act
		employee, err := service.GetByID(2)

		//Assert
		require.Error(t, err)
		require.NotNil(t, err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 3: get all employees successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		mockEmployeeRepo.On("GetAll").Return(map[int]model.Employee{0: mockEmployeeA, 1: mockEmployeeB}, nil)

		//Act
		employees, err := service.GetAll()

		//Assert
		require.NoError(t, err)
		require.Equal(t, map[int]model.Employee{0: mockEmployeeA, 1: mockEmployeeB}, employees)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Update(t *testing.T) {
	t.Run("case 1: update employee successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeA.EmployeeAttributes.FirstName = &firstNameEmpB
		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: warehouseID}, nil)
		mockEmployeeRepo.On("Update", mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes).Return(mockEmployeeA, nil)

		//Act
		employee, err := service.Update(mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes)
		
		//Assert
		require.NoError(t, err)
		require.Equal(t, mockEmployeeA, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - employee doesn't exist", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		
		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: warehouseID}, nil)
		mockEmployeeRepo.On("Update", 2, mockEmployeeA.EmployeeAttributes).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		//Act
		employee, err := service.Update(2, mockEmployeeA.EmployeeAttributes)
		
		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, eh.GetErrNotFound(eh.EMPLOYEE), err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 3: foreign key - warehouse foreign key doesn't exist", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE))

		//Act
		employee, err := service.Update(mockEmployeeB.ID, mockEmployeeB.EmployeeAttributes)
		
		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.WAREHOUSE), err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 4: conflict - card number id duplicated ", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: warehouseID}, nil)
		mockEmployeeRepo.On("Update", mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes).Return(model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER))

		//Act
		employee, err := service.Update(mockEmployeeA.ID, mockEmployeeA.EmployeeAttributes)
		
		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER), err)
		require.Equal(t, model.Employee{}, employee)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Delete(t *testing.T) {
	t.Run("case 1: delete employee successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		
		mockEmployeeRepo.On("Delete", 1).Return(nil)

		//Act
		err := service.Delete(1)

		//Assert
		require.NoError(t, err)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - employee id doesn't exist", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)
		mockEmployeeRepo.On("Delete", 2).Return(eh.GetErrNotFound(eh.EMPLOYEE))

		//Act
		err := service.Delete(2)
		
		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, eh.GetErrNotFound(eh.EMPLOYEE), err)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_Report(t *testing.T) {
	t.Run("case 1: get all employees report successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		reportMap := map[int]model.InboundOrdersReport{
			0: {
				ID:           1,
				FirstName:    "John",
				LastName:     "Doe",
				CardNumberID: 12345,
				WarehouseID:  1,
			},
			1: {
				ID:                 2,
				FirstName:          "Jane",
				LastName:           "Wick",
				CardNumberID:       54321,
				WarehouseID:        1,
				InboundOrdersCount: 3,
			},
		}
		mockEmployeeRepo.On("Report", -1).Return(reportMap, nil)

		//Act
		report, err := service.Report(-1)

		//Assert
		require.NoError(t, err)
		require.Equal(t, reportMap, report)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: get report by employee id successfully", func(t *testing.T) {
		//Arrange
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockEmployeeRepo, mockWarehouseRepo)

		reportMap := map[int]model.InboundOrdersReport{
			0: {
				ID:           1,
				FirstName:    "John",
				LastName:     "Doe",
				CardNumberID: 12345,
				WarehouseID:  1,
			},
		}
		mockEmployeeRepo.On("Report", 1).Return(reportMap, nil)

		//Act
		report, err := service.Report(1)

		//Assert
		require.NoError(t, err)
		require.Equal(t, reportMap, report)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 3: not found - get by employee id report", func(t *testing.T) {
		//Arrange
		mockRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewEmployeeService(mockRepo, mockWarehouseRepo)

		mockRepo.On("Report", 10).Return(map[int]model.InboundOrdersReport{}, eh.GetErrNotFound(eh.EMPLOYEE))

		//Act
		report, err := service.Report(10)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, eh.GetErrNotFound(eh.EMPLOYEE), err)
		require.Equal(t, map[int]model.InboundOrdersReport{}, report)
		mockRepo.AssertExpectations(t)
	})
}
