package domain

type Patient struct {
	Id                 int    `json:"id"`
	Name               string `json:"name" binding:"required"`
	LastName           string `json:"last_name" binding:"required"`
	Address            string `json:"address" binding:"required"`
	Personal_id_number int    `json:"personal_id_number" binding:"required"`
	Creation_date      string `json:"creation_date" binding:"required"`
}

type PatientDTO struct {
	Name               string `json:"name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	Address            string `json:"address,omitempty"`
	Personal_id_number int    `json:"personal_id_number,omitempty"`
	Creation_date      string `json:"creation_date,omitempty"`
}
