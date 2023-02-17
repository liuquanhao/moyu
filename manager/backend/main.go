package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/liuquanhao/moyu/controller/page_data_controller"
	"github.com/liuquanhao/moyu/controller/page_url_controller"
	"github.com/liuquanhao/moyu/controller/user_controller"
	"github.com/liuquanhao/moyu/middleware"
	"github.com/liuquanhao/moyu/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/gofiber/websocket/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getCurrentPwd() string {
	path, _ := os.Getwd()
	return path
}
func initDatabase() {
	var err error
	model.DBConn, err = gorm.Open(sqlite.Open(getCurrentPwd() + "/../db/moyu_manager.db"))
	if err != nil {
		log.Fatal("failed to connect database: {}", err)
	}
}

func initSession() {
	storage := sqlite3.New(sqlite3.Config{
		Database:        getCurrentPwd() + "/../db/moyu_manager.db",
		Table:           "sessions",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})

	model.SessionStore = session.New(session.Config{
		Storage: storage,
	})
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", user_controller.Login)
	api.Get("/page_url", page_url_controller.List)

	logged := app.Group("/api", middleware.Logged)
	logged.Post("/logout", user_controller.Logout)
	logged.Get("/account", user_controller.Account)
	logged.Post("/page_url", page_url_controller.Add)
	logged.Delete("/page_url/:id", page_url_controller.Delete)

	app.Get("/ws/page_data", websocket.New(page_data_controller.PushPageData))
}

//go:embed dist
var frontend embed.FS

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://194.29.186.121:8080",
		AllowCredentials: true,
	}))
	app.Use("/ws/*", middleware.UpgradeOptions)

	initDatabase()
	initSession()
	setupRoutes(app)

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
