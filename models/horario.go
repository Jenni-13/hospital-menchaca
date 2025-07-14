package models

type Horario struct {
	IdHorario     int    `json:"id_horario"`
	Turno         string `json:"turno"`
	IdConsultorio int    `json:"id_consultorio"`
	IdMedico      int    `json:"id_medico"`
}
