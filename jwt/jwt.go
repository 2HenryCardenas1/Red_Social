package jwt

import (
	"Red_Social/models"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

//Recordar que los JWT tienen tres partes el HEADER-PAYLOAD-NUESTRA FIRMA
func GeneroJWT(token models.User) (string, error) {
	//LLave privada -FIRMA

	miClave := []byte("Esta_es_mi_clave_privada")
	//Asignamos los valores que llevara el token - PAYLOAD
	payload := jwt.MapClaims{
		"email":     token.Email,
		"name":      token.Name,
		"lastName":  token.LastName,
		"dateBirth": token.DateBirth,
		"biografia": token.Biografia,
		"ubicacion": token.Biografia,
		"webSite":   token.WebSite,
		"_id":       token.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	//Creamos un token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//Ponemos la firma de nosotros
	tokensStr, err := newToken.SignedString(miClave)

	if err != nil {
		return tokensStr, err
	}

	return tokensStr, nil
}
