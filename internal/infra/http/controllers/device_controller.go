package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type DeviceController struct {
	dvService app.DeviceService
}

func NewDeviceController(ds app.DeviceService) DeviceController {
	return DeviceController{
		dvService: ds,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return

		}

		org := r.Context().Value(OrgKey).(domain.Organization)
		dev.OrganizationId = org.Id

		dev, err = c.dvService.Save(dev)
		if err != nil {
			log.Printf("DeviceController.Save(c.dvService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		dvDto := resources.DeviceDto{}
		dvDto = dvDto.DomainToDto(dev)
		Success(w, dvDto)
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		newDev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Update(requests.Update): %s", err)
			BadRequest(w, err)
			return
		}

		/*
			Глибоко
			if *newRoom.RoomId != rm.OrganizationId {
				if user.Id != org.UserId {
					Forbidden(w, errors.New("access denied"))
					return
				}
				Forbidden(w, errors.New("access denied"))
				return
			}
		*/

		dev.RoomId = newDev.RoomId
		dev.InventoryNumber = newDev.InventoryNumber
		dev.SerialNumber = newDev.SerialNumber
		dev.Characteristics = newDev.Characteristics
		dev.Units = newDev.Units
		dev.PowerConsumption = newDev.PowerConsumption

		dev, err = c.dvService.Update(dev)
		if err != nil {
			log.Printf("DeviceController.Update(c.dvService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		err := c.dvService.Delete(dev.Id)
		if err != nil {
			log.Printf("DeviceController.Delete(c.dvService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
