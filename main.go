package main

import (
	"big-genius/core/config"
	ctx "big-genius/core/context"
	"big-genius/core/log"
	trace "big-genius/core/middleware"
	"big-genius/internal/app/controllers"
	"big-genius/internal/app/models/database"
	"big-genius/internal/app/models/openai"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {

	app := iris.New()
	app.Use(trace.Inject)

	initComponents()

	// Index
	app.Get("/", func(c *context.Context) {
		controllers.Index(ctx.Context{
			Context: c,
		})
	})

	app.Get("/ask", func(c *context.Context) {
		controllers.Ask(ctx.Context{
			Context: c,
		})
	})

	app.Post("/ask", func(c *context.Context) {
		controllers.Ask(ctx.Context{
			Context: c,
		})
	})

	if err := app.Listen(":" + config.GlobalConfig.App.Port); err != nil {
		panic(err)
	}

	log.Logger.Info("server is shutdown")
}

func initComponents() {
	// 初始化配置
	config.InitConfig()

	// 初始化日志记录器
	log.InitLogger()

	// 初始化数据库
	database.DatabaseInit()

	// 初始化数据库
	openai.ChatGPTInit()
}
