package controller

import (
	"fmt"
	"log"
	"net/url"

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
