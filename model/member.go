package model

type Member struct {
	Name     string `bson:"name"`
	Account  string `bson:"account"`
	Password string `bson:"password"`
}