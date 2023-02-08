package main

import (
	"log"
	"os"
	"time"

	"github.com/liuquanhao/moyu/controller/page_url_controller"
	"github.com/liuquanhao/moyu/controller/user_controller"
	"github.com/liuquanhao/moyu/middleware"
	"github.com/liuquanhao/moyu/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	model.DBConn, err = gorm.Open(sqlite.Open("./database/moyu_manager.db"))
	if err != nil {
		log.Fatal("failed to connect database: {}", err)
	}
}

func initSession() {
	storage := sqlite3.New(sqlite3.Config{
		Database:        "./database/moyu_manager.db",
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

	logged := app.Group("/api", middleware.Logged)
	logged.Post("/logout", user_controller.Logout)
	logged.Get("/account", user_controller.Account)
	logged.Post("/page_url", page_url_controller.Add)
	logged.Delete("/page_url/:id", page_url_controller.Delete)
	logged.Get("/page_url", page_url_controller.List)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	initSession()
	setupRoutes(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(host + ":" + port))
}
