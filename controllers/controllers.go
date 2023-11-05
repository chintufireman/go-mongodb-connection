package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://chintufireman:root@cluster0.fjqreg9.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMP
var collection *mongo.Collection

// connect with mongodb
func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection was successfull")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}
