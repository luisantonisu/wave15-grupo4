package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISection interface {
	GetAll() (map[int]model.Section, error)
	GetByID(id int) (model.Section, error)
	Create(section model.Section) (model.Section, error)
	Delete(id int) error
}
