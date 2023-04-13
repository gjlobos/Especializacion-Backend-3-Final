package appointment

import (
	"errors"
	"final/internal/domain"
	"final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Appointment, error)
	GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	Storage store.StoreInterfaceAppointment
}

func NewRepository(storage store.StoreInterfaceAppointment) Repository {
	return &repository{
		Storage: storage,
	}
}

// GetByID busca un turno por su id
func (r *repository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.Storage.ReadAppointment(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

// GetByPersonalIdNumber busca un turno por su dni
func (r *repository) GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error) {
	appointment, err := r.Storage.ReadAppointmentByPersonalIdNumber(personal_id_number)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

// Create crea un nuevo turno
func (r *repository) Create(appointment domain.Appointment) (domain.Appointment, error) {
	err := r.Storage.CreateAppointment(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating appointment")
	}
	return appointment, nil
}

// Update actualiza un turno
func (r *repository) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	appointment.Id = id
	err := r.Storage.UpdateAppointment(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("error updating appointment")
	}
	return appointment, nil
}

// Delete elimina un paciente
func (r *repository) Delete(id int) error {
	err := r.Storage.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}
