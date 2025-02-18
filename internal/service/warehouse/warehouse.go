package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

const (
	ErrWarehouseCodeEmpty           = "warehouse code is empty"
	ErrTelephoneEmptyOrInvalid      = "telephone is empty or invalid"
	ErrMinimumCapacityNegative      = "minimum capacity is negative"
	ErrMinimumTemperatureOutOfRange = "minimum temperature is out of range"
)

func NewWarehouseService(warehouseRp warehouseRepository.IWarehouse, localityRp localityRepository.ILocality) *WarehouseService {
	return &WarehouseService{
		warehouseRp: warehouseRp,
		localityRp:  localityRp,
	}
}

type WarehouseService struct {
	warehouseRp warehouseRepository.IWarehouse
	localityRp  localityRepository.ILocality
}

func (s *WarehouseService) GetAll() ([]model.Warehouse, error) {
	return s.warehouseRp.GetAll()
}

func (s *WarehouseService) GetByID(id int) (model.Warehouse, error) {
	return s.warehouseRp.GetByID(id)
}

func (s *WarehouseService) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	if warehouse.WarehouseCode == nil || *warehouse.WarehouseCode == "" {
		return model.Warehouse{}, eh.GetErrInvalidData(ErrWarehouseCodeEmpty)
	}

	// Check if warehouse code already exists
	exists, err := s.warehouseRp.GetByCode(*warehouse.WarehouseCode)
	if exists != (model.Warehouse{}) {
		return model.Warehouse{}, eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
	}
	if err != nil && !errors.Is(err, eh.ErrNotFound) {
		return model.Warehouse{}, err
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

	if warehouse.LocalityID != nil {
		if _, err := s.localityRp.GetByID(*warehouse.LocalityID); err != nil {
			return model.Warehouse{}, eh.GetErrForeignKey(eh.LOCALITY)
		}
	}

	return s.warehouseRp.Create(warehouse)
}

func (s *WarehouseService) Update(id int, warehouse model.WarehouseAttributes) (model.Warehouse, error) {
	existingWarehouse, err := s.warehouseRp.GetByID(id)
	if err != nil {
		return model.Warehouse{}, err
	}

	if warehouse.WarehouseCode != nil && *warehouse.WarehouseCode != *existingWarehouse.WarehouseCode {
		// Check if warehouse code already exists
		exists, err := s.warehouseRp.GetByCode(*warehouse.WarehouseCode)
		if exists != (model.Warehouse{}) {
			return model.Warehouse{}, eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
		}
		if err != nil && !errors.Is(err, eh.ErrNotFound) {
			return model.Warehouse{}, err
		}
		existingWarehouse.WarehouseCode = warehouse.WarehouseCode
	}

	if warehouse.Telephone != nil {
		if err := validatePhone(*warehouse.Telephone); err != nil {
			return model.Warehouse{}, err
		}
		existingWarehouse.Telephone = warehouse.Telephone
	}

	if warehouse.MinimumCapacity != nil {
		if err := validateCapacity(*warehouse.MinimumCapacity); err != nil {
			return model.Warehouse{}, err
		}
		existingWarehouse.MinimumCapacity = warehouse.MinimumCapacity
	}

	if warehouse.MinimumTemperature != nil {
		if err := validateTemperature(*warehouse.MinimumTemperature); err != nil {
			return model.Warehouse{}, err
		}
		existingWarehouse.MinimumTemperature = warehouse.MinimumTemperature
	}

	if warehouse.LocalityID != nil {
		if _, err := s.localityRp.GetByID(*warehouse.LocalityID); err != nil {
			return model.Warehouse{}, eh.GetErrForeignKey(eh.LOCALITY)
		}
	}

	return s.warehouseRp.Update(id, existingWarehouse)
}

func (s *WarehouseService) Delete(id int) error {
	if _, err := s.warehouseRp.GetByID(id); err != nil {
		return eh.GetErrNotFound(eh.WAREHOUSE)
	}
	return s.warehouseRp.Delete(id)
}

func validatePhone(phone uint) error {
	if phone < 1000000 || phone > 9999999999 {
		return eh.GetErrInvalidData(ErrTelephoneEmptyOrInvalid)
	}
	return nil
}

func validateCapacity(capacity int) error {
	if capacity < 0 {
		return eh.GetErrInvalidData(ErrMinimumCapacityNegative)
	}
	return nil
}

func validateTemperature(temperature float32) error {
	if temperature < -20 || temperature > 40 {
		return eh.GetErrInvalidData(ErrMinimumTemperatureOutOfRange)
	}
	return nil
}
