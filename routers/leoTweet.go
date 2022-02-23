package routers

import (
	"Red_Social/db"
	"encoding/json"
	"net/http"
	"strconv"
)

func LeoTweet(rw http.ResponseWriter, r *http.Request) {
	// vamos a obtener el id del tweet por url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	//Comprovamos que nuestra pagina existe
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(rw, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	//Convertimos a entero lo que nos pongan en valor de pagina
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(rw, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	//Convertimos a entero lo que nos pongan en valor de pagina
	pag := int64(pagina)

	respuesta, status := db.LeoTweet(ID, pag)

	if !status {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	//Editamos nuestro header para que acepte un formato json
	rw.Header().Set("Content-Type", "application/json")
	//Escribimos en nuestro header
	rw.WriteHeader(http.StatusCreated)
	//mandamos los datos optenidos de la consulta de LeoTweet
	json.NewEncoder(rw).Encode(respuesta)

}
