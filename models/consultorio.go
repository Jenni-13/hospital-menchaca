package models

type Consultorio struct {
	IdConsultorio int    `json:"id_consultorio"`
	Tipo          string `json:"tipo"`
	Ubicacion     string `json:"ubicacion"`
	Nombre        string `json:"nombre"`
	IdMedico      int    `json:"id_medico"`
}
