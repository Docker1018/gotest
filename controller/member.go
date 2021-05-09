package controller

import (
	"context"
	"fmt"
	"learning/model"
	"learning/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insert() (res *mongo.InsertOneResult) {
	fmt.Println("insert")
	coll := mongodb.ConnectDB()
	data := &model.Member{Name: "nameee", Account: "Accccc", Password: "passwordddd"}
	result, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
