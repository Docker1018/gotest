package main

import (
	"learning/controller"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	memberApi := app.Party("/member")
	{
		memberApi.Use(iris.Compression)

		// GET: http://localhost:8080/member
		memberApi.Get("/", list)

		// POST: http://localhost:8080/member
		memberApi.Post("/", list)

	}

	app.Listen(":8080")
}

func list(ctx iris.Context) {
	// member := []Member{
	// 	{"1111", "name1", "acc1", "pass1"},
	// 	{"2222", "name2", "acc2", "pass2"},
	// 	{"3333", "name3", "acc3", "pass3"},
	// 	{"4444", "name4", "acc4", "pass4"},
	// }
	// ctx.JSON(member)
	res := controller.Insert()
	ctx.JSON(res)
}
