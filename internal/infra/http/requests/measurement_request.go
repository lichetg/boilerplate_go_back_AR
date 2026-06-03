package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementRequest struct {
	RoomId *uint64 `json:"room_id"`
	Value  float64 `json:"value"`
}

type MeasurementListRequest struct {
	Page         uint `json:"page"`
	CountPerPage uint `json:"count_per_page"`

	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`

	Sort string `json:"sort"`
}

func (m MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		RoomId: m.RoomId,
		Value:  m.Value,
	}, nil
}
