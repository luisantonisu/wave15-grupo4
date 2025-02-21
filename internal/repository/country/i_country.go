package repository

type ICountry interface {
	GetCountryIDByCountryName(countryName string) (int, error)
}