package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
	"time"
)

type eventService struct {
	eveRepo database.EventRepository
}

type EventService interface {
	Save(o domain.Event) (domain.Event, error)
	Find(id uint64) (interface{}, error)
	Update(o domain.Event) (domain.Event, error)
	Delete(id uint64) error

	GetDay(deviceId uint64) ([]domain.Event, error)
	GetWeak(deviceId uint64) ([]domain.Event, error)
	GetMonth(deviceId uint64) ([]domain.Event, error)
}

func NewEventService(
	er database.EventRepository) EventService {
	return eventService{
		eveRepo: er,
	}
}

func (s eventService) Save(o domain.Event) (domain.Event, error) {
	event, err := s.eveRepo.Save(o)
	if err != nil {
		log.Printf("eventService.Save(s.eveRepo.Save): %s", err)
		return domain.Event{}, err
	}

	return event, nil
}

func (s eventService) Find(id uint64) (interface{}, error) {
	eve, err := s.eveRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.measRepo.Find): %s", err)
		return nil, err
	}

	return eve, nil
}

func (s eventService) Update(o domain.Event) (domain.Event, error) {
	eve, err := s.eveRepo.Update(o)
	if err != nil {
		log.Printf("eventService.Update(s.eveRepo.Update): %s", err)
		return domain.Event{}, err
	}

	return eve, nil
}

func (s eventService) Delete(id uint64) error {
	err := s.eveRepo.Delete(id)
	if err != nil {
		log.Printf("eventService.Delete(s.eveRepo.Delete): %s", err)
		return err
	}

	return nil
}

func (s eventService) GetDay(deviceId uint64) ([]domain.Event, error) {

	now := time.Now()

	from := now.AddDate(0, 0, -1)

	return s.eveRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}

func (s eventService) GetWeak(deviceId uint64) ([]domain.Event, error) {

	now := time.Now()

	from := now.AddDate(0, 0, -7)

	return s.eveRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}

func (s eventService) GetMonth(deviceId uint64) ([]domain.Event, error) {

	now := time.Now()

	from := now.AddDate(0, -1, 0)

	return s.eveRepo.FindByDeviceAndPeriod(
		deviceId,
		from,
		now,
	)
}
