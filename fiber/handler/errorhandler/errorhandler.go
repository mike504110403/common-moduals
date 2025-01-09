package errorhandler

import (
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/baseProtocol"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/loghttp"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/recoverhandler"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/tracehandler"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/ilog/v2"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.ErrorHandler {
	cfg := configDefault(config...)
	return func(c *fiber.Ctx, e error) error {
		if e != nil {
			_log := ilog.Trace(tracehandler.Get(c)).HTTP(loghttp.Get(c))
			if cfg.ResCloudTraceHeader {
				c.Set(tracehandler.XCloudTraceContext, tracehandler.Get(c).Raw)
			}
			switch e := e.(type) {
			case *baseProtocol.BaseResponse:
				if !e.RetStatus.IsSuccess() {
					if cfg.LogRetStatus {
						_log.Ret(e.RetStatus.StatusCode).Msg()
					}
					return c.JSON(e)
				}
				return c.JSON(baseProtocol.BaseResponse{})

			case *ilog.Data:
				if retStatus := e.BaseProtocol(); retStatus == nil {
					return c.Status(fiber.StatusInternalServerError).
						JSON(fiber.ErrInternalServerError)
				} else {
					if retStatus.IsSuccess() {
						return c.JSON(baseProtocol.BaseResponse{})
					}
					return c.JSON(baseProtocol.BaseResponse{RetStatus: *retStatus})
				}
			case *fiber.Error:
				if cfg.Log {
					_log.Err("HTTP_%d / %s", e.Code, e.Message)
				}
				return c.Status(e.Code).JSON(e)
			case *recoverhandler.Recover:
				_log.Call().Criti("%v\n%s", e.Recover, e.Stack)
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.ErrInternalServerError)
			default:
				if cfg.Log {
					_log.Err("HTTP_%d / %s", fiber.StatusInternalServerError, e)
				}
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.ErrInternalServerError)
			}
		}
		return c.Next()
	}
}
