package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewOrderStatusRepository(db *sql.DB) *OrderStatusRepository {
	return &OrderStatusRepository{
		db: db,
	}
}

type OrderStatusRepository struct {
	db *sql.DB
}

// Get order status by id
func (r *OrderStatusRepository) GetByID(id int) (model.OrderStatus, error) {
	// Create response entity
	var orderStatus model.OrderStatus

	// Get order status from db
	err := r.db.QueryRow("SELECT id, description FROM order_status WHERE id = ?", id).Scan(
		&orderStatus.ID, &orderStatus.Description,
	)
	if err != nil {
		return model.OrderStatus{}, eh.GetErrNotFound(eh.ORDER_STATUS)
	}
	return orderStatus, nil
}
