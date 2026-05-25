package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const MeasurementsTableName = "measurements"

type measurement struct {
	Id          uint64     `db:"id"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      *uint64    `db:"room_id"`
	Value       float64    `db:"value"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}
type MeasurementRepository interface {
	Save(o domain.Measurement) (domain.Measurement, error)
	Find(id uint64) (domain.Measurement, error)
}
type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

func NewMeasurementRepository(session db.Session) measurementRepository {
	return measurementRepository{
		coll: session.Collection(MeasurementsTableName),
		sess: session,
	}
}

func (r measurementRepository) Save(o domain.Measurement) (domain.Measurement, error) {
	meas := r.mapDomainToModel(o)
	now := time.Now()
	meas.CreatedDate = now
	meas.UpdatedDate = now

	err := r.coll.InsertReturning(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	o = r.mapModelToDomain(meas)
	return o, nil
}

func (r measurementRepository) Find(id uint64) (domain.Measurement, error) {
	var meas measurement

	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	o := r.mapModelToDomain(meas)
	return o, nil
}

func (r measurementRepository) mapDomainToModel(rm domain.Measurement) measurement {
	return measurement{
		Id:          rm.Id,
		DeviceId:    rm.DeviceId,
		RoomId:      rm.RoomId,
		Value:       rm.Value,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}

func (r measurementRepository) mapModelToDomain(rm measurement) domain.Measurement {
	return domain.Measurement{
		Id:          rm.Id,
		DeviceId:    rm.DeviceId,
		RoomId:      rm.RoomId,
		Value:       rm.Value,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}
