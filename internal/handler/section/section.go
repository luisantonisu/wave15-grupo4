package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/section"
)

func NewSectionHandler(sv service.ISection) *SectionHandler {
	return &SectionHandler{sv: sv}
}

type SectionHandler struct {
	sv service.ISection
}