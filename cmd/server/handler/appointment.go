package handler

import (
	"errors"
	"strconv"

	"final/internal/appointment"
	"final/internal/domain"
	"final/pkg/web"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

// NewAppointmentHandler crea un nuevo controller de turnos
func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

// GetByID obtiene un turno por su id
func (h *appointmentHandler) GetAppointmentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}

// GetAppointmentByPersonalIdNumber obtiene un turno por el dni del paciente
func (h *appointmentHandler) GetAppointmentByPersonalIdNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		personalidnumberParam := c.Param("personal_id_number")
		personal_id_number, err := strconv.Atoi(personalidnumberParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		appointment, err := h.s.GetByPersonalIdNumber(personal_id_number)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}

// PostAppointment crea un nuevo turno
func (h *appointmentHandler) PostAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyFieldsAppointment(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// PutAppointment actualiza un turno
func (h *appointmentHandler) PutAppointment() gin.HandlerFunc {
	type Request struct {
		Dentist          domain.Dentist `json:"dentist"`
		Patient          domain.Patient `json:"patient"`
		Appointment_date string         `json:"appointment_date"`
		Appointment_time string         `json:"appointment_time"`
		Description      string         `json:"description"`
	}
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var r Request
		err = c.ShouldBindJSON(&r)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		update := domain.Appointment{
			Dentist:          r.Dentist,
			Patient:          r.Patient,
			Appointment_date: r.Appointment_date,
			Appointment_time: r.Appointment_time,
			Description:      r.Description,
		}

		t, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t)
	}
}

// PatchAppointment actualiza un turno
func (h *appointmentHandler) PatchAppointment() gin.HandlerFunc {
	type Request struct {
		Dentist          domain.Dentist `json:"dentist,omitempty"`
		Patient          domain.Patient `json:"patient,omitempty"`
		Appointment_date string         `json:"appointment_date,omitempty"`
		Appointment_time string         `json:"appointment_time,omitempty"`
		Description      string         `json:"description,omitempty"`
	}
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var r Request
		err = c.ShouldBindJSON(&r)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		update := domain.Appointment{
			Dentist:          r.Dentist,
			Patient:          r.Patient,
			Appointment_date: r.Appointment_date,
			Appointment_time: r.Appointment_time,
			Description:      r.Description,
		}

		t, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t)
	}
}

// DeleteAppointment elimina un turno
func (h *appointmentHandler) DeleteAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}
