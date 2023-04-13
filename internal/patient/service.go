package patient

import (
	"final/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un paciente por su id
func (s *service) GetByID(id int) (domain.Patient, error) {
	patient, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

// Create crea un nuevo paciente
func (s *service) Create(patient domain.Patient) (domain.Patient, error) {
	patient, err := s.r.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

// Update actualiza un paciente
func (s *service) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.r.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

// Delete elimina un paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
