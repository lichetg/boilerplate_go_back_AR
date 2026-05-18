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
//	FindByOrgId(uId uint64) ([]domain.Room, error)
	Find(id uint64) (interface{}, error)
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

//func (s roomService) FindByOrgId(uId uint64) ([]domain.Room, error) {
//	rms, err := s.roomRepo.FindByOrgId(uId)
//	if err != nil {
//		log.Printf("roomService.FindByOrgId(s.roomRepo.FindByOrgId): %s", err)
//		return nil, err
//	}
//
//	return rms, nil
//}

func (s roomService) Find(id uint64) (interface{}, error) {
	rm, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.orgRepo.Find): %s", err)
		return nil, err
	}

	return rm, nil
}
