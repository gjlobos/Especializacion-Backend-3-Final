package patient

import (
	"errors"
	"final/internal/domain"
	"final/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	Storage store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &repository{
		Storage: storage,
	}
}

// GetByID busca un paciente por su id
func (r *repository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.Storage.ReadPatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}

// Create crea un nuevo paciente
func (r *repository) Create(patient domain.Patient) (domain.Patient, error) {
	err := r.Storage.CreatePatient(patient)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return patient, nil
}

// Update actualiza un paciente
func (r *repository) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patient.Id = id
	err := r.Storage.UpdatePatient(patient)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return patient, nil
}

// Delete elimina un paciente
func (r *repository) Delete(id int) error {
	err := r.Storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
