package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPagingParams(ctx *fiber.Ctx) (perPage int, page int) {
	perPage, _ = strconv.Atoi(ctx.Query("perpage"))
	page, _ = strconv.Atoi(ctx.Query("page"))
	return
}
