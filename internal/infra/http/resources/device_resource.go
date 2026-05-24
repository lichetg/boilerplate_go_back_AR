package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type DeviceDto struct {
	Id               uint64                `json:"id"`
	OrganizationId   uint64                `json:"organization_id"`
	RoomId           *uint64               `json:"room_id"`
	GUID             uuid.UUID             `json:"guid"`
	InventoryNumber  string                `json:"inventorynumber"`
	SerialNumber     string                `json:"serialnumber"`
	Characteristics  string                `json:"characteristics"`
	Category         domain.DeviceCategory `json:"category"`
	Units            string                `json:"units"`
	PowerConsumption float64               `json:"powerconsumption"`
}

func (d DeviceDto) DomainToDto(dv domain.Device) DeviceDto {
	return DeviceDto{
		Id:               dv.Id,
		OrganizationId:   dv.OrganizationId,
		RoomId:           dv.RoomId,
		GUID:             dv.GUID,
		InventoryNumber:  dv.InventoryNumber,
		SerialNumber:     dv.SerialNumber,
		Characteristics:  dv.Characteristics,
		Category:         dv.Category,
		Units:            dv.Units,
		PowerConsumption: dv.PowerConsumption,
	}
}
