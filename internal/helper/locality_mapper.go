package helper

import (
	"strconv"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func LocalityRequestDTOToLocality(localityDto dto.LocalityRequestDTO) model.Locality {
	data := model.Locality{
		Id: *localityDto.Data.Id,
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: localityDto.Data.LocalityName,
			ProvinceName: localityDto.Data.ProvinceName,
			CountryName:  localityDto.Data.CountryName,
		},
	}

	return data
}

func LocalityToLocalityDataResponseDTO(locality model.Locality) dto.LocalityDataDTO {
	data := dto.LocalityDataDTO{
		Id:           &locality.Id,
		LocalityName: locality.LocalityName,
		ProvinceName: locality.ProvinceName,
		CountryName:  locality.CountryName,
	}

	return data
}

func CarriersByLocalityReportToCarriersByLocalityReportResponseDTO(carriersReport model.CarriersByLocalityReport) dto.CarriersByLocalityReportResponseDTO {
	data := dto.CarriersByLocalityReportResponseDTO{
		LocalityID:    carriersReport.LocalityID,
		LocalityName:  carriersReport.LocalityName,
		CarriersCount: carriersReport.CarriersCount,
	}

	return data
}

func LocalityReportToLocalityReportResponseDto(locality model.LocalityReport) dto.LocalityReportResponseDTO {
	data := dto.LocalityReportResponseDTO{
		Id:           strconv.Itoa(locality.Id),
		LocalityName: locality.LocalityName,
		SellerCount:  locality.SellerCount,
	}

	return data
}
