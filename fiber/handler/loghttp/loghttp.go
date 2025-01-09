package loghttp

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"google.golang.org/api/logging/v2"
)

func Get(c *fiber.Ctx) *logging.HttpRequest {
	h := &logging.HttpRequest{
		ResponseSize:  int64(len(c.Response().Body())),
		Status:        int64(c.Response().StatusCode()),
		Latency:       time.Since(c.Context().Time()).String(),
		Protocol:      utils.GetString(c.Request().Header.Protocol()),
		Referer:       c.Get(fiber.HeaderReferer),
		RemoteIp:      c.Get(fiber.HeaderXForwardedFor),
		RequestMethod: c.Method(),
		RequestSize:   int64(len(c.Request().Body())),
		RequestUrl:    c.Path(),
		UserAgent:     c.Get(fiber.HeaderUserAgent),
	}
	return h
}
