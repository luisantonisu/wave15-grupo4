package model

type Employee struct {
	ID int
	EmployeeAttributes
}

type EmployeeAttributes struct {
	CardNumberID int
	FirstName    string
	LastName     string
	WarehouseID  int
}

type EmployeeAttributesPtr struct {
	CardNumberID *int
	FirstName    *string
	LastName     *string
	WarehouseID  *int
}
