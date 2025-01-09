package notfound

import (
	"github.com/mike504110403/common-moduals/fiber/handler/loghttp"
	"github.com/mike504110403/common-moduals/fiber/handler/tracehandler"
	"github.com/mike504110403/common-moduals/ilog/v2"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)
	return func(c *fiber.Ctx) error {
		notFound := fiber.ErrNotFound
		c.Status(notFound.Code)
		if cfg.LogNotFound {
			ilog.HTTP(loghttp.Get(c)).
				Trace(tracehandler.Get(c)).
				Err("[HTTP_%d] %s", notFound.Code, notFound.Message)
		}
		return fiber.ErrNotFound
	}
}
