package db

import (
	"Red_Social/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := mongoConect.Database("twittor")
	col := db.Collection("users")

	var perfil models.User

	objId, _ := primitive.ObjectIDFromHex(id)

	//La consulta que le realizaremos a nuestra bd
	condicion := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	//Seteamos la password para no poder manejarla
	perfil.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado ", err.Error())
		return perfil, err
	}
	return perfil, nil
}
