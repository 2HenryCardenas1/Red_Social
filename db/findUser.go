package db

import (
	"Red_Social/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserFind(email string) (models.User, bool, string) {
	//Esperar un tiempo de respuesta
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Llamando a la base de datos
	db := MongoConect.Database("twittor")
	//Llamando a la coleccion usuario
	col := db.Collection("users")

	//bson.M nos devuelve en formato JSON
	//La consulta que le realizaremos a nuestra bd
	condicion := bson.M{
		"email": email,
	}

	var resultado models.User

	/*Si no me encuentra el usuario guarda el error en la variable err
	y si lo encuentra lo decodifica y lo guarda en la variable RESULTADO
	*/
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	//Converitmos el parametro ID en string tipo hexadecimal
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	//Si encontro al usurio lo devulve
	return resultado, true, ID

}
