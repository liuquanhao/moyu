package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"
	"github.com/liuquanhao/moyu/controller"
	"github.com/liuquanhao/moyu/middleware"
)

//go:embed dist
var frontend embed.FS

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(cors.New())

	app.Use("/ws/*", middleware.UpgradeOptions)

	app.Get("/sys_info", controller.GetSysInfo)
	app.Get("/sys_status", controller.GetSysStatus)
	app.Get("/api/page_data", controller.GetPageData)
	app.Get("/ws/sys_status", websocket.New(controller.PushSysStatus))

	stripped, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(stripped),
		Browse: true,
	}))

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Println(host + ":" + port)
	log.Fatal(app.Listen(host + ":" + port))
}
