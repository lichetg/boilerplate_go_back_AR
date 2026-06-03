package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/upper/db/v4"
	"math"
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
	FindList(p domain.Pagination, f domain.EventFilters) (resources.Events, error)
	Update(o domain.Event) (domain.Event, error)
	Delete(Id uint64) error
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

func (r eventRepository) FindList(
	p domain.Pagination,
	f domain.EventFilters,
) (resources.Events, error) {

	var models []Event

	if p.Page == 0 {
		p.Page = 1
	}

	if p.CountPerPage == 0 {
		p.CountPerPage = 20
	}

	query := r.coll.Find(
		db.Cond{
			"device_id":    f.DeviceId,
			"deleted_date": nil,
		},
	)

	if f.CreatedDateFrom != nil {
		query = query.And("created_date >= ?", *f.CreatedDateFrom)
	}

	if f.CreatedDateTo != nil {
		query = query.And("created_date <= ?", *f.CreatedDateTo)
	}

	switch f.Sort {

	case "created_date":
		query = query.OrderBy("created_date")

	case "-created_date":
		query = query.OrderBy("-created_date")

	case "action":
		query = query.OrderBy("action")

	case "-action":
		query = query.OrderBy("-action")

	default:
		query = query.OrderBy("-created_date")
	}

	result := query.Paginate(uint(p.CountPerPage))

	err := result.Page(uint(p.Page)).All(&models)
	if err != nil {
		return resources.Events{}, err
	}

	total, err := result.TotalEntries()
	if err != nil {
		return resources.Events{}, err
	}

	items := make([]domain.Event, len(models))

	for i, m := range models {
		items[i] = r.mapModelToDomain(m)
	}

	return resources.Events{
		Items: items,
		Total: total,
		Pages: uint64(math.Ceil(
			float64(total) / float64(p.CountPerPage),
		)),
	}, nil
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
