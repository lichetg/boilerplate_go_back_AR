package database

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/upper/db/v4"
	"log"
	"math"
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
	FindByDeviceId(devId uint64) ([]domain.Measurement, error)
	Update(o domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
	FindList(p domain.Pagination, f domain.MeasurementFilters) (resources.Measurements, error)
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

func (r measurementRepository) FindList(
	p domain.Pagination,
	f domain.MeasurementFilters,
) (resources.Measurements, error) {

	var models []Measurement

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

	case "value":
		query = query.OrderBy("value")

	case "-value":
		query = query.OrderBy("-value")

	default:
		query = query.OrderBy("-created_date")
	}

	result := query.Paginate(uint(p.CountPerPage))

	err := result.Page(uint(p.Page)).All(&models)
	if err != nil {
		return resources.Measurements{}, err
	}

	total, err := result.TotalEntries()
	if err != nil {
		return resources.Measurements{}, err
	}

	items := make([]domain.Measurement, len(models))

	for i, m := range models {
		items[i] = r.mapModelToDomain(m)
	}

	return resources.Measurements{
		Items: items,
		Total: total,
		Pages: uint(math.Ceil(
			float64(total) / float64(p.CountPerPage),
		)),
	}, nil
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
