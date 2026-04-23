package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type OrganizationController struct {
	orgService app.OrganizationService
}

func NewOrganizationController(os app.OrganizationService) OrganizationController {
	return OrganizationController{
		orgService: os,
	}
}

func (c OrganizationController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org, err := requests.Bind(r, requests.OrganizationRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return 
		}

		user := r.Context().Value(UserKey).(domain.User)
		org.UserId = user.Id

		org, err = c.orgService.Save(org)
		if err != nil{
			log.Printf("OrganizationController.Save(c.orgService.Save): %s", err)
			InternalServerError(w, err)
			return 
		}
		Success(w, resources.UserDto{}.DomainToDto(user))
	}
}
