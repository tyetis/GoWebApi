package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tyetis/goapiexample/app/data/entity"
	"github.com/tyetis/goapiexample/app/services"
	"github.com/tyetis/goapiexample/app/utils"
)

type userController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) userController {
	return userController{
		service: service,
	}
}

func (c *userController) Create(ctx *fiber.Ctx) error {
	user := c.service.FindOne(&entity.User{})
	return ctx.JSON(user)
}

func (c *userController) Index(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	perPage, page := utils.GetPagingParams(ctx)
	name := ctx.Query("name")
	email := ctx.Query("email")

	users := c.service.GetAll(&entity.User{
		Id:    id,
		Name:  name,
		Email: email,
	}, perPage, page)

	return ctx.JSON(users)
}

func SetSession() {
	// sess := utils.GetSession(ctx)
	// name := sess.Get("name")
	// sess.Set("name", "john")

	// if err := sess.Save(); err != nil {
	// 	panic(err)
	// }
}
