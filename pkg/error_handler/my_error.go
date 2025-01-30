package error_handler

import (
	"errors"
	"fmt"
)

const (
	BUYER        = "buyer"
	EMPLOYEE     = "employee"
	PRODUCT      = "product"
	SECTION      = "section"
	SELLER       = "seller"
	WAREHOUSE    = "warehouse"
	ID           = "id"
	INVALID_BODY = "invalid request body"
	INVALID_ID   = "invalid id"
	CARD_NUMBER  = "card number ID"
	WAREHOUSE_CODE = "warehouse code"
)

var (
	ErrNotFound      = errors.New("not found")      // 404
	ErrAlreadyExists = errors.New("already exists") // 409
	ErrInvalidData   = errors.New("invalid data")   // 422
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

func HandleError(err error) (int, string) {
	if errors.Is(err, ErrNotFound) {
		return 404, err.Error()
	}

	if errors.Is(err, ErrAlreadyExists) {
		return 409, err.Error()
	}

	if errors.Is(err, ErrInvalidData) {
		return 422, err.Error()
	}

	return 500, err.Error()
}
