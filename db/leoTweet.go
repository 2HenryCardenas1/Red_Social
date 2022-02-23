package db

import (
	"Red_Social/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//recibe un id de usuario y una pagina(pagiandor)
func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConect.Database("twittor")

	col := db.Collection("tweet")

	//Esta variable guardara los tweets encontrados
	var resultados []*models.DevuelvoTweets
	//La condicion es la consulta a la base de datos, aca traemos los tweets los cuales sean del usuario registrado
	condicion := bson.M{
		"userid": ID,
	}

	//CRACION PAGINADOR

	//con la funcion options podemos interactuar con los datos de la db
	opciones := options.Find()
	//Con setLimit ponemos el limite de archivos que traera
	opciones.SetLimit(20)
	//Con Sort le estamos diciendo que nos organice segun la key fecha en orden descendente,Se usa bson.D
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	//Va a tomar un valor entero el caul sera el total de valores por pagina que apareceran
	//Este es nuestro pagiandor
	opciones.SetSkip((pagina - 1) * 20)

	//Variable que funciona como puntero para recorrer los resultados

	puntero, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	//Con next le indicamos que me busque cada ves un elemento de la variable puntero
	for puntero.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := puntero.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)

	}

	return resultados, true
}
