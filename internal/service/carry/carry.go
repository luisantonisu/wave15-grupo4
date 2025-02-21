package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	carryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"

	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

const (
	ErrCarryIDEmpty = "carry id is empty"
)

func NewCarryService(rp carryRepository.ICarry, localityRp localityRepository.ILocality) *CarryService {
	return &CarryService{
		carryRp:    rp,
		localityRp: localityRp,
	}
}

type CarryService struct {
	carryRp    carryRepository.ICarry
	localityRp localityRepository.ILocality
}

func (s *CarryService) Create(carry model.Carry) (model.Carry, error) {
	if carry.CarryID == nil || *carry.CarryID == "" {
		return model.Carry{}, eh.GetErrInvalidData(ErrCarryIDEmpty)
	}

	exists, err := s.carryRp.GetByCarryID(*carry.CarryID)
	if exists != (model.Carry{}) {
		return model.Carry{}, eh.GetErrAlreadyExists(eh.CARRY_ID)
	}
	if err != nil && !errors.Is(err, eh.ErrNotFound) {
		return model.Carry{}, err
	}

	if carry.LocalityID != nil {
		if _, err := s.localityRp.GetByID(*carry.LocalityID); err != nil {
			return model.Carry{}, eh.GetErrForeignKey(eh.LOCALITY)
		}
	}
	return s.carryRp.Create(carry)
}
