package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type roomService struct {
	roomRepo database.RoomRepository
}

type RoomService interface {
	Save(o domain.Room) (domain.Room, error)
	Find(id uint64) (interface{}, error)
	Update(o domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

func NewRoomService(
	rr database.RoomRepository) RoomService {
	return roomService{
		roomRepo: rr,
	}
}

func (s roomService) Save(o domain.Room) (domain.Room, error) {
	rm, err := s.roomRepo.Save(o)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	rm, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return nil, err
	}

	return rm, nil
}

func (s roomService) Update(o domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Update(o)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}

	return nil
}
