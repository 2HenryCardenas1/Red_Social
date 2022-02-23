package routers

import (
	"Red_Social/db"
	"Red_Social/models"
	"encoding/json"
	"net/http"
)

func EditProfile(rw http.ResponseWriter, r *http.Request) {

	var usr models.User

	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		http.Error(rw, "Datos incorrectos "+err.Error(), 400)
		return
	}

	//IDuser es nuestra variable global la cual traemos cuando el token se ha verificado
	//Esta en el archivo processTooken.go
	var status bool

	status, err = db.ModificarRegistro(usr, IDUsuario)

	if err != nil {
		http.Error(rw, "Ocurrio un error al actualizar el usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "No se ha logrado modificar el registro "+err.Error(), 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
