package routers

import (
	"Red_Social/db"
	"encoding/json"
	"net/http"
)

func VerPerfil(rw http.ResponseWriter, r *http.Request) {
	//Vamos a traer del body los parametros que nos llegaron

	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(rw, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(id)

	if err != nil {
		http.Error(rw, "Ocurrio un error al buscar el registro "+err.Error(), 400)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(perfil)
}
