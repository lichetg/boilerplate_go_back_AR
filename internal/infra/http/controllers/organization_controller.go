package controllers

import (
	"net/http"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
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
		user := r.Context().Value(UserKey).(domain.User)
		Success(w, resources.UserDto{}.DomainToDto(user))
	}
}
