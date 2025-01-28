package model

type Employee struct {
	Id int
	EmployeeAttributes
}

type EmployeeAttributes struct {
	CardNumberId int
	FirstName    string
	LastName     string
	WarehouseId  int
}
