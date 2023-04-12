package dentist

import (
	"errors"
	"final/internal/domain"
	"final/pkg/store"
	"final/pkg/web"
	"fmt"
)

type IRepository interface {
	GetByID(id int) (*domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type Repository struct {
	Store store.StoreInterfaceDentist
}

func NewRepository(storage store.StoreInterfaceDentist) Repository {
	return &Repository{
		Store: storage,
	}
}

// GetByID busca un dentista por su id
func (r *Repository) GetByID(id int) (*domain.Dentist, error) {
	dentist, err := r.Store.ReadDentist(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist_id %d not found"))
	}
	return dentist, nil
}

// Create crea un nuevo dentista
func (r *Repository) Create(dentist domain.Dentist) (domain.Dentist, error) {
	err := r.Store.CreateDentist(dentist)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating dentist")
	}
	return dentist, nil
}

// Update actualiza un dentista
func (r *Repository) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentist.Id = id
	err := r.Store.UpdateDentist(dentist)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return dentist, nil
}

// Delete elimina un dentista
func (r *Repository) Delete(id int) error {
	err := r.Store.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
