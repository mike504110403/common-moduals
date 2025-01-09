package ilog

import (
	"fmt"
	"strconv"
	"time"

	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/baseProtocol"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/fiber/handler/tracehandler"
	"ec2-15-168-3-237.ap-northeast-3.compute.amazonaws.com/gogogo/common-moduals/tools"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/utils"
	"google.golang.org/api/logging/v2"
)

// Init : 初始化
func Init(config ...Config) error {
	cfg := configDefault(config...)
	reportSet = cfg.ReportedErrorEvent
	reportLevel = parseLevel(cfg.ReportLevel)
	return nil
}

func (l *logger) cloudWrite(e Entry) error {
	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, isGet := reportLevel[e.Severity]; isGet {
		e.GcpType = reportSet
	}
	e.TimeStamp = now.Format(time.RFC3339Nano)
	b, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	l.buf = l.buf[:0]
	l.buf = append(l.buf, b...)
	l.buf = append(l.buf, '\n')
	_, err = l.out.Write(l.buf)
	return err
}

func localWrite(e Entry) error {
	b := []byte{}
	if e.SourceLocation != nil {
		b = append(b, fmt.Sprintf("[call->%s:%d(%s)]", e.SourceLocation.File, e.SourceLocation.Line, e.SourceLocation.Function)...)
	}

	if e.TraceEntry != nil {
		b = append(b, fmt.Sprintf("[traceID->%s]", e.TraceEntry.Trace)...)
	}

	if e.HttpRequest != nil {
		b = append(b, fmt.Sprintf("%#v", e.HttpRequest)...)
	}

	if e.Header != nil {
		b = append(b, fmt.Sprintf("%#v", e.Header)...)
	}

	if e.Labels != nil {
		for k, v := range e.Labels {
			b = append(b, fmt.Sprintf("[%s->%s]", k, v)...)
		}
	}

	if e.JsonPayload != nil {
		b = append(b, fmt.Sprintf("%s", e.JsonPayload)...)
	}

	if len(b) > 0 {
		b = append(b, " "...)
	}
	b = append(b, e.Message...)

	l, isGet := logMap[e.Severity]
	if !isGet {
		l = logMap[EMERGENCY]
	}
	l.Print(utils.GetString(b))
	return nil
}

func (d *Data) write(se severity, format string, v ...interface{}) *Data {
	defer func() { d.Message = ""; d.JsonPayload = nil; d.SourceLocation = nil }()
	d.Message = fmt.Sprintf(format, v...)
	d.Severity = se

	if err := write(d.Entry); err != nil {
		writeErr.Printf("[%s] %#v", err, d)
	}
	return d
}

// Writef : 寫入log(包含format)
func (d *Data) Writef(se severity, format string, v ...interface{}) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.write(se, format, v...)
}

// Writef : 寫入log(包含format)
func Writef(se severity, format string, v ...interface{}) *Data {
	return (&Data{}).Writef(se, format, v...)
}

// Writef : 寫入log(包含format)
func (d *Data) Write(se severity, s string) *Data { return d.Writef(se, "%s", s) }

func Write(se severity, s string) *Data { return (&Data{}).Writef(se, "%s", s) }

// Call : 新增程式行數、位置、func name 等資料
func (d *Data) Call(skip ...int) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	_skip := 2
	if len(skip) > 0 {
		_skip = skip[0]
	}
	_file, _func, _line := tools.CallerInfo2(_skip)
	d.Entry.SourceLocation = &logging.LogEntrySourceLocation{
		File:     _file,
		Function: _func,
		Line:     int64(_line),
	}
	return d
}

// Call : 新增程式行數、位置、func name 等資料
func Call(skip ...int) *Data { return (&Data{}).Call(skip...) }

// Label : 為 log 加上標籤
func (d *Data) label(key, value string) *Data {
	if d.Entry.Labels == nil {
		d.Entry.Labels = make(map[string]string)
	}
	d.Entry.Labels[key] = value
	return d
}

// Label : 為 log 加上標籤
func (d *Data) Label(key, value string) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.label(key, value)
}

// Label : 為 log 加上標籤
func Label(key, value string) *Data { return (&Data{}).Label(key, value) }

// Member : 於 log 的 header 新增 MemberID
func (d *Data) Member(MemberID string) *Data { return d.Label(LabelMemberID, MemberID) }

