package main

import (
	"big-genius/core/config"
	ctx "big-genius/core/context"
	"big-genius/core/log"
	"big-genius/internal/app/controllers"
	"big-genius/internal/app/models/database"
	"big-genius/internal/app/models/mq"
	"big-genius/internal/app/models/openai"
	"big-genius/internal/app/models/redis"
	smq "big-genius/internal/app/services/mq"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {

	app := iris.New()
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

	app.Post("/wx/mock", controllers.MockWxWebhook)

	app.Get("/wx/airobot", controllers.AIRobot)
	app.Post("/wx/airobot", controllers.AIRobot)

	if err := app.Listen(":" + config.GlobalConfig.App.Port); err != nil {
		panic(err)
	}

	log.Logger.Info("server is shutdown")
}

func initComponents() {

	// 初始化配置
	config.Init()

	// 初始化日志记录器
	log.Init()

	// 初始化数据库
	database.Init()

	// 初始化Redis
	redis.Init()

	// 初始化RabbitMQ && 注册消费任务
	mq.Init()
	smq.RegisterTrigger()

	//// 初始化全局Proxy
	//proxy.Init()

	// 初始化OpenAI
	//openai.InitChatGPT()
	openai.InitAzure()
}
