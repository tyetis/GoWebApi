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
	var user entity.User
	ctx.BodyParser(&user)

	return ctx.JSON(utils.Validator(user, func() interface{} {
		return c.service.Create(&user)
	}))
}

func (c *userController) Update(ctx *fiber.Ctx) error {
	var user entity.User
	ctx.BodyParser(&user)
	id, _ := strconv.Atoi(ctx.Params("id"))

	return ctx.JSON(utils.Validator(user, func() interface{} {
		user.Id = id
		return c.service.Update(&user)
	}))
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

func (c *userController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return ctx.JSON(c.service.Delete(id))
}

func SetSession() {
	// sess := utils.GetSession(ctx)
	// name := sess.Get("name")
	// sess.Set("name", "john")

	// if err := sess.Save(); err != nil {
	// 	panic(err)
	// }
}
