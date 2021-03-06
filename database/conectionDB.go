package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConnectDatabase()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user1:contraseña1@cluster-careerpath.1yinv.mongodb.net/test")

/*ConectarBD es la función que me permite conectar a la base de datos*/
func ConnectDatabase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection to the Database")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

//prueba de agregar algo.
