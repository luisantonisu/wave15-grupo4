package repository


type IProvince interface {
	GetProvinceID(countryId int, provinceName string) (int, error)
}