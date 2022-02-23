package routers

import (
	"Red_Social/db"
	"net/http"
)

func EliminarTweet(rw http.ResponseWriter, r *http.Request) {
	IdTweet := r.URL.Query().Get("id")
	if len(IdTweet) < 1 {
		http.Error(rw, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(IdTweet, IDUsuario)

	if err != nil {
		http.Error(rw, "Ocuerio un error al borar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)

}
