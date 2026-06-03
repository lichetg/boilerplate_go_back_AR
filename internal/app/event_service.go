package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
)

type eventService struct {
	eveRepo database.EventRepository
}

type EventService interface {
	Save(o domain.Event) (domain.Event, error)
	Find(id uint64) (interface{}, error)
	FindList(p domain.Pagination, f domain.EventFilters) (resources.Events, error)
	Update(o domain.Event) (domain.Event, error)
	Delete(id uint64) error
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

func (s eventService) FindList(p domain.Pagination, f domain.EventFilters) (resources.Events, error) {
	eves, err := s.eveRepo.FindList(p, f)
	if err != nil {
		log.Printf("eventService.FindList(s.eveRepo.FindList): %s", err)
		return resources.Events{}, err
	}

	return eves, nil
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
