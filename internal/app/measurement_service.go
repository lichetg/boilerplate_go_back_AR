package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
	"time"
)

type measurementService struct {
	measRepo database.MeasurementRepository
}

type MeasurementService interface {
	Save(o domain.Measurement) (domain.Measurement, error)
	Find(id uint64) (interface{}, error)
	Update(o domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error

	GetDay(deviceId uint64) ([]domain.Measurement, error)
	GetWeek(deviceId uint64) ([]domain.Measurement, error)
	GetMonth(deviceId uint64) ([]domain.Measurement, error)
}

func NewMeasurementService(
	mr database.MeasurementRepository) MeasurementService {
	return measurementService{
		measRepo: mr,
	}
}

func (s measurementService) Save(o domain.Measurement) (domain.Measurement, error) {
	meas, err := s.measRepo.Save(o)
	if err != nil {
		log.Printf("measuremenetService.Save(s.measRepo.Save): %s", err)
		return domain.Measurement{}, err
	}

	return meas, nil
}

func (s measurementService) Find(id uint64) (interface{}, error) {
	meas, err := s.measRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.measRepo.Find): %s", err)
		return nil, err
	}

	return meas, nil
}

func (s measurementService) Update(o domain.Measurement) (domain.Measurement, error) {
	meas, err := s.measRepo.Update(o)
	if err != nil {
		log.Printf("measurementService.Update(s.measRepo.Update): %s", err)
		return domain.Measurement{}, err
	}

	return meas, nil
}

func (s measurementService) Delete(id uint64) error {
	err := s.measRepo.Delete(id)
	if err != nil {
		log.Printf("measurementService.Delete(s.measRepo.Delete): %s", err)
		return err
	}

	return nil
}

func (s measurementService) GetDay(deviceId uint64) ([]domain.Measurement, error) {

	now := time.Now()

	from := now.AddDate(0, 0, -1)

	return s.measRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}

func (s measurementService) GetWeek(deviceId uint64) ([]domain.Measurement, error) {

	now := time.Now()

	from := now.AddDate(0, 0, -7)

	return s.measRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}

func (s measurementService) GetMonth(deviceId uint64) ([]domain.Measurement, error) {

	now := time.Now()

	from := now.AddDate(0, -1, 0)

	return s.measRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}
