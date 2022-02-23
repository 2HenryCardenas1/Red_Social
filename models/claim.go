package models

import (
	jwt "github.com/dgrijalva/jwt-go/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim es la estructura usada para procesar los JWT
recordar que el ID es el ID que nos recibe desde la bd
no podemos decirle que en la bd se llama _id y aca se va a llamar solo id
*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
