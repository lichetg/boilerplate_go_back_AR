package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"time"
)

type MeasurementDto struct {
	Id          uint64    `json:"id"`
	DeviceId    uint64    `json:"deviceId"`
	RoomId      *uint64   `json:"roomId"`
	Value       float64   `json:"value"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func (d MeasurementDto) DomainToDto(ms domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:          ms.Id,
		DeviceId:    ms.DeviceId,
		RoomId:      ms.RoomId,
		Value:       ms.Value,
		UpdatedDate: ms.CreatedDate,
	}
}

func (d MeasurementDto) MeasurementDomainToDtoCollection(meass []domain.Measurement) []MeasurementDto {
	meassDto := make([]MeasurementDto, len(meass))
	for i, _ := range meass {
		meassDto[i] = d.DomainToDto(meass[i])
	}
	return meassDto
}
