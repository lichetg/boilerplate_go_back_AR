package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
)

type measurementService struct {
	measRepo database.MeasurementRepository
}

type MeasurementService interface {
	Save(o domain.Measurement) (domain.Measurement, error)
	Find(id uint64) (interface{}, error)
	FindList(p domain.Pagination, f domain.MeasurementFilters) (resources.Measurements, error)
	Update(o domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
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

func (s measurementService) FindList(p domain.Pagination, f domain.MeasurementFilters) (resources.Measurements, error) {
	meass, err := s.measRepo.FindList(p, f)
	if err != nil {
		log.Printf("measurementService.FindList(s.measRepo.FindList): %s", err)
		return resources.Measurements{}, err
	}

	return meass, nil
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
