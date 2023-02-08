package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuquanhao/moyu/model"
)

func Logged(c *fiber.Ctx) error {

	s, _ := model.SessionStore.Get(c)
	if s.Fresh() {
		c.Status(fiber.StatusUnauthorized)
		return nil
	}
	return c.Next()
}
