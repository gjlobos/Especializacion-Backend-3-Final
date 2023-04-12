package store

import (
	"database/sql"
	"final/internal/domain"
)

type SqlStoreAppointment struct {
	DB *sql.DB
}

func NewSqlStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &SqlStoreAppointment{
		DB: db,
	}
}

// ReadAppointment devuelve un turno por su id
func (s *SqlStoreAppointment) ReadAppointment(id int) (*domain.Appointment, error) {
	var appointment domain.Appointment
	row := s.DB.QueryRow("SELECT appointments.id, appointments.appointment_date, appointments.appointment_time, appointments.description, dentists.id, dentists.name, dentists.last_name, dentists.registration_number, patients.id, patients.name, patients.last_name, patients.address, patients.personal_id_number, patients.creation_date FROM appointments JOIN dentists ON dentists.registration_number = appointments.dentist_registration_number JOIN patients ON patients.personal_id_number = appointments.patient_personal_id_number WHERE appointments.id = ?;", id)

	err := row.Scan(&appointment.Id, &appointment.Appointment_date, &appointment.Appointment_time, &appointment.Description,
		&appointment.Dentist.Id, &appointment.Dentist.Name, &appointment.Dentist.LastName, &appointment.Dentist.RegistrationNumber,
		&appointment.Patient.Id, &appointment.Patient.Name, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.Personal_id_number, &appointment.Patient.Creation_date)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// ReadAppointmentByPersonalIdNumber devuelve un turno por el DNI del paciente
func (s *SqlStoreAppointment) ReadTurnByPersonalIdNumber(personal_id_number int) (*domain.Appointment, error) {
	var appointment domain.Appointment
	row := s.DB.QueryRow("SELECT appointments.id, appointments.appointment_date, appointments.appointment_time, appointments.description, dentists.id, dentists.name, dentists.last_name, dentists.registration_number, patients.id, patients.name, patients.last_name, patients.address, patients.personal_id_number, patients.creation_date FROM appointments JOIN dentists ON dentists.registration_number = appointments.dentist_registration_number JOIN patients ON patients.personal_id_number = appointments.patient_personal_id_number WHERE appointments.patient_personal_id_number = ?;", personal_id_number)

	err := row.Scan(&appointment.Id, &appointment.Appointment_date, &appointment.Appointment_time, &appointment.Description,
		&appointment.Dentist.Id, &appointment.Dentist.Name, &appointment.Dentist.LastName, &appointment.Dentist.RegistrationNumber,
		&appointment.Patient.Id, &appointment.Patient.Name, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.Personal_id_number, &appointment.Patient.Creation_date)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

// CreateAppointment crea un turno
func (s *SqlStoreAppointment) CreateAppointment(appointment domain.Appointment) error {
	query := "INSERT INTO appointments (id, dentist_registration_number, patient_personal_id_number, appointment_date, appointment_time, description) VALUES (?, ?, ?, ?, ?, ?);"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(&appointment.Id, &appointment.Dentist.RegistrationNumber, &appointment.Patient.Personal_id_number, &appointment.Appointment_date, &appointment.Appointment_time, &appointment.Description)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// UpdateAppointment actualiza un turno por su id
func (s *SqlStoreAppointment) UpdateAppointment(appointment domain.Appointment) error {
	stmt, err := s.DB.Prepare("UPDATE appointments SET appointment_date = ?, appointment_time = ?, description = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&appointment.Appointment_date, &appointment.Appointment_time, &appointment.Description, &appointment.Id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAppointment elimina un turno por su id
func (s *SqlStoreAppointment) DeleteAppointment(id int) error {
	stmt := "DELETE FROM appointments WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
