package apiprotocol

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"cloud.google.com/go/logging"
	"github.com/gofiber/fiber/v2"
	mlog "github.com/mike504110403/goutils/log"
	"github.com/valyala/fasthttp"
)

type Config struct {
	ServiceNo int
}

type ErrorHandlerConfig struct {
	// ResponseTraceHeader : 回應 cloud trace header
	ResCloudTraceHeader bool
	ResError            bool
	Log                 bool
}

var cfg = Config{}
var errHandlerCfg = ErrorHandlerConfig{}

func Init(initCfg Config) {
	cfg = initCfg
}
func Append(initRetStatusContent map[Code]RetStatusContent) {
	for k, v := range initRetStatusContent {
		retStatusList[k] = v
	}
}

// ToRes : 直接組成可以回傳的資料結構(通常直接JSON送出)
func (code Code) ToRes() *BaseResponse {
	msg := "未知的錯誤"
	level := logging.Error
	code = PadServiceNo(Code(cfg.ServiceNo), code)
	if retStatus, exist := retStatusList[code]; exist {
		msg = retStatus.Msg
		level = retStatus.Level
	}
	return &BaseResponse{
		RetStatus: RetStatus{
			Code:  code,
			Msg:   msg,
			Level: level,
		},
	}
}

// ToRet : 組裝出基礎結構待之後可以進行進一步處理
func (code Code) ToRet() *RetStatus {
	code = PadServiceNo(Code(cfg.ServiceNo), code)
	msg := "未知的錯誤"
	level := logging.Error
	if retStatus, exist := retStatusList[code]; exist {
		msg = retStatus.Msg
		level = retStatus.Level
	}
	return &RetStatus{
		Code:  code,
		Msg:   msg,
		Level: level,
	}

}

func (br *BaseResponse) Err(err string) *BaseResponse {
	br.RetStatus.Error = err
	return br
}

func (br *BaseResponse) ToErr() error {
	if br.RetStatus.Code == Success10000 {
		return nil
	}
	return br
}

func (br *BaseResponse) Msg() string {
	return br.RetStatus.Msg
}
func (br *BaseResponse) Error() string {
	return br.RetStatus.Error
}

func (br *BaseResponse) LogError() string {
	return fmt.Sprintf("Msg: %s, Error: %s", br.Msg(), br.RetStatus.Error)
}

// PadServiceNo : 增加服務編號
func PadServiceNo(serviceNo, code Code) Code {
	serviceNoString := strconv.Itoa(int(serviceNo))
	codeString := strconv.Itoa(int(code))
	if result, err := strconv.Atoi(serviceNoString + codeString); err != nil {
		return code
	} else {
		return Code(result)
	}
}

func ErrorHandler(config ...ErrorHandlerConfig) fiber.ErrorHandler {
	return func(c *fiber.Ctx, e error) error {
		if e != nil {
			switch err := e.(type) {
			case *BaseResponse:
				code := err.RetStatus.Code
				msg := fmt.Sprintf("[code]->%d ", code)
				if err.RetStatus.Error != "" {
					msg += fmt.Sprintf("[error]->%s ", err.RetStatus.Error)
				}
				switch err.RetStatus.Level {
				case logging.Debug:
					mlog.Debug(msg)
				case logging.Error:
					mlog.Error(msg)
				case logging.Alert:
					mlog.Fatal(msg)
				case logging.Warning:
					mlog.Warning(msg)
				default:
				}
				// 輸出一個只有code 的 RetStatus
				return c.JSON(BaseResponse{RetStatus: RetStatus{Code: code}})
			case *fiber.Error:
				code := err.Code
				if code != fasthttp.StatusNotFound {
					msg := fmt.Sprintf("[code]->%d ", code)
					if err.Message != "" {
						msg += fmt.Sprintf("[error]->%s ", err.Message)
						mlog.Error(msg)
					}
				}
				return c.SendStatus(err.Code)
			default:
				if err.Error() != "" {
					msg := fmt.Sprintf("[error]->%s ", err.Error())
					mlog.Error(msg)
				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
			}
		}
		return c.Next()
	}
}

func CallInfo(callstack int) (file string, funcName string, line int) {
	pc, fileFullPath, l, _ := runtime.Caller(callstack)
	funcFullPath := runtime.FuncForPC(pc).Name()
	functionNameSplit := strings.Split(funcFullPath, "/")
	fileFullPathSplit := strings.Split(fileFullPath, "/")
	if len(fileFullPathSplit) == 1 {
		fileFullPathSplit = strings.Split(fileFullPath, "\\")
	}
	fileName := fileFullPathSplit[len(fileFullPathSplit)-1]
	if len(functionNameSplit) > 2 {
		file = strings.Join(functionNameSplit[1:len(functionNameSplit)-1], "/") + "/" + fileName
	} else {
		file = strings.Join(functionNameSplit, "/") + "/" + fileName
	}

	funcNameSplit := strings.Split(functionNameSplit[len(functionNameSplit)-1], ".")
	funcName = funcNameSplit[len(funcNameSplit)-1]
	return file, funcName, l
}
