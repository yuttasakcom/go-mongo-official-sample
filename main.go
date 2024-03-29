package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/yuttasakcom/go-mongo-official-sample/users"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("main: init(): Error loading .env file")
	}
}

func main() {
	filter := bson.M{"name": "yoo"}
	user, err := users.FindOne(filter)
	if err != nil {
		log.Fatalf("main: users.FindOne fail: %v", err)
	}

	fmt.Println(string(user))
}
