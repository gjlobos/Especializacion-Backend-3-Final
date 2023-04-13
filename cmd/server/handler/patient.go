package handler

import (
	"errors"
	"strconv"

	"final/internal/domain"
	"final/internal/patient"
	"final/pkg/web"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

// NewPatientHandler crea un nuevo controller de pacientes
func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// GetByID obtiene un paciente por su id
func (h *patientHandler) GetPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// PostPatient crea un nuevo paciente
func (h *patientHandler) PostPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyFieldsPatient(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// PutPatient actualiza un paciente
func (h *patientHandler) PutPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyFieldsPatient(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// PatchPatchPatient actualiza un paciente por id
func (h *patientHandler) PatchPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var PatientNew domain.PatientDTO
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if err := c.ShouldBindJSON(&PatientNew); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		update := domain.Patient{
			Name:               PatientNew.Name,
			LastName:           PatientNew.LastName,
			Address:            PatientNew.Address,
			Personal_id_number: PatientNew.Personal_id_number,
			Creation_date:      PatientNew.Creation_date,
		}

		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// DeletePatient elimina un turno
func (h *patientHandler) DeletePatient() gin.HandlerFunc {
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
