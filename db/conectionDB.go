package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConect = ConectarBD()

//Esta es nuestra llamada a la bd
var clientOptions = options.Client().ApplyURI("Aca ponen la direccion que le da mongodb")

//Conectamos la base de datos
func ConectarBD() *mongo.Client {
	/*En esta linea estamos dicendo que si no responde la bd en 10 segundos
	entonces que me mande de una vez a cancelar.
	contexT me guarda TODO el contexto valido, es como si se llamara de la siguiente forma
	context.TODO()
	*/
	contexT, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/*Realizamos la conexion y si es exitosa nos guardara el contenido de exito en la variable
	client de lo contrario procedera a mostrar los errores los cuales los capturo en la variable err*/
	client, err := mongo.Connect(contexT, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(contexT, nil) //revisar si la bd esta encendida
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion existosa a la bd :)")

	return client

}

//Verificamos la conexion
func CheckConection() int {
	contexT, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := MongoConect.Ping(contexT, nil) //revisar si la bd esta encendida

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
