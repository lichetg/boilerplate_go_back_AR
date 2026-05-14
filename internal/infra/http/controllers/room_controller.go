package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type RoomControler struct {
	rmService app.RoomService
}

func NewRoomController(rs app.RoomService) RoomController {
	return RoomController{
		rmService: rs,
	}
}

func (c RoomControler) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org, err := requests.Bind(r, requests.OrganizationRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org.UserId = user.Id

		org, err = c.rmService.Save(org)
		if err != nil {
			log.Printf("OrganizationController.Save(c.orgService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		orgDto := resources.OrganizationDto{}
		orgDto = orgDto.DomainToDto(org)
		Success(w, orgDto)
	}
}
