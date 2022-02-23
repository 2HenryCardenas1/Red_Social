package db

import (
	"Red_Social/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(user models.User) (string, bool, error) {

	//ctx = context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Llamando a la base de datos
	db := MongoConect.Database("twittor")
	//Llamando a la coleccion usuario
	col := db.Collection("users")

	//Encriptacion de password
	user.Password, _ = EncryptPassword(user.Password)

	//result es un json
	//El InsertOne es el metodo para insertar datos en mogodb
	resul, err := col.InsertOne(ctx, user)

	//Si salio error devolvemos un json vacio un false y el error
	if err != nil {
		return "", false, err
	}

	//Forma de obtener el ID que se acaba de insertar
	ObjID, _ := resul.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}
