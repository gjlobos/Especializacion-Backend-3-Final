package store

import (
	"final/internal/domain"
)

type StoreInterfaceDentist interface {
	// ReadDentist devuelve un dentista por su id
	ReadDentist(id int) (domain.Dentist, error)
	// CreateDentist agrega un nuevo dentista
	CreateDentist(dentist domain.Dentist) error
	// UpdateDentist actualiza un dentista
	UpdateDentist(dentist domain.Dentist) error
	// DeleteDentist elimina un dentista
	DeleteDentist(id int) error
}

type StoreInterfacePatient interface {
	// ReadPatient devuelve un paciente por su id
	ReadPatient(id int) (domain.Patient, error)
	// CreatePatient agrega un nuevo paciente
	CreatePatient(patient domain.Patient) error
	// UpdatePatient actualiza un paciente
	UpdatePatient(patient domain.Patient) error
	// DeletePatient elimina un paciente
	DeletePatient(id int) error
}

type StoreInterfaceAppointment interface {
	// ReadAppointment devuelve un turno por su id
	ReadAppointment(id int) (domain.Appointment, error)
	// CreateAppointment agrega un nuevo turno
	CreateAppointment(appointment domain.Appointment) error
	// UpdateAppointment actualiza un turno
	UpdateAppointment(appointment domain.Appointment) error
	// DeleteAppointment elimina un turno
	DeleteAppointment(id int) error
	// ReadAppointmentByPersonalIdNumber devuelve un turno por el DNI del paciente
	ReadAppointmentByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
}
