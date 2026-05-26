package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type measurementService struct {
	measRepo database.MeasurementRepository
}

type MeasurementService interface {
	Save(o domain.Measurement) (domain.Measurement, error)
	Find(id uint64) (interface{}, error)
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
	rm, err := s.measRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.measRepo.Find): %s", err)
		return nil, err
	}

	return rm, nil
}
