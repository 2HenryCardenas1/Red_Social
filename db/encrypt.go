package db

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pwd string) (string, error) {
	//El costo es  la cantidad de veces que se encriptara la password
	costo := 8
	//GenerateFromPassword encriptar password
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), costo)

	return string(bytes), err

}
