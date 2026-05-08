package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id             uint64     `db:"id,omitempty"`
	OrganizationId uint64     `db:"organizationId"`
	Name           string     `db:"name"`
	Description    *string    `db:"description"`
	CreatedDate    time.Time  `db:"createdDate"`
	UpdatedDate    time.Time  `db:"updatedDate"`
	DeletedDate    *time.Time `db:"deletedDate"`
}

type RoomRepository interface {
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

func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(rm room) domain.Room {
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
