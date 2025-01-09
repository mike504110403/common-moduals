package tracehandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// New : 新增中間層用於處理 Cloud Trace
//
// config 暫時無值
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)
	traceKey = cfg.TraceHaderKey
	return func(c *fiber.Ctx) error {
		_trace := auto(c)
		c.Locals(KeyTrace, _trace)
		if cfg.AllResAddTraceID {
			defer c.Set(traceKey, _trace.Raw)
		}
		return c.Next()
	}
}

// auto : 自動處理，如果header沒有就自動產生
func auto(c *fiber.Ctx) Trace {
	if _trace, status := Conv(c.Get(traceKey)); status {
		return _trace
	}
	return NewTrace(c)
}

func NewTrace(c *fiber.Ctx) Trace {
	_trace := Trace{
		TraceID:   utils.UUID(),
		SpanID:    strconv.FormatUint(c.Context().ID(), 10),
		TraceTrue: false,
	}
	_trace.Raw = _trace.Conv()
	c.Request().Header.Add(traceKey, _trace.Raw)
	return _trace
}

// Get : 取得傳入 Trace 相關 code
func Get(c *fiber.Ctx) Trace {
	if trace, isTrace := c.Locals(KeyTrace).(Trace); isTrace {
		return trace
	}
	return NewTrace(c)
}

func GetTraceKey() string {
	return traceKey
}

// Conv : 從 header 上的 X-Cloud-Trace-Context 取值
//
// curl "http://www.example.com" --header "X-Cloud-Trace-Context: 105445aa7843bc8bf206b12000100000/1;o=1"
func Conv(rawStr string) (t Trace, status bool) {
	t.Raw = rawStr
	if st := strings.Split(rawStr, "/"); len(st) > 1 {
		status = true
		t.TraceID = st[0]
		t.SpanID = st[1]
		if span := strings.Split(st[1], ";o="); len(span) > 1 {
			t.SpanID = span[0]
			t.TraceTrue = span[1] == "1"
		}
	}
	return
}

// Conv : 將 GoogleCloudTrace 轉為string結構
//
// "X-Cloud-Trace-Context: TRACE_ID/SPAN_ID;o=TRACE_TRUE"
func (t Trace) Conv() (str string) {
	return t.TraceID + "/" + t.SpanID + traceTrue(t.TraceTrue)
}

func traceTrue(t bool) string {
	if t {
		return ";o=1"
	}
	return ";o=0"
}

// HTTPHeader : 設定 http 發出的 header
func (t Trace) HTTPHeader(req *http.Request) {
	req.Header.Set(traceKey, t.Raw)
}

// FastHTTPHeader : 設定 fasthttp 發出的 header
func (t Trace) FastHTTPHeader(req *fasthttp.Request) {
	req.Header.Set(traceKey, t.Raw)
}

// Res : 回應追蹤碼
func Res(c *fiber.Ctx) error {
	defer c.Set(traceKey, Get(c).Raw)
	return c.Next()
}
