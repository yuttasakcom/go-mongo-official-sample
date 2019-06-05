package users

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yuttasakcom/go-mongo-official-sample/database"
	"github.com/yuttasakcom/go-mongo-official-sample/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindOne user
func FindOne(b bson.M) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.Connect(ctx)
	if err != nil {
		return "", fmt.Errorf("users: FindOne(): %v", err)
	}

	collection := db.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("COLLECTION_USERS"))
	singleResult := collection.FindOne(ctx, b)

	result := models.Users{}
	singleResult.Decode(&result)

	resultByte, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("users: FindOne(): unabled to marshal json: %v", err)
	}

	return string(resultByte), nil
}