// Member : 於 log 的 header 新增 MemberID
func Member(MemberID string) *Data { return (&Data{}).Member(MemberID) }

// Trade : 於 log 的 header 新增 TradeSN
func (d *Data) Trade(TradeSN string) *Data { return d.Label(LabelTradeSN, TradeSN) }

// Trade : 於 log 的 header 新增 TradeSN
func Trade(TradeSN string) *Data { return (&Data{}).Trade(TradeSN) }

// Ret : 於 log 的 header 新增 retStatusCode 並更新 log 內部結構資料, 可用於接續Msg
func (d *Data) Ret(retStatusCode int) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.retStatus = baseProtocol.Create(retStatusCode)
	d.label("retStatus", strconv.Itoa(retStatusCode))
	return d
}

// Ret : 於 log 的 header 新增 retStatusCode 並更新 log 內部結構資料, 可用於接續Msg
func Ret(retStatus int) *Data { return (&Data{}).Ret(retStatus) }

// BaseProtocol : 輸出結果給 BaseProtocol
func (d *Data) BaseProtocol() *baseProtocol.RetStatus {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.retStatus
}

// RetErr : 輸出 RetStatus 結果給 error handler
func (d *Data) RetErr() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.retStatus.Err()
}

// Trace : 新增 cloud trace id
func (d *Data) Trace(t tracehandler.Trace) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.TraceEntry = &TraceEntry{
		Trace:   string(t.TraceID),
		SpanID:  string(t.SpanID),
		Sampled: t.TraceTrue,
	}
	return d
}

// Trace : 新增 cloud trace id
func Trace(t tracehandler.Trace) *Data { return (&Data{}).Trace(t) }

// Body : 紀錄HTTP Body(自動判斷是否為json)
func (d *Data) Body(b []byte) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	if json.ConfigStd.Valid(b) {
		d.Entry.JsonPayload = b
	} else {
		d.Entry.Message = utils.GetString(b)
	}
	return d
}

// Body : 紀錄HTTP Body(自動判斷是否為json)
func Body(b []byte) *Data { return (&Data{}).Body(b) }

// JSON : 紀錄JSON Byte
func (d *Data) JSON(b []byte) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entry.JsonPayload = b
	return d
}

// JSON : 紀錄JSON Byte
func JSON(b []byte) *Data { return (&Data{}).JSON(b) }

// Struct : 紀錄結構並轉為JSON儲存(會被golang的JSON機制限制)
func (d *Data) Struct(i interface{}) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	b, err := json.Marshal(i)
	if err != nil {
		d.Entry.Message = fmt.Sprintf("[json.Marshal: %s] %#v", err, i)
	} else {
		d.Entry.JsonPayload = b
	}
	return d
}

// Struct : 紀錄結構並轉為JSON儲存(會被golang的JSON機制限制)
func Struct(i interface{}) *Data { return (&Data{}).Struct(i) }

// Header : 紀錄HTTP Header
func (d *Data) Header(m map[string][]string) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entry.Header = m
	return d
}

// Header : 紀錄HTTP Header
func Header(m map[string][]string) *Data { return (&Data{}).Header(m) }

// HTTP : 紀錄HTTP相關資料
//
//	一般會放置於 defer 並配合 logreq 套件
func (d *Data) HTTP(h *logging.HttpRequest) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.HttpRequest = h
	return d
}

// HTTP : 紀錄HTTP相關資料
//
//	一般會放置於 defer 並配合 logreq 套件
func HTTP(h *logging.HttpRequest) *Data { return (&Data{}).HTTP(h) }

// Msg : 將 RetStatus 結構內的訊息轉為 Log 並紀錄 (不填寫內容自動判斷是否為 ERROR)
func (d *Data) Msg(se ...severity) *Data {
	d.mu.Lock()
	defer d.mu.Unlock()
	switch {
	case len(se) > 0:
		return d.write(se[0], "[%d]%s", d.retStatus.StatusCode, d.retStatus.StatusMsg)

	case d != nil && d.retStatus != nil:
		if d.retStatus.IsSuccess() {
			return d.write(INFO, "[%d]%s", d.retStatus.StatusCode, d.retStatus.StatusMsg)
		}
		return d.write(ERROR, "[%d]%s", d.retStatus.StatusCode, d.retStatus.StatusMsg)

	default:
		return d

	}
}

func (d *Data) Error() string {
	return d.RetErr().Error()
}
