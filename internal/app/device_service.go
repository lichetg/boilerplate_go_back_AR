package app

import (
	"errors"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type deviceService struct {
	devRepo database.DeviceRepository
}

type DeviceService interface {
	Save(o domain.Device) (domain.Device, error)
	Find(id uint64) (interface{}, error)
	Update(o domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

func NewDeviceService(
	dr database.DeviceRepository) DeviceService {
	return deviceService{
		devRepo: dr,
	}
}

func (s deviceService) Save(o domain.Device) (domain.Device, error) {

	if o.Category == domain.Actuator {
		if o.Units == nil || *o.Units == "" {
			return domain.Device{}, errors.New("units is required for actuator")
		}
	} else {
		o.Units = nil
	}

	dv, err := s.devRepo.Save(o)

	if err != nil {
		log.Printf("deviceService.Save(s.devRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return dv, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	rm, err := s.devRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.Find(s.devRepo.Find): %s", err)
		return nil, err
	}

	return rm, nil
}

func (s deviceService) Update(o domain.Device) (domain.Device, error) {
	dev, err := s.devRepo.Update(o)
	if err != nil {
		log.Printf("deviceService.Update(s.devRepo.Update): %s", err)
		return domain.Device{}, err
	}

	return dev, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.devRepo.Delete(id)
	if err != nil {
		log.Printf("deviceService.Delete(s.devRepo.Delete): %s", err)
		return err
	}

	return nil
}
