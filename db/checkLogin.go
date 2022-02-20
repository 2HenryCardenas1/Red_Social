package db

import (
	"Red_Social/models"

	"golang.org/x/crypto/bcrypt"
)

/*Funcion CheckLogin, retorna datos de usuraio(models.User) y un estado booleano*/
func CheckLogin(email string, password string) (models.User, bool) {
	user, encontrado, _ := CheckUserFind(email)

	if !encontrado {
		return user, false
	}

	passwordBytes := []byte(password)   //pwd texto plano
	passwordDB := []byte(user.Password) //pwd encriptada

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
