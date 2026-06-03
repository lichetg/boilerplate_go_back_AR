package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"time"
)

type EventDto struct {
	Id          uint64    `json:"id"`
	DeviceId    uint64    `json:"device_id"`
	RoomId      *uint64   `json:"room_id"`
	Action      bool      `json:"action"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Events struct {
	Items []domain.Event `json:"items"`
	Total uint64         `json:"total"`
	Pages uint64         `json:"pages"`
}

func (d EventDto) DomainToDto(ev domain.Event) EventDto {
	return EventDto{
		Id:          ev.Id,
		DeviceId:    ev.DeviceId,
		RoomId:      ev.RoomId,
		Action:      ev.Action,
		UpdatedDate: ev.CreatedDate,
	}
}

func (d EventDto) EventDomainToDtoCollection(eves []domain.Event) []EventDto {
	evesDto := make([]EventDto, len(eves))
	for i, _ := range eves {
		evesDto[i] = d.DomainToDto(eves[i])
	}
	return evesDto
}
