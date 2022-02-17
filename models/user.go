package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Modelo de user
type User struct {
	//La parte del bson es para indicarle la estructura del dato a mongo db
	//Con omitempty no mostramos los datos de salida
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"` //Le pusimos el omitempty en el json por si no lo encuentra
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	DateBirth time.Time          `bson:"dateBird" json:"dateBird,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Baner     string             `bson:"baner" json:"baner,omitempty"`
	Biografia string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
