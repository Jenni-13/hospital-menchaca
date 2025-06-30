package models

type Expediente struct {
	IdExpediente     int    `json:"id_expediente"`
	Antecedentes     string `json:"antecedentes"`
	HistorialClinico string `json:"historial_clinico"`
	Seguro           string `json:"seguro"`
	IdPaciente       int    `json:"id_paciente"`
}
