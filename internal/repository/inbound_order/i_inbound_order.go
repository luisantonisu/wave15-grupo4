package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IInboundOrder interface {
	CreateInboundOrder(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error)
}
