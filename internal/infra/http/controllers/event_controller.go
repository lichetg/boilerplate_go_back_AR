package controllers

import (
	"errors"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
	"net/http"
)

type EventController struct {
	evService app.EventService
	rmService app.RoomService
}

func NewEventController(evs app.EventService, rms app.RoomService) EventController {
	return EventController{
		evService: evs,
		rmService: rms,
	}
}

func (c EventController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		eve, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return

		}

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
		}

		if dev.OrganizationId != org.Id {
			Forbidden(w, errors.New("access denied (wrong organization)"))
			return
		}

		if *eve.RoomId != *dev.RoomId {
			Forbidden(w, errors.New("access denied (wrong room)"))
			return
		}

		if dev.Category != domain.Sensor {
			Forbidden(w, errors.New("access denied (wrong device category)"))
			return
		}
		eve.DeviceId = dev.Id

		eve, err = c.evService.Save(eve)
		if err != nil {
			log.Printf("EventController.Save(c.evService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		eveDto := resources.EventDto{}
		eveDto = eveDto.DomainToDto(eve)
		Success(w, eveDto)
	}
}

func (c EventController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dv := r.Context().Value(DeviceKey).(domain.Device)
		eve := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dv.OrganizationId {
			Forbidden(w, errors.New("wrong organization"))
			return
		}

		if dv.Id != eve.DeviceId {
			Forbidden(w, errors.New("wrong device"))
			return
		}

		Success(w, resources.EventDto{}.DomainToDto(eve))
	}
}

func (c EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)
		eve := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		if dev.Category != domain.Sensor {
			Forbidden(w, errors.New("access denied (wrong device category)"))
			return
		}

		newEve, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController.Update(requests.Update): %s", err)
			BadRequest(w, err)
			return
		}

		if *newEve.RoomId != *dev.RoomId {
			Forbidden(w, errors.New("wrong room"))
			return
		}

		if newEve.RoomId != nil {
			room, err := c.rmService.Find(*newEve.RoomId)
			if err != nil {
				log.Printf("EventController.Update(c.evService.Find): %s", err)
				BadRequest(w, errors.New("room not found"))
				return
			}

			rm, ok := room.(domain.Room)
			if !ok {
				InternalServerError(w, errors.New("invalid room type"))
				return
			}

			if rm.OrganizationId != org.Id {
				Forbidden(w, errors.New("room does not belong to organization"))
				return
			}
		}

		eve.RoomId = newEve.RoomId
		eve.Action = newEve.Action

		eve, err = c.evService.Update(eve)
		if err != nil {
			log.Printf("EventController.Update(c.evService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.DomainToDto(eve))
	}
}

func (c EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)
		eve := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		if dev.Id != eve.DeviceId {
			Forbidden(w, errors.New("access denied (another device)"))
			return
		}

		err := c.evService.Delete(eve.Id)
		if err != nil {
			log.Printf("EventController.Delete(c.evService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}

func (c EventController) Day() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		device := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != device.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		eve, err := c.evService.GetDay(device.Id)

		if err != nil {
			log.Printf("EventController.Day(c.evService.GetDay): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.EventDomainToDtoCollection(eve))
	}
}

func (c EventController) Weak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		device := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != device.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		eve, err := c.evService.GetWeak(device.Id)

		if err != nil {
			log.Printf("EventController.Weak(c.evService.GetWeak): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.EventDomainToDtoCollection(eve))
	}
}

func (c EventController) Month() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		device := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != device.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		eve, err := c.evService.GetMonth(device.Id)

		if err != nil {
			log.Printf("EventController.Month(c.evService.GetMount): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.EventDomainToDtoCollection(eve))
	}
}
