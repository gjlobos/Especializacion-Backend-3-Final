package appointment

import (
	"final/internal/domain"
)

type IService interface {
	GetByID(id int) (*domain.Appointment, error)
	GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type Service struct {
	Repository IRepository
}

func NewService(r IRepository) Service {
	return &Service{r}
}

// GetByID busca un turno por su id
func (s *Service) GetByID(id int) (*domain.Appointment, error) {
	appointment, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

// GetByPersonalIdNumber busca un turno por su dni
func (s *Service) GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error) {
	appointment, err := s.Repository.GetByPersonalIdNumber(personal_id_number)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// Create crea un nuevo turno
func (s *Service) Create(appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.Repository.Create(appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// Update actualiza un turno
func (s *Service) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.Repository.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// Delete elimina un turno
func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
