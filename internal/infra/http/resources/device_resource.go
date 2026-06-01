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
	Units            *string               `json:"units"`
	PowerConsumption float64               `json:"powerconsumption"`
	Measurements     []MeasurementDto      `json:"measurements"`
	Events           []EventDto            `json:"events"`
}

func (d DeviceDto) DomainToDto(dv domain.Device) DeviceDto {
	dto := DeviceDto{
		Id:              dv.Id,
		OrganizationId:  dv.OrganizationId,
		RoomId:          dv.RoomId,
		GUID:            dv.GUID,
		InventoryNumber: dv.InventoryNumber,
		SerialNumber:    dv.SerialNumber,
		Characteristics: dv.Characteristics,
		Category:        dv.Category,
		Measurements:    MeasurementDto{}.MeasurementDomainToDtoCollection(dv.Measurements),
		Events:          EventDto{}.EventDomainToDtoCollection(dv.Events),
	}

	switch dv.Category {
	case domain.Sensor:
		dto.PowerConsumption = dv.PowerConsumption
		dto.Units = nil

	case domain.Actuator:
		dto.PowerConsumption = dv.PowerConsumption
		dto.Units = dv.Units

	}

	return dto
}

func (d DeviceDto) DeviceDomainToDtoCollection(dvs []domain.Device) []DeviceDto {
	dvsDto := make([]DeviceDto, len(dvs))
	for i, _ := range dvs {
		dvsDto[i] = d.DomainToDto(dvs[i])
	}
	return dvsDto
}
