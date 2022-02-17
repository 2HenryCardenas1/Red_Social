package routers

import (
	"Red_Social/models"
	"encoding/json"
	"net/http"

	db "Red_Social/db"
)

/*Registro es la funcion para crear en la bd el registro de usuario */
func Registro(rw http.ResponseWriter, r *http.Request) {

	var usr models.User                         //Traemos los valores del modelo
	err := json.NewDecoder(r.Body).Decode(&usr) //Decodificar lo que viene por post en json

	//Validaciones
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(usr.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", 400)
		return
	}
	if len(usr.Password) <= 6 {
		http.Error(rw, "La contraseña debe de tener al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := db.CheckUserFind(usr.Email)
	/* encontrado solo es igual a escribir encontrado == true*/
	if encontrado {
		http.Error(rw, "Ya existe un usuario con este email", 400)
		return
	}

	_, status, err := db.InsertRegister(usr)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar realizar el registro de usurio"+err.Error(), 400)
		return
	}
	/* !status es igual a decir status == false*/
	if !status {
		http.Error(rw, "Ocurrio un error al registrar el usuario ", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
