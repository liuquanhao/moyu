package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	exePath, _ := os.Executable()
	return filepath.Dir(exePath)
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))
	base := app.Group(os.Getenv("BASEURL"))

	ws := base.Group("/ws")
	ws.Use("/*", middleware.UpgradeOptions)
	ws.Get("/page_data", websocket.New(page_data_controller.PushPageData))

	api := base.Group("/api")
	api.Post("/login", user_controller.Login)
	api.Get("/page_url", page_url_controller.List)

	logged := app.Group("/api", middleware.Logged)
	logged.Post("/logout", user_controller.Logout)
	logged.Get("/account", user_controller.Account)
	logged.Post("/page_url", page_url_controller.Add)
	logged.Delete("/page_url/:id", page_url_controller.Delete)

	stripped, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal(err)
	}
	base.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(stripped),
		Browse: true,
	}))
}

//go:embed dist
var frontend embed.FS

func main() {
	app := fiber.New()

	initDatabase()
	initSession()
	setupRoutes(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Println(host + ":" + port)
	log.Fatal(app.Listen(host + ":" + port))
}
