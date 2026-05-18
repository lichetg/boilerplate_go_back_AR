package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type RoomController struct {
	rmService app.RoomService
}

func NewRoomController(rs app.RoomService) RoomController {
	return RoomController{
		rmService: rs,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rm, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		org := r.Context().Value(OrgKey).(domain.Organization)
		rm.OrganizationId = org.Id

		rm, err = c.rmService.Save(rm)
		if err != nil {
			log.Printf("RoomController.Save(c.rmService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		rmDto := resources.RoomDto{}
		rmDto = rmDto.DomainToDto(rm)
		Success(w, rmDto)
	}
}


//func (c RoomController) FindList() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		user := r.Context().Value(UserKey).(domain.User)
//
//		orgs, err := c.rmService.FindList(user.Id)
//		if err != nil {
//			log.Printf("OrganizationController.FindList(c.orgService.FindList): %s", err)
//			InternalServerError(w, err)
//			return
//		}
//
//		Success(w, resources.OrganizationDto{}.DomainToDtoCollection(orgs))
//	}
//}

func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)
		rm := r.Context().Value(RoomKey).(domain.Room)

		if org.Id != rm.OrganizationId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		Success(w, resources.OrganizationDto{}.DomainToDto(org))
	}
}