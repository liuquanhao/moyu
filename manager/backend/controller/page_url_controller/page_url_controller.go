package page_url_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuquanhao/moyu/model"
	"github.com/liuquanhao/moyu/model/remote_url_model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Add(c *fiber.Ctx) error {
	db := model.DBConn
	pageUrl := new(remote_url_model.PageUrl)
	if err := c.BodyParser(pageUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := validate.Struct(pageUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	db.Create(&pageUrl)
	return c.JSON(pageUrl)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := model.DBConn
	var pageUrl remote_url_model.PageUrl
	db.Take(&pageUrl, id)
	if pageUrl.Title == "" {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}
	db.Delete(&pageUrl)
	return c.SendString("Delete success")
}

func List(c *fiber.Ctx) error {
	db := model.DBConn
	var pageUrls []remote_url_model.PageUrl
	db.Find(&pageUrls)
	return c.JSON(pageUrls)
}
