package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name"`
	Age    int                `json:"age"`
	Status bool               `json:"status"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("main: connot create new client. : ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("main: connot connect to database.")
	}

	collection := client.Database("golang").Collection("users")
	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFindOne()
	filter := bson.M{"name": "pi"}
	singleResult := collection.FindOne(ctx, filter)

	if err != nil {
		log.Fatal("main : findOne : ", err)
	}

	result := Users{}
	singleResult.Decode(&result)

	resultByte, err := json.Marshal(result)
	if err != nil {
		log.Fatal("main : unabled to marshal json: ", result)
	}

	fmt.Println(string(resultByte))
}
