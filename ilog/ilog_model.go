package ilog

import (
	"log"
	"os"

	"github.com/mike504110403/common-moduals/baseProtocol"
	"github.com/mike504110403/common-moduals/fiber/handler/tracehandler"

	"github.com/valyala/fasthttp"
)

// LogData : log 基礎結構
type LogData struct {
	LogString  string
	MemberID   *string
	TradeSN    *string
	anyHeader  *[]anyHeader
	ctx        *fasthttp.RequestCtx
	cloudtrace *tracehandler.Trace
	callerInfo *callerInfo
	RetStatus  *baseProtocol.RetStatus
}

type anyHeader struct {
	Key, Value string
}

var (
	// WriteLogFile 寫入檔案
	serviceName = ""

	logInfo  = log.New(os.Stdout, "[Info] ", log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
	logError = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
)
