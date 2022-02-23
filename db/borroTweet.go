package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(IdTweet string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConect.Database("twittor")

	col := db.Collection("tweet")

	//Convertimos el id que entra por parametro en un ObjectID
	objIdTweet, _ := primitive.ObjectIDFromHex(IdTweet)

	condicion := bson.M{
		"_id":    objIdTweet,
		"userid": UserId,
	}

	_, err := col.DeleteOne(ctx, condicion)

	return err
}
