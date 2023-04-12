package domain

type Appointment struct {
	Id               int     `json:"id"`
	Dentist          Dentist `json:"dentist" binding:"required"`
	Patient          Patient `json:"patient" binding:"required"`
	Appointment_date string  `json:"appointment_date" binding:"required"`
	Appointment_time string  `json:"appointment_time" binding:"required"`
	Description      string  `json:"description" binding:"required"`
}

type AppointmentDTO struct {
	Appointment_date string `json:"appointment_date,omitempty"`
	Appointment_time string `json:"appointment_time,omitempty"`
	Description      string `json:"description,omitempty"`
}
