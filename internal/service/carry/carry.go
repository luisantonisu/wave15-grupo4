package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
)

func NewCarryService(rp repository.ICarry) *CarryService {
	return &CarryService{rp: rp}
}

type CarryService struct {
	rp repository.ICarry
}

func (s *CarryService) Create(carry model.Carry) (model.Carry, error) {
	return s.rp.Create(carry)
}
