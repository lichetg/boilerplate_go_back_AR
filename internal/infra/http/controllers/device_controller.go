package controllers

import (
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

		rm := r.Context().Value(RoomKey).(domain.Room)
		dev.RoomId = &rm.Id

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
