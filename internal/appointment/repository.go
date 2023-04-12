package appointment

import (
	"errors"
	"final/internal/domain"
	"final/pkg/store"
	"final/pkg/web"
	"fmt"
)

type IRepository interface {
	GetByID(id int) (*domain.Appointment, error)
	GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type Repository struct {
	Store store.StoreInterfaceAppointment
}

func NewRepository(storage store.StoreInterfaceAppointment) Repository {
	return &Repository{
		Store: storage,
	}
}

// GetByID busca un turno por su id
func (r *Repository) GetByID(id int) (*domain.Appointment, error) {
	turn, err := r.Store.ReadAppointment(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("appointment_id %d not found"))
	}
	return turn, nil
}

// GetByPersonalIdNumber busca un turno por su dni
func (r *Repository) GetByPersonalIdNumber(personal_id_number int) (*domain.Appointment, error) {
	appointment, err := r.Store.ReadAppointmentByPersonalIdNumber(personal_id_number)
	if err != nil {
		return nil, errors.New("appointment not found")
	}
	return appointment, nil
}

// Create crea un nuevo turno
func (r *Repository) Create(appointment domain.Appointment) (domain.Appointment, error) {
	err := r.Store.CreateAppointment(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating turn")
	}
	return appointment, nil
}

// Update actualiza un turno
func (r *Repository) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	appointment.Id = id
	err := r.Store.UpdateAppointment(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("error updating turn")
	}
	return appointment, nil
}

// Delete elimina un paciente
func (r *Repository) Delete(id int) error {
	err := r.Store.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}
