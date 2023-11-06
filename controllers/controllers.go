package controller

import (
	"context"
	"fmt"
	"log"
	model "mongotut/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//Mongodb helpers - file

//insert 1 record
func insertOneMovie(movie model.Netflix)  {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updatedOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update:= bson.M{"set":bson.M{"watched":true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ",result.ModifiedCount)
}