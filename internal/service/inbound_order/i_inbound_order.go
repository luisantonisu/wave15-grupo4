package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IInboundOrder interface {
	Create(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error)
}
