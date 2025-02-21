package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISection interface {
	GetAll() ([]model.Section, error)
	GetByID(id int) (model.Section, error)
	Create(section model.SectionAttributes) (model.Section, error)
	Patch(id int, section model.SectionAttributes) (model.Section, error)
	Delete(id int) error
	Report(id *int) ([]model.ReportProductsBatches, error)
}
