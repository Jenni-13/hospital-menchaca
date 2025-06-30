package models

type Usuario struct {
	IdUsuario int    `json:"id_usuario"`
	Nombre    string `json:"nombre"`
	Rol       string `json:"rol"`
}
