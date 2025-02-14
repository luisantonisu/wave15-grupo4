package error_handler

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	BUYER          = "buyer"
	EMPLOYEE       = "employee"
	PRODUCT        = "product"
	PRODUCT_RECORD = "product record"
	SECTION        = "section"
	SELLER         = "seller"
	WAREHOUSE      = "warehouse"
	ID             = "id"
	INVALID_BODY   = "invalid request body"
	INVALID_ID     = "invalid id"
	CARD_NUMBER    = "card number ID"
	WAREHOUSE_CODE = "warehouse code"
	SECTION_NUMBER = "section number"
	ORDER_NUMBER   = "order number"
	INBOUND_ORDER  = "inbound order"
	PURCHASE_ORDER = "purchase order"
	COMPANY_ID     = "company ID"
	LOCALITY       = "locality"
	COUNTRY_NAME   = "country name"
	PROVINCE_NAME  = "province name"
	COUNTRY_ID     = "country ID"
	PROVINCE_ID    = "province ID"
	LOCALITY_ID    = "locality ID"
	CARRY          = "carry"
)

var (
	ErrNotFound      = errors.New("not found")             // 404
	ErrAlreadyExists = errors.New("already exists")        // 409
	ErrInvalidData   = errors.New("invalid data")          // 422
	ErrForeignKey    = errors.New("foreign key not found") // 409
	ErrGettingData   = errors.New("error getting data")    // 500
	ErrParsingData   = errors.New("error parsing data")    // 500
)

func GetErrNotFound(entity string) error {
	return fmt.Errorf("%s %w", entity, ErrNotFound)
}

func GetErrAlreadyExists(entity string) error {
	return fmt.Errorf("%s %w", entity, ErrAlreadyExists)
}

func GetErrInvalidData(entity string) error {
	return fmt.Errorf("%w: %s", ErrInvalidData, entity)
}

func GetErrForeignKey(entity string) error {
	return fmt.Errorf("%s %w", entity, ErrForeignKey)
}

func GetErrGettingData(entity string) error {
	return fmt.Errorf("%w: %s", ErrGettingData, entity)
}

func GetErrParsingData(entity string) error {
	return fmt.Errorf("%w: %s", ErrParsingData, entity)
}

func HandleError(err error) (int, string) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, err.Error()
	}

	if errors.Is(err, ErrAlreadyExists) {
		return http.StatusConflict, err.Error()
	}

	if errors.Is(err, ErrInvalidData) {
		return http.StatusUnprocessableEntity, err.Error()
	}

	if errors.Is(err, ErrForeignKey) {
		return http.StatusConflict, err.Error()
	}

	return http.StatusInternalServerError, err.Error()
}
