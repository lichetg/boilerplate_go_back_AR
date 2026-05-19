package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type Room struct {
	Id             uint64     `db:"id,omitempty"`
	OrganizationId uint64     `db:"organization_id"`
	Name           string     `db:"name"`
	Description    *string    `db:"description"`
	CreatedDate    time.Time  `db:"created_date"`
	UpdatedDate    time.Time  `db:"updated_date"`
	DeletedDate    *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	Save(o domain.Room) (domain.Room, error)
	FindByOrgId(oId uint64) ([]domain.Room, error)
	Find(id uint64) (domain.Room, error)
	Update(o domain.Room) (domain.Room, error)
	Delete(Id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(session db.Session) roomRepository {
	return roomRepository{
		coll: session.Collection(RoomsTableName),
		sess: session,
	}
}

func (r roomRepository) Save(o domain.Room) (domain.Room, error) {
	org := r.mapDomainToModel(o)
	now := time.Now()
	org.CreatedDate = now
	org.UpdatedDate = now

	err := r.coll.InsertReturning(&org)
	if err != nil {
		return domain.Room{}, err
	}

	o = r.mapModelToDomain(org)
	return o, nil
}

func (r roomRepository) FindByOrgId(oId uint64) ([]domain.Room, error) {
	var rooms []Room

	err := r.coll.Find(db.Cond{
		"organization_id": oId,
		"deleted_date":    nil,
	}).All(&rooms)

	if err != nil {
		return nil, err
	}

	rms := r.mapModelToDomainCollection(rooms)
	return rms, nil
}

func (r roomRepository) Find(id uint64) (domain.Room, error) {
	var room Room

	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&room)
	if err != nil {
		return domain.Room{}, err
	}

	o := r.mapModelToDomain(room)
	return o, nil
}

func (r roomRepository) Update(o domain.Room) (domain.Room, error) {
	room := r.mapDomainToModel(o)
	room.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&room)
	if err != nil {
		return domain.Room{}, err
	}

	o = r.mapModelToDomain(room)
	return o, nil
}

func (r roomRepository) Delete(Id uint64) error {
	return r.coll.Find(db.Cond{"Id": Id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapDomainToModel(rm domain.Room) Room {
	return Room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(rm Room) domain.Room {
	return domain.Room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(rooms []Room) []domain.Room {
	rms := make([]domain.Room, len(rooms))
	for i := range rooms {
		rms[i] = r.mapModelToDomain(rooms[i])
	}
	return rms
}
