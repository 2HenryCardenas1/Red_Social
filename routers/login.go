package routers

import (
	"Red_Social/db"
	"Red_Social/jwt"
	"Red_Social/models"
	"encoding/json"
	"net/http"
	"time"
)

/*Se realiza el login*/
func Login(rw http.ResponseWriter, r *http.Request) {
	//Seteamos el header para indicarle que el contenido va a ser de tipo JSON

	rw.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user) //Decodificar lo que viene en json
	//Validaciones
	if err != nil {
		http.Error(rw, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(rw, "El email es requerido", 400)
		return
	}

	//Validacion de usuario
	documento, existe := db.CheckLogin(user.Email, user.Password)

	if !existe {
		http.Error(rw, "Usuario y/o contraseña invalidos ", 400)
		return
	}

	//Uso de JWT JsonWebToken

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	//Aca seteamos el header
	rw.Header().Set("content-type", "application/json")
	//Editamos la cabecera con nuestro token generado
	rw.WriteHeader(http.StatusCreated)
	//Encodeamos nuestro token
	json.NewEncoder(rw).Encode(resp)

	//Como se graba una cookie desde el backend

	//generar un campo fecha para la expiracion de la cookie
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
