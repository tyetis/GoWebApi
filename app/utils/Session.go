package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func GetSession(ctx *fiber.Ctx) *session.Session {
	sess, _ := store.Get(ctx)
	return sess
}
