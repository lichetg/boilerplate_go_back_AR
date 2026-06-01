package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
	"log"
	"time"
)

const MeasurementsTableName = "measurements"

type Measurement struct {
	Id          uint64     `db:"id,omitempty"`
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
	Update(o domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
	FindByDeviceAndPeriod(id uint64, from time.Time, to time.Time) ([]domain.Measurement, error)
	FindByDeviceId(devId uint64) ([]domain.Measurement, error)
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

	log.Println(meas)
	err := r.coll.InsertReturning(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	o = r.mapModelToDomain(meas)
	return o, nil
}

func (r measurementRepository) Find(id uint64) (domain.Measurement, error) {
	var meas Measurement

	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	o := r.mapModelToDomain(meas)
	return o, nil
}

func (r measurementRepository) FindByDeviceId(devId uint64) ([]domain.Measurement, error) {
	var meass []Measurement

	err := r.coll.Find(db.Cond{
		"device_id":    devId,
		"deleted_date": nil,
	}).All(&meass)

	if err != nil {
		return nil, err
	}

	o := r.mapModelToDomainCollection(meass)
	return o, nil
}

func (r measurementRepository) FindByDeviceAndPeriod(
	id uint64,
	from time.Time,
	to time.Time,
) ([]domain.Measurement, error) {
	var meas []Measurement

	err := r.coll.Find(
		db.Cond{
			"device_id":       id,
			"created_date >=": from,
			"created_date <=": to,
			"deleted_date":    nil,
		},
	).All(&meas)

	if err != nil {
		return nil, err
	}

	measurements := make([]domain.Measurement, len(meas))

	for i, m := range meas {
		measurements[i] = r.mapModelToDomain(m)
	}

	return measurements, nil
}

func (r measurementRepository) Update(o domain.Measurement) (domain.Measurement, error) {
	meas := r.mapDomainToModel(o)
	meas.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	o = r.mapModelToDomain(meas)
	return o, nil
}

func (r measurementRepository) Delete(Id uint64) error {
	return r.coll.Find(db.Cond{"id": Id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r measurementRepository) mapDomainToModel(rm domain.Measurement) Measurement {
	return Measurement{
		Id:          rm.Id,
		DeviceId:    rm.DeviceId,
		RoomId:      rm.RoomId,
		Value:       rm.Value,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}

func (r measurementRepository) mapModelToDomain(rm Measurement) domain.Measurement {
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

func (r measurementRepository) mapModelToDomainCollection(meass []Measurement) []domain.Measurement {
	mss := make([]domain.Measurement, len(meass))
	for i := range meass {
		mss[i] = r.mapModelToDomain(meass[i])
	}
	return mss
}
