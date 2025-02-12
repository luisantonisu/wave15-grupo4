package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type InbounderOrderRepository struct {
	db *sql.DB
}

func NewInboundOrderRepository(defaultDB *sql.DB) *InbounderOrderRepository {
	return &InbounderOrderRepository{
		db: defaultDB,
	}
}

func (i *InbounderOrderRepository) AlreadyExits(atribute string, value int) bool {
	var exists bool
	err := i.db.QueryRow("SELECT EXISTS(SELECT 1 FROM inbound_orders WHERE ? = ?)", atribute, value).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (i *InbounderOrderRepository) CreateInboundOrder(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error) {
	if i.AlreadyExits("employee_id", inboundOrder.EmployeeID) {
		return model.InboundOrder{}, eh.GetErrAlreadyExists(eh.EMPLOYEE)
	}

	if i.AlreadyExits("order_id", inboundOrder.OrderNumber) {
		return model.InboundOrder{}, eh.GetErrAlreadyExists(eh.ORDER_NUMBER)
	}

	row, err := i.db.Exec("INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?, ?, ?, ?, ?)", inboundOrder.OrderDate, inboundOrder.OrderNumber, inboundOrder.EmployeeID, inboundOrder.ProductBatchID, inboundOrder.WarehouseID)
	if err != nil {
		return model.InboundOrder{}, eh.GetErrInvalidData(eh.INBOUND_ORDER)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.InboundOrder{}, err
	}

	var inb model.InboundOrder
	inb.ID = int(id)
	inb.InboundOrderAttributes = inboundOrder
	return inb, nil
}
