package gracefulshutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	mlog "github.com/mike504110403/goutils/log"
)

func Listen(app *fiber.App, port string) {
	go func() {
		if err := app.Listen(port); err != nil {
			mlog.Error(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)      // Create channel to signify a signal being sent
	signal.Notify(quit, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	_ = <-quit                           // This blocks the main thread until an interrupt is received
	mlog.Info(fmt.Sprintf("Wait Shutdown for Port: %s", port))
	time.Sleep(15 * time.Second)
	mlog.Info(fmt.Sprintf("Server Start Shutdown for Port: %s", port))
	if err := app.Shutdown(); err != nil {
		mlog.Error(fmt.Sprintf("Server Shutdown Error: %s", err))
	}
	mlog.Info(fmt.Sprintf("Port %s Server Shutdown", port))
}
