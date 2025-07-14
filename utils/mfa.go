package utils

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// GenerarMFA genera un secreto MFA y su URL para escanear con Google Authenticator
func GenerarMFA(nombreUsuario string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Hospital App",
		AccountName: nombreUsuario,
		Algorithm:   otp.AlgorithmSHA1,
		Period:      30,
		Digits:      otp.DigitsSix,
	})
	if err != nil {
		return "", "", err
	}

	// Retorna el secreto y la URL del QR
	return key.Secret(), key.URL(), nil
}
