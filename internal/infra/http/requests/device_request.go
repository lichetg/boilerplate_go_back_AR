package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceCategory string

const (
	Sensor   DeviceCategory = "SENSOR"
	Actuator DeviceCategory = "ACTUATOR"
)

type DeviceRequest struct {
	RoomId           *uint64               `json:"room_id"`
	InventoryNumber  string                `json:"inventory_number"`
	SerialNumber     string                `json:"serial_number"`
	Characteristics  string                `json:"characteristics"`
	Category         domain.DeviceCategory `json:"category"`
	Units            *string               `json:"units"`
	PowerConsumption float64               `json:"power_consumption"`
}

func (d DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		RoomId:           d.RoomId,
		InventoryNumber:  d.InventoryNumber,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
	}, nil
}
