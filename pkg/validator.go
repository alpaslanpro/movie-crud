package pkg

import (
	_ "regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate


func init() {
	Validate = validator.New()
	//Validate.RegisterValidation("password_complex", passwordValidator)
}
/*
func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*()_+{}:;.,<>?-]`).MatchString(password)

	return len(password) >= 8 && hasUpper && hasNumber && hasSpecial
}
*/