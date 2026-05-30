package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventRequest struct {
	RoomId *uint64 `json:"room_id"`
	Action bool    `json:"action"`
}

func (e EventRequest) ToDomainModel() (interface{}, error) {
	return domain.Event{
		RoomId: e.RoomId,
		Action: e.Action,
	}, nil
}
