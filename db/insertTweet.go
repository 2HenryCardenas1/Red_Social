package db

import (
	"Red_Social/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Llamando a la base de datos
	db := MongoConect.Database("twittor")
	//Llamando a la coleccion tweet
	col := db.Collection("tweet")

	//convertimos el json a bson

	registro := bson.M{
		"userid":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	//convertimos el id del tweet en string
	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil

}
