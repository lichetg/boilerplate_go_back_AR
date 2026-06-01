package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const EventTableName = "events"

type Event struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      *uint64    `db:"room_id"`
	Action      bool       `db:"action"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type EventRepository interface {
	Save(o domain.Event) (domain.Event, error)
	Find(id uint64) (domain.Event, error)
	Update(o domain.Event) (domain.Event, error)
	Delete(Id uint64) error
	FindByDeviceAndPeriod(id uint64, from time.Time, to time.Time) ([]domain.Event, error)
	FindByDeviceId(devId uint64) ([]domain.Event, error)
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

func (r eventRepository) Find(id uint64) (domain.Event, error) {
	var event Event

	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&event)
	if err != nil {
		return domain.Event{}, err
	}

	o := r.mapModelToDomain(event)
	return o, nil
}

func (r eventRepository) FindByDeviceId(devId uint64) ([]domain.Event, error) {
	var eves []Event

	err := r.coll.Find(db.Cond{
		"device_id":    devId,
		"deleted_date": nil,
	}).All(&eves)

	if err != nil {
		return nil, err
	}

	o := r.mapModelToDomainCollection(eves)
	return o, nil
}

func (r eventRepository) Update(o domain.Event) (domain.Event, error) {
	eve := r.mapDomainToModel(o)
	eve.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&eve)
	if err != nil {
		return domain.Event{}, err
	}

	o = r.mapModelToDomain(eve)
	return o, nil
}

func (r eventRepository) Delete(Id uint64) error {
	return r.coll.Find(db.Cond{"id": Id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r eventRepository) FindByDeviceAndPeriod(
	id uint64,
	from time.Time,
	to time.Time,
) ([]domain.Event, error) {
	var eves []Event

	err := r.coll.Find(
		db.Cond{
			"device_id":       id,
			"created_date >=": from,
			"created_date <=": to,
			"deleted_date":    nil,
		},
	).All(&eves)

	if err != nil {
		return nil, err
	}

	events := make([]domain.Event, len(eves))

	for i, m := range eves {
		events[i] = r.mapModelToDomain(m)
	}

	return events, nil
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

func (r eventRepository) mapModelToDomainCollection(eves []Event) []domain.Event {
	evs := make([]domain.Event, len(eves))
	for i := range eves {
		evs[i] = r.mapModelToDomain(eves[i])
	}
	return evs
}
