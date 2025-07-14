package models

type Expediente struct {
    IdExpediente   int     `json:"id_expediente"`
    NombrePaciente string  `json:"nombre_paciente"` // Añadir este campo para el join
    Antecedentes   string  `json:"antecedentes"`
    HistorialClinico string `json:"historial_clinico"`
    Seguro         string  `json:"seguro"`
    IdPaciente     int     `json:"id_paciente"`
    Estatura       float64 `json:"estatura"`
    Peso           float64 `json:"peso"`
    Edad           int     `json:"edad"`
}

