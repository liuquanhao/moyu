package user_controller

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/liuquanhao/moyu/model"
	"github.com/liuquanhao/moyu/model/user_model"
)

var validate = validator.New()

func Login(c *fiber.Ctx) error {
	db := model.DBConn
	user := new(user_model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	db.Where(&user).Take(&user)
	if user.Id == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Found User")
	}
	s, _ := model.SessionStore.Get(c)
	s.Set("uid", user.Id)
	s.Set("sid", s.ID())
	s.Set("ip", c.Context().RemoteIP().String())
	s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
	s.Set("ua", string(c.Request().Header.UserAgent()))
	err := s.Save()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "uid"
	cookie.Value = strconv.FormatUint(uint64(user.Id), 10)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)
	c.Status(fiber.StatusNoContent)
	return nil
}

func Account(c *fiber.Ctx) error {
	s, _ := model.SessionStore.Get(c)
	if s.Fresh() {
		c.Status(fiber.StatusNotFound)
		return nil
	}
	return c.JSON(fiber.Map{
		"sid":   s.ID(),
		"uid":   s.Get("uid"),
		"ip":    s.Get("ip"),
		"login": s.Get("login"),
		"ua":    s.Get("ua"),
	})
}

func Logout(c *fiber.Ctx) error {
	s, _ := model.SessionStore.Get(c)
	if err := s.Destroy(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	c.Status(fiber.StatusNoContent)
	return nil
}
