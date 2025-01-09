package healthcheck

import "github.com/gofiber/fiber/v2"

// healthcheck 的 Handler，當不符合使用時，可在 project 覆蓋此變數
var Handler = func(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
