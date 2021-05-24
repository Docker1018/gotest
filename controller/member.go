package controller

import (
	"context"
	"learning/model"
	"learning/mongodb"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
)

func Insert(in *model.Member) (err error) {
	coll := mongodb.ConnectDB()
	_, err = coll.InsertOne(context.TODO(), in)
	if err != nil {
		spew.Dump(err)
		return
	}
	return
}

func Find(in string) (result *model.Member, err error) {
	coll := mongodb.ConnectDB()
	f := bson.M{"name": in}
	res := coll.FindOne(context.TODO(), f)
	err = res.Decode(&result)
	if err != nil {
		return
	}
	return
}
