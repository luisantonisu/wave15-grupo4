package model

type Buyer struct {
	ID int
	BuyerAttributes
}
type BuyerAttributes struct {
	CardNumberId *string
	FirstName    *string
	LastName     *string
}

type ReportPurchaseOrders struct {
	ID                  int    `json:"id"`
	CardNumberId        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}
