package healthcheck

import (
	"github.com/mike504110403/common-moduals/fiber/handler/healthcheck"

	"github.com/gofiber/fiber/v2"
)

// SetRouter 設定路由，當在不符合使用時，可以在自己的 project 覆蓋此變數
//
// 使用方法 (在 main.go 中): healthcheck.SetRouter(router)
//
// 注意必須放在 r.Use(NotFound) 之前
var SetRouter = func(router fiber.Router) {
	router.Get("/gateway/healthcheck", healthcheck.Handler)
}
