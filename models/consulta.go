package models

import (
	"database/sql"
)


type Consulta struct {
	IdConsulta    int     `json:"id_consulta"`
	Tipo          string  `json:"tipo"`
	Diagnostico   string  `json:"diagnostico"`
	Costo         float64 `json:"costo"`
	Fecha         sql.NullTime    `json:"fecha"` 
	IdPaciente    int     `json:"id_paciente"`
	IdMedico      int     `json:"id_medico"`
	IdConsultorio int     `json:"id_consultorio"`
	IdHorario     int     `json:"id_horario"`
}
