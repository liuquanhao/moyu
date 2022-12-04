package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/liuquanhao/moyu/controller"
	"github.com/liuquanhao/moyu/middleware"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(cors.New())
	app.Use("/ws/*", middleware.UpgradeOptions)

	app.Get("/sys_info", controller.GetSysInfo)
	app.Get("/sys_status", controller.GetSysStatus)
	app.Get("/ws/sys_status", websocket.New(controller.PushSysStatus))

	app.Listen("0.0.0.0:8080")
}
