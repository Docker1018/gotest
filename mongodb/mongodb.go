package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Hi() {
	fmt.Println("?")
}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}

	collection := client.Database("test").Collection("member")
	fmt.Println(collection)
	fmt.Println("ddd")
}

// func ListAll() {
// 	res, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("fasf")
// 	fmt.Println(res)
// }
