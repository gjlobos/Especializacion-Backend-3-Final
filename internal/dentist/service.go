package dentist

import (
	"final/internal/domain"
)

type IService interface {
	GetByID(id int) (*domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type Service struct {
	Repository IRepository
}

func NewService(r IRepository) Service {
	return &Service{r}
}

// GetByID busca un dentista por su id
func (s *Service) GetByID(id int) (*domain.Dentist, error) {
	dentist, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

// Create crea un nuevo dentista
func (s *Service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.Repository.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Update actualiza un dentista
func (s *Service) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.Repository.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Delete elimina un dentista
func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
