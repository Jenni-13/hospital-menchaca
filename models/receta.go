package models

import "time"

type Receta struct {
	IdReceta      int    `json:"id_receta"`
	Fecha          time.Time  `json:"fecha"`
	Medicamento   string `json:"medicamento"`
	Dosis         string `json:"dosis"`
	IdMedico      int    `json:"id_medico"`
	IdPaciente    int    `json:"id_paciente"`
	IdConsultorio int    `json:"id_consultorio"`
}
