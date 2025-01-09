package recoverhandler

import (
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

type Recover struct {
	Recover interface{}
	Stack   []byte
}

// Error implements error
func (r *Recover) Error() string {
	return fmt.Sprintf("%v\n%s", r.Recover, r.Stack)
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) (e error) {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				e = &Recover{
					Recover: r,
					Stack:   debug.Stack(),
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
