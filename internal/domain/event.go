package domain

import "time"

// Event в честь мого відчислення

type Event struct {
	Id          uint64
	Device      uint64
	RoomId      *uint64
	Action      bool
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
