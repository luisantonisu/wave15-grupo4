package service

import (
	"regexp"
	"strconv"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	countryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/country"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	ProvinceRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/province"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewLocalityService(countryRp countryRepository.ICountry, provinceRp ProvinceRepository.IProvince, localityRp localityRepository.ILocality) *LocalityService {
	return &LocalityService{
		countryRp:  countryRp,
		provinceRp: provinceRp,
		localityRp: localityRp,
	}
}

type LocalityService struct {
	countryRp  countryRepository.ICountry
	provinceRp ProvinceRepository.IProvince
	localityRp localityRepository.ILocality
}

func (s *LocalityService) Create(locality model.Locality) (model.Locality, error) {
	//validate locality
	err := s.validateLocality(locality)
	if err != nil {
		return model.Locality{}, err
	}

	//validate CountryName and get countryID
	countryId, err := s.countryRp.GetCountryIDByCountryName(locality.CountryName)
	if err != nil {
		return model.Locality{}, err
	}

	//validate ProvinceName and get ProvinceID
	provinceID, err := s.provinceRp.GetProvinceID(countryId, locality.ProvinceName)
	if err != nil {
		return model.Locality{}, err
	}

	//Convert localityID
	localityID, err := strconv.Atoi(locality.Id)
	if err != nil {
		return model.Locality{}, err
	}

	//create model localityDb
	var localityDb = model.LocalityDBModel{
		Id:           localityID,
		LocalityName: locality.LocalityName,
		ProvinceID:   provinceID,
	}

	newLocality, err := s.localityRp.Create(localityDb)
	if err != nil {
		return model.Locality{}, err
	}

	var respLocality = model.Locality{
		Id: strconv.Itoa(newLocality.Id),
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: newLocality.LocalityName,
			ProvinceName: locality.ProvinceName,
			CountryName:  locality.CountryName,
		},
	}

	return respLocality, nil
}

func (s *LocalityService) SellersReport(id *int) (map[int]model.LocalityReport, error) {
	return s.localityRp.SellersReport(id)
}

func (s *LocalityService) validateLocality(locality model.Locality) error {
	//validate if locality_Id only contains numbers and is not empty
	pattern := regexp.MustCompile("^[0-9]+$")
	match := pattern.MatchString(locality.Id)
	if !match {
		return eh.GetErrInvalidData(eh.LOCALITY)
	}

	hasLocalityName := locality.LocalityName != ""
	hasProvinceName := locality.ProvinceName != ""
	hasCountryName := locality.CountryName != ""

	var validators = []bool{hasLocalityName, hasProvinceName, hasCountryName}
	for _, item := range validators {
		if !item {
			return eh.GetErrInvalidData(eh.LOCALITY)
		}
	}

	return nil

}
