package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventDto struct {
	Id       uint64  `json:"id"`
	DeviceId uint64  `json:"device_id"`
	RoomId   *uint64 `json:"room_id"`
	Action   bool    `json:"action"`
}

func (d EventDto) DomainToDto(ev domain.Event) EventDto {
	return EventDto{
		Id:       ev.Id,
		DeviceId: ev.DeviceId,
		RoomId:   ev.RoomId,
		Action:   ev.Action,
	}
}
