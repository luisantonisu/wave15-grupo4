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
	service "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	companyID   = "1234"
	CompanyIDB  = "2345"
	companyName = "Company X"
	address     = "221B Baker Street"
	telephone   = "123590"
	localityId  = "1"

	seller = model.Seller{
		ID: 1,
		SellerAttributes: model.SellerAttributes{
			CompanyID:   &companyID,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityId:  &localityId,
		},
	}
	sellerB = model.Seller{
		ID: 2,
		SellerAttributes: model.SellerAttributes{
			CompanyID:   &CompanyIDB,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityId:  &localityId,
		},
	}
	sellerRequest = dto.SellerRequestDTO{
		CompanyID:   &CompanyIDB,
		CompanyName: &companyName,
		Address:     &address,
		Telephone:   &telephone,
		LocalityId:  &localityId,
	}
)

func TestSellerHandler_GetAll(t *testing.T) {
	t.Run("case 1: Get all sellers successfully", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("GetAll").Return([]model.Seller{seller, sellerB}, nil)

		rt := chi.NewRouter()
		rt.Get("/sellers", sellerHd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"data":[
        {"id":1,"cid":"1234","company_name":"Company X","address":"221B Baker Street","telephone":"123590","locality_id":"1"},
		{"id":2,"cid":"2345","company_name":"Company X","address":"221B Baker Street","telephone":"123590","locality_id":"1"}
		]}`

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})

	t.Run("case 2: sellers not found", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("GetAll").Return([]model.Seller{}, nil)

		rt := chi.NewRouter()
		rt.Get("/sellers", sellerHd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":null}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 3: internal server error", func(t *testing.T) {
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("GetAll").Return([]model.Seller{}, eh.GetErrDatabase(eh.SELLER))

		rt := chi.NewRouter()
		rt.Get("/sellers", sellerHd.GetAll())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: seller"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
}

func TestSellersHandler_GetByID(t *testing.T) {
	t.Run("case 1: get sellers by ID", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)
		sellerSv.On("GetByID", 1).Return(seller, nil)

		rt := chi.NewRouter()
		rt.Get("/sellers/{id}", sellerHd.GetByID())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"cid":"1234","company_name":"Company X","address":"221B Baker Street","telephone":"123590","locality_id":"1"}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 2: not found - Seller ID doesn't exist", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)
		sellerSv.On("GetByID", 2).Return(model.Seller{}, eh.GetErrNotFound(eh.SELLER))

		rt := chi.NewRouter()
		rt.Get("/sellers/{id}", sellerHd.GetByID())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "seller not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 3: bad request - Invalid Id", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		rt := chi.NewRouter()
		rt.Get("/sellers/{id}", sellerHd.GetByID())

		//Act
		req, res := httptest.NewRequest(http.MethodGet, "/sellers/Xyz", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
}

func TestSellerHandler_Create(t *testing.T) {
	t.Run("case 1: seller create successfully", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Create", mock.Anything).Return(seller, nil)

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sellers", sellerHd.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"cid":"1234","company_name":"Company X","address":"221B Baker Street","telephone":"123590","locality_id":"1"}}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 2: bad request - invalid body", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Create", mock.Anything).Return(model.Seller{}, eh.GetErrInvalidData(eh.SELLER))

		body, err := json.Marshal(dto.SellerRequestDTO{
			CompanyID:   new(string),
			CompanyName: new(string),
			Address:     new(string),
			Telephone:   new(string),
			LocalityId:  new(string),
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sellers", sellerHd.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Unprocessable Entity","message": "invalid data: seller"}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 3: conflict - invalid locality id", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Create", mock.Anything).Return(model.Seller{}, eh.GetErrForeignKey(eh.LOCALITY))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sellers", sellerHd.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Conflict","message": "locality foreign key not found"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 4: internal server error", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Create", mock.Anything).Return(model.Seller{}, eh.GetErrDatabase(eh.SELLER))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sellers", sellerHd.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: seller"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 5: conflict - invalid company ID", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Create", mock.Anything).Return(model.Seller{}, eh.GetErrAlreadyExists(eh.SELLER))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sellers", sellerHd.Create())

		//Act
		req, res := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Conflict","message": "seller already exists"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
}

func TestSellerHandler_Update(t *testing.T) {
	t.Run("case 1: seller update successfully", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", 1, mock.Anything).Return(seller, nil)

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/1", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"cid":"1234","company_name":"Company X","address":"221B Baker Street","telephone":"123590","locality_id":"1"}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 2: not found - seller ID doesn't exist", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", 2, mock.Anything).Return(model.Seller{}, eh.GetErrNotFound(eh.SELLER))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/2", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "seller not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 3: bad request - invalid id", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", "Xyz", mock.Anything).Return(model.Seller{}, eh.INVALID_ID)

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/Xyz", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertNotCalled(t, "Update")
	})
	t.Run("case 4: conflict - invalid Company ID", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", 1, mock.Anything).Return(model.Seller{}, eh.GetErrAlreadyExists(eh.SELLER))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Conflict","message": "seller already exists"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})	
	t.Run("case 5: conflict - foreign key - seller id doesn't exist", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", 1, mock.Anything).Return(model.Seller{}, eh.GetErrForeignKey(eh.LOCALITY))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Conflict","message": "locality foreign key not found"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 6: internal server error", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		sellerSv.On("Update", 1, mock.Anything).Return(model.Seller{}, eh.GetErrDatabase(eh.SELLER))

		body, err := json.Marshal(sellerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sellers/{id}", sellerHd.Update())

		//Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sellers/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		//Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: seller"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
}

func TestSellerHandler_Delete(t *testing.T) {
	t.Run("case 1: delete seller successfully", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)
		sellerSv.On("Delete", 1).Return(nil)

		rt := chi.NewRouter()
		rt.Delete("/sellers/{id}", sellerHd.Delete())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sellers/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 2: not found - Seller ID doesn't exist", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)
		sellerSv.On("Delete", 2).Return(eh.GetErrNotFound(eh.SELLER))

		rt := chi.NewRouter()
		rt.Delete("/sellers/{id}", sellerHd.Delete())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sellers/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "seller not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
	t.Run("case 3: bad request - Invalid Id", func(t *testing.T) {
		//Arrange
		sellerSv := service.NewSellerServiceMock()
		sellerHd := NewSellerHandler(sellerSv)

		rt := chi.NewRouter()
		rt.Delete("/sellers/{id}", sellerHd.Delete())

		//Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sellers/Xyz", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		sellerSv.AssertExpectations(t)
	})
}
