package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementDto struct {
	Id       uint64  `json:"id"`
	DeviceId uint64  `json:"deviceId"`
	RoomId   *uint64 `json:"roomId"`
	Value    float64 `json:"value"`
}

func (d MeasurementDto) DomainToDto(ms domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:       ms.Id,
		DeviceId: ms.DeviceId,
		RoomId:   ms.RoomId,
		Value:    ms.Value,
	}
}
