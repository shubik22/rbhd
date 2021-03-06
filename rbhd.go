package main

import (
	"os"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"

	"github.com/shubik22/rbhd/lib/app"
	"github.com/shubik22/rbhd/lib/handlers"
)

func main() {
	a := app.NewApp()
	go a.Init()

	iris.Use(logger.New(iris.Logger))
	iris.Use(recovery.New())

	iris.Get("/users", handlers.UsersHandlerDev)
	iris.Get("/leaderboard", func(ctx *iris.Context) {
		handlers.LeaderboardHandler(a, ctx)
	})

	iris.Static("/assets", "./static_files/", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("./static_files/index.html", false)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	iris.Listen(port)
}
