package patient

import (
	"final/internal/domain"
)

type IService interface {
	GetByID(id int) (*domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type Service struct {
	Repository IRepository
}

func NewService(r IRepository) Service {
	return &Service{r}
}

// GetByID busca un paciente por su id
func (s *Service) GetByID(id int) (*domain.Patient, error) {
	patient, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

// Create crea un nuevo paciente
func (s *Service) Create(patient domain.Patient) (domain.Patient, error) {
	patient, err := s.Repository.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

// Update actualiza un paciente
func (s *Service) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.Repository.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

// Delete elimina un paciente
func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
