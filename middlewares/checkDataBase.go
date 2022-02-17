package middlewares

import (
	"Red_Social/db"
	"net/http"
)

//Middleware que permite conocer el estado de la base de datos
//Como parametro recibe HandlerFunc por que necesita ver la ruta ingresada
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConection() == 0 {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
