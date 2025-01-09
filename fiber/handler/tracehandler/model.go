package tracehandler

// Trace : https://cloud.google.com/trace/docs/setup
type Trace struct {
	TraceID   TraceID   `json:"TraceID"`
	SpanID    SpanID    `json:"SpanID"`
	TraceTrue TraceTrue `json:"TraceTrue"`
	Raw       string    `json:"raw"`
}

type (
	// TraceID : is a 32-character hexadecimal value representing a 128-bit number. It should be unique between your requests, unless you intentionally want to bundle the requests together. You can use UUIDs.
	TraceID = string

	// SpanID : is the decimal representation of the (unsigned) span ID. It should be randomly generated and unique in your trace. For subsequent requests, set SPAN_ID to the span ID of the parent request. See the description of TraceSpan (REST, RPC) for more information about nested traces.
	SpanID = string

	// TraceTrue : must be 1 to trace this request. Specify 0 to not trace the request.
	//
	// Cloud Trace doesn't sample every request. For example, if you use Java and OpenCensus, then only 1 request out of every 10,000 is traced. If you are using App Engine, requests are sampled at a rate of 0.1 requests per second for each App Engine instance. If you use the Cloud Trace API, then you can configure customer rates. Some packages, such as the Java OpenCensus package, support configuring the sampling rate.
	TraceTrue = bool
)

const (
	// KeyTrace : 寫入ctx locals 的 key
	KeyTrace = "Trace"
	// KeySpanID : 寫入ctx locals 的 key
	KeySpanID = "SpanID"
	// KeyTraceTrue : 寫入ctx local s的 key
	KeyTraceTrue = "TraceTrue"
)

// XCloudTraceContext : https://cloud.google.com/trace/docs/setup
const XCloudTraceContext = "X-Cloud-Trace-Context"

// ClientSet : 自訂轉發參數
type ClientSet struct {
	SpanID    SpanID
	TraceTrue TraceTrue
}
