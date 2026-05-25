package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementRequest struct {
	RoomId *uint64 `json:"room_id"`
	Value  float64 `json:"value"`
}

func (m MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		RoomId: m.RoomId,
		Value:  m.Value,
	}, nil
}
