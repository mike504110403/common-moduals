package notfound

import (
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/loghttp"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/tracehandler"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/ilog/v2"

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
