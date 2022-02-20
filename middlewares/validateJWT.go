package middlewares

import (
	"Red_Social/routers"
	"net/http"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion
Recordar que cuando realizamos middlewares debemos recibir y devolver lo mismo
*/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//r.Header.Get("Authorization") esto es un string por eso en el metodo ProcesoToken recibe
		//una variable string
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Error en el TOKEN ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
