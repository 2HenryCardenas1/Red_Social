package routers

import (
	"Red_Social/db"
	"Red_Social/models"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

//Variables exportadas que se usaran en todos los endPoints (routers)

var Email string
var IDuser string

/*Proceso token para extraer sus valores*/
func ProcesoToken(token string) (*models.Claim, bool, string, error) {

	miClave := []byte("Esta_es_mi_clave_privada")
	// el &var es para poner un puntero
	claims := &models.Claim{}

	/*Separamos la palabra Bearer ya que al crear el token nos lo crea con esto al inicio
	es un estandar para la autenticacion de usuarios */
	splitToken := strings.Split(token, "Bearer")

	//Verificamos si se separo el Bearer de nuestro token
	//errors.New con esta funcion creamos nuestros propio errores
	if len(splitToken) != 2 {
		errores := errors.New("formato de token invalido")
		return claims, false, string(""), errores
	}
	//TrimSpace con esta funcion le quita los espacios al token
	token = strings.TrimSpace(splitToken[1])
	//Validacion token
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return miClave, nil })

	if err == nil {
		_, encontrado, _ := db.CheckUserFind(claims.Email)
		if encontrado {
			Email = claims.Email
			IDuser = claims.ID.Hex()
		}
		return claims, encontrado, IDuser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
