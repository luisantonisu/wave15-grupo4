package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

var (
	ErrWarehouseCodeEmpty           = errors.New("warehouse code is empty")
	ErrTelephoneEmptyOrInvalid      = errors.New("telephone is empty or invalid")
	ErrMinimumCapacityNegative      = errors.New("minimum capacity is negative")
	ErrMinimumTemperatureOutOfRange = errors.New("minimum temperature is out of range")
)

func NewWarehouseService(rp repository.IWarehouse) *WarehouseService {
	return &WarehouseService{rp: rp}
}

type WarehouseService struct {
	rp repository.IWarehouse
}

func (ws *WarehouseService) GetAll() (map[int]model.Warehouse, error) {
	return ws.rp.GetAll()
}

func (ws *WarehouseService) GetByID(id int) (model.Warehouse, error) {
	return ws.rp.GetByID(id)
}

func (ws *WarehouseService) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	if err := validateCode(warehouse.WarehouseCode); err != nil {
		return model.Warehouse{}, err
	}

	if warehouse.Telephone != 0 {
		if err := validatePhone(warehouse.Telephone); err != nil {
			return model.Warehouse{}, err
		}
	}

	if warehouse.MinimumCapacity != 0 {
		if err := validateCapacity(warehouse.MinimumCapacity); err != nil {
			return model.Warehouse{}, err
		}
	}

	if warehouse.MinimumTemperature != 0 {
		if err := validateTemperature(warehouse.MinimumTemperature); err != nil {
			return model.Warehouse{}, err
		}
	}
	return ws.rp.Create(warehouse)
}

func (ws *WarehouseService) Update(id int, warehouse model.WarehouseAttributesPtr) (model.Warehouse, error) {
	if warehouse.WarehouseCode != nil {
		if err := validateCode(*warehouse.WarehouseCode); err != nil {
			return model.Warehouse{}, err
		}
	}

	if warehouse.Telephone != nil {
		if err := validatePhone(*warehouse.Telephone); err != nil {
			return model.Warehouse{}, err
		}
	}

	if warehouse.MinimumCapacity != nil {
		if err := validateCapacity(*warehouse.MinimumCapacity); err != nil {
			return model.Warehouse{}, err
		}
	}

	if warehouse.MinimumTemperature != nil {
		if err := validateTemperature(*warehouse.MinimumTemperature); err != nil {
			return model.Warehouse{}, err
		}
	}
	return ws.rp.Update(id, warehouse)
}

func (ws *WarehouseService) Delete(id int) error {
	return ws.rp.Delete(id)
}

func validateCode(code string) error {
	if code == "" {
		return eh.GetErrInvalidData(ErrWarehouseCodeEmpty.Error())
	}
	return nil
}

func validatePhone(phone uint) error {
	if phone < 1000000 || phone > 9999999999 {
		return eh.GetErrInvalidData(ErrTelephoneEmptyOrInvalid.Error())
	}
	return nil
}

func validateCapacity(minimumCapacity int) error {
	if minimumCapacity < 0 {
		return eh.GetErrInvalidData(ErrMinimumCapacityNegative.Error())
	}
	return nil
}

func validateTemperature(temperature float32) error {
	if temperature < -20 || temperature > 40 {
		return eh.GetErrInvalidData(ErrMinimumTemperatureOutOfRange.Error())
	}

	return nil
}
