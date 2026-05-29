package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const EventTableName = "events"

type Event struct {
	Id          uint64     `db:"id"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      *uint64    `db:"room_id"`
	Action      bool       `db:"action"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type EventRepository interface {
}

type eventRepository struct {
	coll db.Collection
	sess db.Session
}

func NewEventRepository(session db.Session) eventRepository {
	return eventRepository{
		coll: session.Collection(EventTableName),
		sess: session,
	}
}
func (r eventRepository) Save(o domain.Event) (domain.Event, error) {
	eve := r.mapDomainToModel(o)
	now := time.Now()
	eve.CreatedDate = now
	eve.UpdatedDate = now

	err := r.coll.InsertReturning(&eve)
	if err != nil {
		return domain.Event{}, err
	}

	o = r.mapModelToDomain(eve)
	return o, nil
}

func (r eventRepository) mapDomainToModel(eve domain.Event) Event {
	return Event{
		Id:          eve.Id,
		DeviceId:    eve.DeviceId,
		RoomId:      eve.RoomId,
		Action:      eve.Action,
		CreatedDate: eve.CreatedDate,
		UpdatedDate: eve.UpdatedDate,
		DeletedDate: eve.DeletedDate,
	}
}
func (r eventRepository) mapModelToDomain(eve Event) domain.Event {
	return domain.Event{
		Id:          eve.Id,
		DeviceId:    eve.DeviceId,
		RoomId:      eve.RoomId,
		Action:      eve.Action,
		CreatedDate: eve.CreatedDate,
		UpdatedDate: eve.UpdatedDate,
		DeletedDate: eve.DeletedDate,
	}
}
