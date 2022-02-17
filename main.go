package main

import (
	contection "Red_Social/db"
	handler "Red_Social/handlers"
	"log"
)

func main() {
	/*Revisamos la conexion a la bd*/

	if contection.CheckConection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	/*una vez chequeada la conexion realice la verificacion de si se puede comunicar
	con el puerto especifico y esa direccion puede comunicarce con la API
	*/
	handler.Managers()

}
