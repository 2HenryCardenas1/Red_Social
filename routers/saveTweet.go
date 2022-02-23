package routers

import (
	"Red_Social/db"
	"Red_Social/models"
	"encoding/json"
	"net/http"
	"time"
)

func GraboTweet(rw http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	//recibimos el json en nuestro body y lo decodificamos

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	//Validaciones
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	registro := models.SaveTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(registro)

	if err != nil {
		http.Error(rw, "No se pudo subir el tweet intenta de nuevo "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "No se pudo subir el tweet intenta de nuevo ", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
