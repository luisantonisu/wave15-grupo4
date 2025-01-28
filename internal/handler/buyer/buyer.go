package handler

import (
	"net/http"

	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
)

func NewBuyerHandler(sv service.IBuyer) *BuyerHandler {
	return &BuyerHandler{sv: sv}
}

type BuyerHandler struct {
	sv service.IBuyer
}

func (b *BuyerHandler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}
