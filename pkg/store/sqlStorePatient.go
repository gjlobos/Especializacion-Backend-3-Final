package store

import (
	"database/sql"
	"final/internal/domain"
)

type SqlStorePatient struct {
	DB *sql.DB
}

func NewSqlStorePatient(db *sql.DB) StoreInterfacePatient {
	return &SqlStorePatient{
		DB: db,
	}
}

// ReadPatient devuelve un paciente por su id
func (s *SqlStorePatient) ReadPatient(id int) (*domain.Patient, error) {
	var patient domain.Patient
	row := s.DB.QueryRow("SELECT * FROM patients WHERE id = ?;", id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.Personal_id_number, &patient.Creation_date)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// CreatePatient crea un paciente
func (s *SqlStorePatient) CreatePatient(patient domain.Patient) error {
	query := "INSERT INTO patients (id, name, last_name, address, personal_id_number, creation_date) VALUES (?, ?, ?, ?, ?, ?);"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.Personal_id_number, &patient.Creation_date)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// UpdatePatient actualiza un paciente
func (s *SqlStorePatient) UpdatePatient(patient domain.Patient) error {
	stmt, err := s.DB.Prepare("UPDATE patients SET name = ?, last_name = ?, address = ?, personal_id_number = ?, creation_date = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&patient.Name, &patient.LastName, &patient.Address, &patient.Personal_id_number, &patient.Creation_date, &patient.Id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePatient elimina un paciente por el id
func (s *SqlStorePatient) DeletePatient(id int) error {
	stmt := "DELETE FROM patients WHERE id = ?;"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
