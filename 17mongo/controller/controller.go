package controller

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/abc/mongo/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const username = "agarwalshivam834"
const passowrd = "Ya@300803"

var connectionString = "mongodb+srv://" + url.QueryEscape(username) + ":" + url.QueryEscape(passowrd) + "@cluster0.rqsus3e.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
		fmt.Printf("error : %v ", err)
	}

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// insert a record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update a movie

func updateOneMovie(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete a movie

func deleteOneMovie(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got delete with the delete count: ", deleteCount)
}

// delete all movies

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
