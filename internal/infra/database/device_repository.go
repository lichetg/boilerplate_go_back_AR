package database

import (
	"time"

	"github.com/google/uuid"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type Device struct {
	Id               uint64                `db:"id,omitempty"`
	OrganizationId   uint64                `db:"organization_id"`
	RoomId           *uint64               `db:"room_id"`
	GUID             uuid.UUID             `db:"guid"`
	InventoryNumber  string                `db:"inventory_number"`
	SerialNumber     string                `db:"serial_number"`
	Characteristics  string                `db:"characteristics"`
	Category         domain.DeviceCategory `db:"category"`
	Units            *string               `db:"units"`
	PowerConsumption float64               `db:"power_consumption"`
	CreatedDate      time.Time             `db:"created_date"`
	UpdatedDate      time.Time             `db:"updated_date"`
	DeletedDate      *time.Time            `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(o domain.Device) (domain.Device, error)
	Find(id uint64) (domain.Device, error)
	FindGUID(GUID uuid.UUID) (domain.Device, error)
	FindList(orgId uint64) ([]domain.Device, error)
	Update(o domain.Device) (domain.Device, error)
	Delete(Id uint64) error
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(session db.Session) deviceRepository {
	return deviceRepository{
		coll: session.Collection(DevicesTableName),
		sess: session,
	}
}

func (r deviceRepository) Save(o domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(o)
	dev.GUID = uuid.New()
	now := time.Now()
	dev.CreatedDate = now
	dev.UpdatedDate = now

	err := r.coll.InsertReturning(&dev)
	if err != nil {
		return domain.Device{}, err
	}

	o = r.mapModelToDomain(dev)
	return o, nil
}

func (r deviceRepository) Find(id uint64) (domain.Device, error) {
	var device Device

	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&device)
	if err != nil {
		return domain.Device{}, err
	}

	o := r.mapModelToDomain(device)
	return o, nil
}

func (r deviceRepository) FindGUID(guid uuid.UUID) (domain.Device, error) {
	var device Device

	err := r.coll.Find(db.Cond{"guid": guid, "deleted_date": nil}).One(&device)
	if err != nil {
		return domain.Device{}, err
	}

	o := r.mapModelToDomain(device)
	return o, nil
}

func (r deviceRepository) FindList(orgId uint64) ([]domain.Device, error) {
	var devs []Device

	err := r.coll.
		Find(db.Cond{
			"organization_id": orgId,
			"deleted_date":    nil,
		}).
		All(&devs)
	if err != nil {
		return nil, err
	}

	organizations := r.mapModelToDomainCollection(devs)
	return organizations, nil
}

func (r deviceRepository) Update(o domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(o)
	dev.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&dev)
	if err != nil {
		return domain.Device{}, err
	}

	o = r.mapModelToDomain(dev)
	return o, nil
}

func (r deviceRepository) Delete(Id uint64) error {
	return r.coll.Find(db.Cond{"id": Id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) mapDomainToModel(dev domain.Device) Device {
	return Device{
		Id:               dev.Id,
		OrganizationId:   dev.OrganizationId,
		RoomId:           dev.RoomId,
		GUID:             dev.GUID,
		InventoryNumber:  dev.InventoryNumber,
		SerialNumber:     dev.SerialNumber,
		Characteristics:  dev.Characteristics,
		Category:         dev.Category,
		Units:            dev.Units,
		PowerConsumption: dev.PowerConsumption,
		CreatedDate:      dev.CreatedDate,
		UpdatedDate:      dev.UpdatedDate,
		DeletedDate:      dev.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(dev Device) domain.Device {
	return domain.Device{
		Id:               dev.Id,
		OrganizationId:   dev.OrganizationId,
		RoomId:           dev.RoomId,
		GUID:             dev.GUID,
		InventoryNumber:  dev.InventoryNumber,
		SerialNumber:     dev.SerialNumber,
		Characteristics:  dev.Characteristics,
		Category:         dev.Category,
		Units:            dev.Units,
		PowerConsumption: dev.PowerConsumption,
		CreatedDate:      dev.CreatedDate,
		UpdatedDate:      dev.UpdatedDate,
		DeletedDate:      dev.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomainCollection(devs []Device) []domain.Device {
	devices := make([]domain.Device, len(devs))
	for i := range devs {
		devices[i] = r.mapModelToDomain(devs[i])
	}
	return devices
}
