package domain

import "time"

type Event struct {
	Id          uint64
	DeviceId    uint64
	RoomId      *uint64
	Action      bool
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
