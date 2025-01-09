package ilog

import (
	"log"
	"os"

	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/baseProtocol"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/tracehandler"

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
