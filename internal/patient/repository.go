package patient

import (
	"errors"
	"final/internal/domain"
	"final/pkg/store"
	"final/pkg/web"
	"fmt"
)

type IRepository interface {
	GetByID(id int) (*domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type Repository struct {
	Store store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &Repository{
		Store: storage,
	}
}

// GetByID busca un paciente por su id
func (r *Repository) GetByID(id int) (*domain.Patient, error) {
	patient, err := r.Store.ReadPatient(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("patient_id %d not found"))
	}
	return patient, nil
}

// Create crea un nuevo paciente
func (r *Repository) Create(patient domain.Patient) (domain.Patient, error) {
	err := r.Store.CreatePatient(patient)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return patient, nil
}

// Update actualiza un paciente
func (r *Repository) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patient.Id = id
	err := r.Store.UpdatePatient(patient)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return patient, nil
}

// Delete elimina un paciente
func (r *Repository) Delete(id int) error {
	err := r.Store.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
