package models

type Usuario struct {
	IdUsuario int    `json:"id_usuario"`
	Nombre    string `json:"nombre"`
	Rol       string `json:"rol"`
	Password  string `json:"password"`
	Correo    string `json:"correo"`
	MfaSecret  string  `json:"mfa_secret"`
}
