package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func LocalityRequestDTOToLocality(localityDto dto.LocalityRequestDTO) model.Locality {
	data := model.Locality{
		Id: localityDto.Data.Id,
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: localityDto.Data.LocalityName,
			ProvinceName: localityDto.Data.ProvinceName,
			CountryName:  localityDto.Data.CountryName,
		},
	}

	return data
}

func LocalityToLocalityDataResponseDTO(locality model.Locality) dto.LocalityDataResponseDTO {
	data := dto.LocalityDataResponseDTO{
		Id : locality.Id,
		LocalityName: locality.LocalityName,
		ProvinceName: locality.ProvinceName,
		CountryName: locality.CountryName,
	}

	return data
}
