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
	rmService app.RoomService
}

func NewMeasurementController(mss app.MeasurementService, dvs app.DeviceService, rms app.RoomService) MeasurementController {
	return MeasurementController{
		msService: mss,
		dvService: dvs,
		rmService: rms,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		meas, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Save(requests.Bind): %s", err)
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

		if *meas.RoomId != *dev.RoomId {
			Forbidden(w, errors.New("wrong room"))
			return
		}

		if dev.Category != domain.Actuator {
			Forbidden(w, errors.New("access denied (wrong device category)"))
			return
		}

		meas.DeviceId = dev.Id

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

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dv.OrganizationId {
			Forbidden(w, errors.New("wrong organization"))
			return
		}

		if dv.Id != meas.DeviceId {
			Forbidden(w, errors.New("wrong device"))
			return
		}

		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}

func (c MeasurementController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)
		meas := r.Context().Value(MeasKey).(domain.Measurement)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		if dev.Category != domain.Actuator {
			Forbidden(w, errors.New("access denied (wrong device category)"))
			return
		}

		newMeas, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Update(requests.Update): %s", err)
			BadRequest(w, err)
			return
		}

		if newMeas.RoomId != dev.RoomId {
			Forbidden(w, errors.New("wrong room"))
			return
		}

		if newMeas.RoomId != nil {
			room, err := c.rmService.Find(*newMeas.RoomId)
			if err != nil {
				log.Printf("MeasurementController.Update(c.rmService.Find): %s", err)
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

		meas.RoomId = newMeas.RoomId
		meas.Value = newMeas.Value

		meas, err = c.msService.Update(meas)
		if err != nil {
			log.Printf("MeasurementController.Update(c.msService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}

func (c MeasurementController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)
		meas := r.Context().Value(MeasKey).(domain.Measurement)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		if org.Id != dev.OrganizationId {
			Forbidden(w, errors.New("access denied (another organization)"))
			return
		}

		if dev.Id != meas.DeviceId {
			Forbidden(w, errors.New("access denied (another device)"))
			return
		}

		err := c.msService.Delete(meas.Id)
		if err != nil {
			log.Printf("MeasurementController.Delete(c.msService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}

func (c MeasurementController) Day() http.HandlerFunc {
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

		meas, err := c.msService.GetDay(device.Id)

		if err != nil {
			log.Printf("MeasurementController.Day(c.msService.GetDay): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.MeasurementDto{}.MeasurementDomainToDtoCollection(meas))
	}
}

func (c MeasurementController) Weak() http.HandlerFunc {
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

		meas, err := c.msService.GetWeek(device.Id)

		if err != nil {
			log.Printf("MeasurementController.Weak(c.msService.GetWeak): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.MeasurementDto{}.MeasurementDomainToDtoCollection(meas))
	}
}

func (c MeasurementController) Month() http.HandlerFunc {
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

		meas, err := c.msService.GetMonth(device.Id)

		if err != nil {
			log.Printf("MeasurementController.Month(c.msService.GetMonth): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.MeasurementDto{}.MeasurementDomainToDtoCollection(meas))
	}
}
