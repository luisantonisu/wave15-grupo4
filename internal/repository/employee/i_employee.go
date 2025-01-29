package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IEmployee interface {
	GetAll() (map[int]model.Employee, error)
	GetByID(id int) (model.Employee, error)
	Create(employee model.Employee) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Delete(id int) error
}
