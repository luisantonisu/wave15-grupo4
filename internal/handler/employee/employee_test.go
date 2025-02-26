package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	firstNameEmp   = "John"
	firstNameEmpB  = "Jane"
	lastNameEmp    = "Doe"
	cardNumberID   = 12345
	cardNumberB    = 54321
	warehouseID    = 1
	warehouseBadID = 234
	mockEmployee   = model.Employee{
		ID: 1,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: &cardNumberID,
			FirstName:    &firstNameEmp,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		},
	}
	employeeRequest = dto.EmployeeRequestDTO{
		CardNumberID: &cardNumberID,
		FirstName:    &firstNameEmp,
		LastName:     &lastNameEmp,
		WarehouseID:  &warehouseID,
	}
	mockEmployeeUpdate = model.Employee{
		ID: 1,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		},
	}
)

func TestEmployeeHandler_Create(t *testing.T) {

	t.Run("case 1: create employee successfully", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Create", mock.Anything).Return(mockEmployee, nil)

		body, err := json.Marshal(employeeRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/employees", employeeHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"card_number_id":12345,"first_name":"John","last_name":"Doe","warehouse_id":1}
		}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: invalid data - employee missing fields", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Create", mock.Anything).Return(model.Employee{}, eh.GetErrInvalidData(eh.EMPLOYEE))

		body, err := json.Marshal(dto.EmployeeRequestDTO{
			CardNumberID: nil,
			FirstName:    &firstNameEmp,
			LastName:     nil,
			WarehouseID:  nil,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/employees", employeeHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Unprocessable Entity",
			"message": "invalid data: employee"
		}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: conflict - card number id duplicated", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Create", mock.Anything).Return(model.Employee{}, eh.GetErrAlreadyExistsCompose(eh.EMPLOYEE, eh.CARD_NUMBER))

		body, err := json.Marshal(dto.EmployeeRequestDTO{
			CardNumberID: &cardNumberID,
			FirstName:    &firstNameEmp,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/employees", employeeHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "employee with that card number ID already exists"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: not found - warehouse id doesn't exist", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Create", mock.Anything).Return(model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		body, err := json.Marshal(dto.EmployeeRequestDTO{
			CardNumberID: &cardNumberID,
			FirstName:    &firstNameEmp,
			LastName:     &lastNameEmp,
			WarehouseID:  &warehouseBadID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/employees", employeeHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "warehouse foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 5: bad request - invalid body", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Create", mock.Anything).Return(model.Employee{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			"CardNumberID": "12345",
			"FirstName":    "John",
			"LastName":     "Doe",
			"WarehouseID":  1,
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/employees", employeeHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid request body"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}

func TestEmployeeHandler_Get(t *testing.T) {
	t.Run("case 1: get all employees successfully", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("GetAll").Return([]model.Employee{mockEmployee}, nil)

		rt := chi.NewRouter()
		rt.Get("/employees", employeeHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [{"id":1,"card_number_id":12345,"first_name":"John","last_name":"Doe","warehouse_id":1}]
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: internal server error - get all employees", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("GetAll").Return([]model.Employee{}, eh.GetErrDatabase(eh.EMPLOYEE))

		rt := chi.NewRouter()
		rt.Get("/employees", employeeHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Internal Server Error",
			"message": "database error: employee"
		}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: get employee by id successfully", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("GetByID", 1).Return(mockEmployee, nil)

		rt := chi.NewRouter()
		rt.Get("/employees/{id}", employeeHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"card_number_id":12345,"first_name":"John","last_name":"Doe","warehouse_id":1}
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: not found - employee id doesn't exist", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("GetByID", 2).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		rt := chi.NewRouter()
		rt.Get("/employees/{id}", employeeHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Not Found",
			"message": "employee not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 5: bad request - employee invalid id", func(t *testing.T) {
		// Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("GetByID", "hi").Return(model.Employee{}, eh.INVALID_ID)

		rt := chi.NewRouter()
		rt.Get("/employees/{id}", employeeHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/hi", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestEmployeeHandler_Update(t *testing.T) {

	t.Run("case 1: update employee successfully", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Update", 1, mock.Anything).Return(mockEmployeeUpdate, nil)

		body, err := json.Marshal(model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/employees/{id}", employeeHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/employees/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"card_number_id":54321,"first_name":"Jane","last_name":"Doe","warehouse_id":1}
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: not found - employee id doesn't exist", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Update", 2, mock.Anything).Return(model.Employee{}, eh.GetErrNotFound(eh.EMPLOYEE))

		body, err := json.Marshal(model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/employees/{id}", employeeHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/employees/2", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Not Found",
			"message": "employee not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: bad request - employee invalid id", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Update", "hi", mock.Anything).Return(model.Employee{}, eh.INVALID_ID)

		body, err := json.Marshal(model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/employees/{id}", employeeHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/employees/hi", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: invalid data - invalid request body", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Update", 1, mock.Anything).Return(model.Employee{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			"CardNumberID": "5442",
			"FirstName":    "Jane"
			}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/employees/{id}", employeeHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/employees/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid request body"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 5: foreign key - warehouse id doesn't exist", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Update", 1, mock.Anything).Return(model.Employee{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		body, err := json.Marshal(model.EmployeeAttributes{
			CardNumberID: &cardNumberB,
			FirstName:    &firstNameEmpB,
			WarehouseID: &warehouseBadID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/employees/{id}", employeeHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/employees/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "warehouse foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestEmployeeHandler_Delete(t *testing.T) {
	t.Run("case 1: delete employee successfully", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Delete", 1).Return(nil)

		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", employeeHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/employees/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
	})
	
	t.Run("case 2: not found - employee id doesn't exist", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Delete", 2).Return(eh.GetErrNotFound(eh.EMPLOYEE))

		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", employeeHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/employees/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expected := `{
			"status": "Not Found",
			"message": "employee not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})

	t.Run("case 3: bad request - employee invalid id", func(t *testing.T) {
		//Arranq
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Delete", "hi").Return(eh.INVALID_ID)

		rt := chi.NewRouter()
		rt.Delete("/employees/{id}", employeeHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/employees/hi", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})
}

func TestEmployeeHandler_Report(t *testing.T) {
	t.Run("case 1: get report by employee id successfully", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)

		employeeService.On("Report", 1).Return([]model.InboundOrdersReport{
			{
				ID: 1,
				CardNumberID: cardNumberID,
				FirstName: firstNameEmp,
				LastName: lastNameEmp,
				WarehouseID: warehouseID,
				InboundOrdersCount: 2,
			},
		}, nil)

		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", employeeHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [{"id":1,"card_number_id":12345,"first_name":"John","last_name":"Doe","warehouse_id":1,"inbound_orders_count":2}]
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})
	
	t.Run("case 2: get all employees report successfully", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)

		employeeService.On("Report", -1).Return([]model.InboundOrdersReport{
			{
				ID: 1,
				CardNumberID: cardNumberID,
				FirstName: firstNameEmp,
				LastName: lastNameEmp,
				WarehouseID: warehouseID,
				InboundOrdersCount: 2,
			},
			{
				ID: 2,
				CardNumberID: cardNumberB,
				FirstName: firstNameEmpB,
				LastName: lastNameEmp,
				WarehouseID: warehouseBadID,
				InboundOrdersCount: 0,
			},
		}, nil)

		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", employeeHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [
				{"id":1,"card_number_id":12345,"first_name":"John","last_name":"Doe","warehouse_id":1,"inbound_orders_count":2},
				{"id":2,"card_number_id":54321,"first_name":"Jane","last_name":"Doe","warehouse_id":234,"inbound_orders_count":0}
			]
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: bad request - get by employee id report", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Report", "hi").Return([]model.InboundOrdersReport{}, eh.INVALID_ID)

		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", employeeHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=hi", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)


		// Assert
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})
	
	t.Run("case 4: not found - get by employee id report", func(t *testing.T) {
		//Arrange
		employeeService := service.NewEmployeeMock()
		employeeHandler := NewEmployeeHandler(employeeService)
		employeeService.On("Report", 2).Return([]model.InboundOrdersReport{}, eh.GetErrNotFound(eh.EMPLOYEE))

		rt := chi.NewRouter()
		rt.Get("/employees/reportInboundOrders", employeeHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/employees/reportInboundOrders?id=2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)


		// Assert
		expected := `{
			"status": "Not Found",
			"message": "employee not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})
}
