package utils

import "unicode"
import "fmt"


func EsPasswordSegura(password string) bool {
	if len(password) < 12 {
		return false
	}

	var (
		hasUpper  bool
		hasNumber bool
		hasSymbol bool
	)

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsNumber(char) {
			hasNumber = true
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSymbol = true
		}
	}

	fmt.Printf("Seguridad: upper=%v, num=%v, sym=%v, len=%v\n", hasUpper, hasNumber, hasSymbol, len(password))
	return hasUpper && hasNumber && hasSymbol
}
