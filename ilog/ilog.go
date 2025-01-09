package ilog

import (
	"fmt"

	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/baseProtocol"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/tracehandler"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/tools"

	"github.com/valyala/fasthttp"
)

// New : 初始化
func New(_strviceName string) {
	serviceName = _strviceName
}

// Log : 一般紀錄
func Log(format string, v ...interface{}) {
	logInfo.Printf(format, v...)
}

// ErrorLog : 錯誤紀錄
func ErrorLog(format string, v ...interface{}) {
	logError.Printf(format, v...)
}

// Fatal : log 並跳出
func Fatal(format string, v ...interface{}) {
	logError.Fatalf(format, v...)
}

type callerInfo struct {
	file     string
	funcName string
	line     int
}

// Basic : 起始開頭什麼都不做
func Basic() *LogData {
	return &LogData{}
}

// Call : 於 log 的 header 新增 func資訊 並更新 log 內部結構資料
func Call(skip int) *LogData {
	file, funcName, line := tools.CallerInfo2(skip)
	return &LogData{
		callerInfo: &callerInfo{
			file:     file,
			funcName: funcName,
			line:     line,
		},
	}
}

// Call : 於 log 的 header 新增 func資訊 並更新 log 內部結構資料
func (logdata *LogData) Call(skip int) *LogData {
	logdata.callerInfo = Call(skip + 1).callerInfo
	return logdata
}

func (logdata *LogData) Add(key, value string) *LogData {
	if logdata.anyHeader == nil {
		logdata.anyHeader = &[]anyHeader{
			{
				Key:   key,
				Value: value,
			},
		}
		return logdata
	}
	*logdata.anyHeader = append(*logdata.anyHeader, anyHeader{
		Key:   key,
		Value: value,
	})
	return logdata
}

// Ctx : 於 log 的 header 新增 連線資訊 並更新 log 內部結構資料
func (logdata *LogData) Ctx(ctx *fasthttp.RequestCtx) *LogData {
	logdata.ctx = ctx
	return logdata
}

// BaseHeader : 基礎 header , 於 header 新增 func 及 連線資訊
func BaseHeader(ctx *fasthttp.RequestCtx, skipCaller int) *LogData {
	return &LogData{
		callerInfo: Call(skipCaller + 1).callerInfo,
		ctx:        ctx,
	}
}

// Update : 將 retStatusCode 更新至 回覆用結構 及 log 內部結構資料, 可用於接續Msg
func (logdata *LogData) Update(retStatus *baseProtocol.RetStatus) *LogData {
	retStatus.Update(logdata.RetStatus.StatusCode)
	return logdata
}

// Ret : 於 log 的 header 新增 retStatusCode 並更新 log 內部結構資料, 可用於接續Msg
func (logdata *LogData) Ret(retStatusCode int) *LogData {
	logdata.RetStatus = baseProtocol.Create(retStatusCode)
	return logdata
}

// Member : 於 log 的 header 新增 MemberID
func (logdata *LogData) Member(MemberID string) *LogData {
	logdata.MemberID = &MemberID
	return logdata
}

// Trade : 於 log 的 header 新增 TradeSN
func (logdata *LogData) Trade(TradeSN string) *LogData {
	logdata.TradeSN = &TradeSN
	return logdata
}

func (logdata *LogData) Trace(t tracehandler.Trace) *LogData {
	logdata.cloudtrace = &t
	return logdata
}

// Msg : 將 RetStatus 結構內的訊息轉為 Log 並紀錄 (自動判斷是否為ErrorLog)
func (logdata *LogData) Msg() *LogData {
	logdata.genHeader()
	if logdata != nil && logdata.RetStatus != nil {
		logdata.LogString = logdata.LogString + "" + logdata.RetStatus.StatusMsg
		if logdata.RetStatus.StatusCode != baseProtocol.Success {
			ErrorLog(logdata.LogString)
		} else {
			Log(logdata.LogString)
		}
	}
	return logdata
}

// LogMsg : 將 RetStatus 結構內的訊息轉為 Log 並紀錄 (強制為Log紀錄)
func (logdata *LogData) LogMsg() *LogData {
	logdata.genHeader()
	if logdata != nil && logdata.RetStatus != nil {
		logdata.LogString = logdata.LogString + " " + logdata.RetStatus.StatusMsg
		Log(logdata.LogString)
	}
	return logdata
}

// genHeader : 清空並產生log header
func (logdata *LogData) genHeader() *LogData {
	if logdata == nil {
		logdata = Basic()
	}
	logdata.LogString = ""
	if logdata.cloudtrace != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf(
			`[traceID->%s]`, logdata.cloudtrace.Raw,
		)
	}
	if logdata.ctx != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf(
			`[reqID->%d][connID->%d][reqNum->%d]`,
			logdata.ctx.ID(),
			logdata.ctx.ConnID(),
			logdata.ctx.ConnRequestNum(),
		)
	}
	if logdata.callerInfo != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf(
			`[%s:%d(%s)]`,
			logdata.callerInfo.file,
			logdata.callerInfo.line,
			logdata.callerInfo.funcName,
		)
	}
	if logdata.MemberID != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf("[MemberID->%s]", *logdata.MemberID)
	}
	if logdata.TradeSN != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf("[TradeSN->%s]", *logdata.TradeSN)
	}
	if logdata.RetStatus != nil {
		logdata.LogString = logdata.LogString + fmt.Sprintf(
			`[%s_%d]`,
			serviceName, logdata.RetStatus.StatusCode,
		)
	}
	if logdata.anyHeader != nil {
		for _, v := range *logdata.anyHeader {
			logdata.LogString = logdata.LogString + fmt.Sprintf(`[%s->%s]`, v.Key, v.Value)
		}
	}
	return logdata
}

// Log : 紀錄為一般Log
func (logdata *LogData) Log(format string, v ...interface{}) *LogData {
	logdata.genHeader()
	Log(logdata.LogString+" "+format, v...)
	logdata.LogString = logdata.LogString + " " + fmt.Sprintf(format, v...)
	return logdata
}

// Err : 紀錄為錯誤Log
func (logdata *LogData) Err(format string, v ...interface{}) *LogData {
	logdata.genHeader()
	ErrorLog(logdata.LogString+" "+format, v...)
	logdata.LogString = logdata.LogString + " " + fmt.Sprintf(format, v...)
	return logdata
}

// Fatal : 紀錄為錯誤Log並跳出
func (logdata *LogData) Fatal(format string, v ...interface{}) *LogData {
	logdata.genHeader()
	Fatal(logdata.LogString+" "+format, v...)
	logdata.LogString = logdata.LogString + " " + fmt.Sprintf(format, v...)
	return logdata
}

// ErrStr : 紀錄為錯誤Log
func (logdata *LogData) ErrStr(err error) *LogData {
	logdata.genHeader()
	ErrorLog(logdata.LogString+" %s", err)
	logdata.LogString = logdata.LogString + " " + fmt.Sprintf("%s", err)
	return logdata
}

// Res : 改變http status code
func (logdata *LogData) Res(code int) *LogData {
	logdata.ctx.SetStatusCode(code)
	logdata.ctx.SetBodyString(fasthttp.StatusMessage(code))
	return logdata
}

// BaseProtocol : 輸出結果給BaseProtocol
func (logdata *LogData) BaseProtocol() *baseProtocol.RetStatus {
	return logdata.RetStatus
}
