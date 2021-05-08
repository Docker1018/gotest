package main

import (
	"learning/mongodb"

	"github.com/kataras/iris/v12"
)

func main() {
	// mongodb.ListAll()
	// app := iris.New()

	// memberApi := app.Party("/member")
	// {
	// 	memberApi.Use(iris.Compression)

	// 	// GET: http://localhost:8080/member
	// 	memberApi.Get("/", list)

	// 	// POST: http://localhost:8080/member
	// 	memberApi.Post("/", create)

	// 	// PUT: http://localhost:8080/member
	// 	memberApi.Put("/", update)

	// 	// DELETE: http://localhost:8080/member
	// 	memberApi.Delete("/", delete)

	// }

	// app.Listen(":8080")
	mongodb.Hi()
	mongodb.ConnectDB()
}

type Member struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Account  string `bson:"account"`
	Password string `bson:"password"`
}

func list(ctx iris.Context) {
	// member := []Member{
	// 	{"1111", "name1", "acc1", "pass1"},
	// 	{"2222", "name2", "acc2", "pass2"},
	// 	{"3333", "name3", "acc3", "pass3"},
	// 	{"4444", "name4", "acc4", "pass4"},
	// }
	// ctx.JSON(member)
	// mongodb.ListAll()
}

func create(ctx iris.Context) {
 var m Member
 err := ctx.ReadJSON(&m)

 if err != nil {
	 ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("member create flail").DetailErr(err))
	 return
 }
 println("member" + m.Name)

 ctx.StatusCode(iris.StatusCreated)
}

func update(ctx iris.Context) {

}

func delete(ctx iris.Context) {

}
