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

type MeasurementController struct {
	msService app.MeasurementService
	dvService app.DeviceService
}

func NewMeasurementController(mss app.MeasurementService, dvs app.DeviceService) MeasurementController {
	return MeasurementController{
		msService: mss,
		dvService: dvs,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meas, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return

		}

		meas, err = c.msService.Save(meas)
		if err != nil {
			log.Printf("MeasurementController.Save(c.msService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		mssDto := resources.MeasurementDto{}
		mssDto = mssDto.DomainToDto(meas)
		Success(w, mssDto)
	}
}

func (c MeasurementController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dv := r.Context().Value(DeviceKey).(domain.Device)
		meas := r.Context().Value(MeasKey).(domain.Measurement)

		dev, err := c.dvService.Find(dv.Id)
		if err != nil {
			Forbidden(w, errors.New("access denied"))
			return
		}

		device := dev.(domain.Device)

		if device.Id != meas.DeviceId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if dv.Category != domain.Actuator {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if dv.Id != meas.DeviceId {
			Forbidden(w, errors.New("wrong device"))
			return
		}

		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}
