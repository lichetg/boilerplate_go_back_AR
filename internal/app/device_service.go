package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type deviceService struct {
	devRepo database.DeviceRepository
}

type DeviceService interface {
	Save(o domain.Device) (domain.Device, error)
}

func NewDeviceService(
	dr database.DeviceRepository) DeviceService {
	return deviceService{
		devRepo: dr,
	}
}

func (s deviceService) Save(o domain.Device) (domain.Device, error) {
	dv, err := s.devRepo.Save(o)
	if err != nil {
		log.Printf("deviceService.Save(s.devRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return dv, nil
}
