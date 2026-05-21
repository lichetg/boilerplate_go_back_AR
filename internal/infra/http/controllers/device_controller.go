package controllers

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"net/http"
)

type DeviceController struct {
	dvService app.DeviceService
}

func NewDeviceController(ds app.DeviceService) DeviceController {
	return DeviceController{
		dvService: ds,
	}
}

func (d DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
	}
}
