package db

import (
	"Red_Social/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Modificar perfil
func ModificarRegistro(usr models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Llamando a la base de datos
	db := MongoConect.Database("twittor")
	//Llamando a la coleccion usuario
	col := db.Collection("users")

	registro := make(map[string]interface{})

	if len(usr.Name) > 0 {
		registro["name"] = usr.Name
	}
	if len(usr.LastName) > 0 {
		registro["lastName"] = usr.LastName
	}

	registro["dateBird"] = usr.DateBirth

	if len(usr.Email) > 0 {
		registro["email"] = usr.Email
	}
	if len(usr.Avatar) > 0 {
		registro["avatar"] = usr.Avatar
	}
	if len(usr.Baner) > 0 {
		registro["baner"] = usr.Baner
	}
	if len(usr.Biografia) > 0 {
		registro["biografia"] = usr.Biografia
	}
	if len(usr.Ubicacion) > 0 {
		registro["ubicacion"] = usr.Ubicacion
	}
	if len(usr.WebSite) > 0 {
		registro["webSite"] = usr.WebSite
	}

	//Damos valores al  registro que actualizara
	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	//Aca le hacemos la condicion para que me actualice el usuario con la ID
	/*Esto seria igual a un: update table set nombre='nombre' where ID = ID */
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
