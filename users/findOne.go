package users

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yuttasakcom/go-mongo-official-sample/database"
	"go.mongodb.org/mongo-driver/bson"
)

// FindOne user
func FindOne(b bson.M) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("users: FindOne(): database.Connect fail: %v", err)
	}

	result := Users{}
	collection := db.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("COLLECTION_USERS"))
	collection.FindOne(ctx, b).Decode(&result)

	resultByte, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("users: FindOne(): unabled to marshal json: %v", err)
	}

	return resultByte, nil
}
