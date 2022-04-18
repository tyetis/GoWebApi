package controller

import (
	"github.com/gofiber/fiber/v2"
)

type postController struct {
}

var PostController postController = postController{}

func (c *postController) Create(ctx *fiber.Ctx) error {
	return ctx.JSON("Create")
}

func (c *postController) Index(ctx *fiber.Ctx) error {
	// sess := utils.GetSession(ctx)
	// name := sess.Get("name")
	// sess.Set("name", "john")

	// if err := sess.Save(); err != nil {
	// 	panic(err)
	// }

	// posts := services.PostService.FindOne()

	// return ctx.JSON(fiber.Map{
	// 	"Page":      "Posts",
	// 	"user name": name,
	// 	"post name": posts[0].Name,
	// })
	return nil
}
