package store

import (
	"database/sql"
	"final/internal/domain"
)

type sqlStoreDentist struct {
	DB *sql.DB
}

func NewSqlStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &sqlStoreDentist{
		DB: db,
	}
}

// Read devuelve un dentista por su id
func (s *sqlStoreDentist) ReadDentist(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	row := s.DB.QueryRow("SELECT * FROM dentists WHERE id = ?;", id)
	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.RegistrationNumber)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStoreDentist) CreateDentist(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (id, name, last_name, registration_number) VALUES (?, ?, ?, ?);"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.RegistrationNumber)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreDentist) UpdateDentist(dentist domain.Dentist) error {
	stmt, err := s.DB.Prepare("UPDATE dentists SET name = ?, last_name = ?, registration_number = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&dentist.Name, &dentist.LastName, &dentist.RegistrationNumber, &dentist.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStoreDentist) DeleteDentist(id int) error {
	stmt := "DELETE FROM dentists WHERE id = ?;"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
