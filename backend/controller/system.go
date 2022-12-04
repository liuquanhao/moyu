package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/liuquanhao/moyu/service"
)

func GetSysInfo(c *fiber.Ctx) error {
	return c.JSON(service.GetSystemInfo())
}

func GetSysStatus(c *fiber.Ctx) error {
	return c.JSON(service.GetSystemStatus())
}
func PushSysStatus(c *websocket.Conn) {
	for {
		c.WriteJSON(service.GetSystemStatus())
		time.Sleep(1 * time.Second)
	}
}
