package models

//Decodificar nuestro body con el tweet dentro
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
