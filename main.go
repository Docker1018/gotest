package main

import (
	"learning/controller"
	"learning/model"

	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	memberApi := app.Party("/member")
	{
		memberApi.Use(iris.Compression)

		// GET: http://localhost:8080/member
		memberApi.Get("/{name}", find)

		// POST: http://localhost:8080/member
		memberApi.Post("/", insert)

	}

	app.Listen(":8080")
}

func insert(ctx iris.Context) {
	m := model.Member{}
	err := ctx.ReadJSON(&m)
	if err != nil {
		spew.Dump(err)
		spew.Dump("err")
	}
	res := controller.Insert(&m)
	if res != nil {
		return
	}
	ctx.JSON(map[string]interface{}{
		"msg": "goood!",
	})
}

func find(ctx iris.Context) {
	name := ctx.Params().Get("name")
	spew.Dump(name)
	res, err := controller.Find(name)
	if err != nil {
		return
	}
	ctx.JSON(res)
}
