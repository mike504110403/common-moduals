package ilog

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/mike504110403/common-moduals/baseProtocol"

	"github.com/fatih/color"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/logging/v2"
)

// Entry : log 基礎結構
type (
	logger struct {
		buf []byte
		out io.Writer
		mu  sync.Mutex
	}

	Data struct {
		Entry
		mu sync.Mutex
	}

	Entry struct {
		TimeStamp      string                          `json:"time,omitempty"`
		Severity       severity                        `json:"severity,omitempty"`
		Header         map[string][]string             `json:"header,omitempty"`
		Labels         map[string]string               `json:"logging.googleapis.com/labels,omitempty"`
		SourceLocation *logging.LogEntrySourceLocation `json:"logging.googleapis.com/sourceLocation,omitempty"`
		HttpRequest    *logging.HttpRequest            `json:"httpRequest,omitempty"`
		JsonPayload    googleapi.RawMessage            `json:"json,omitempty"`
		Message        string                          `json:"message,omitempty"`
		GcpType        string                          `json:"@type,omitempty"`
		retStatus      *baseProtocol.RetStatus
		*TraceEntry
	}

	TraceEntry struct {
		Trace   string `json:"logging.googleapis.com/trace,omitempty"`
		SpanID  string `json:"logging.googleapis.com/spanId,omitempty"`
		Sampled bool   `json:"logging.googleapis.com/trace_sampled,omitempty"`
	}

	severity string
	// Config :GCP Error Report設定
	Config struct {
		ReportedErrorEvent string
		ReportLevel        []severity
	}
)

// Config :GCP Error Report設定
var (
	ConfigDefault = Config{
		ReportedErrorEvent: "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent",
		ReportLevel:        []severity{ERROR, CRITICAL, ALERT, EMERGENCY},
	}
)

var (
	sout        = logger{out: os.Stdout}
	writeErr    = log.New(os.Stderr, "[writeErr]", log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
	write       = localWrite
	reportSet   = ConfigDefault.ReportedErrorEvent
	reportLevel = parseLevel(ConfigDefault.ReportLevel)

	colorMap = map[severity]*color.Color{
		DEFAULT:   color.New(color.FgGreen),
		DEBUG:     color.New(color.FgCyan),
		INFO:      color.New(color.FgWhite),
		NOTICE:    color.New(color.FgBlue),
		WARNING:   color.New(color.FgHiYellow),
		ERROR:     color.New(color.FgHiRed),
		CRITICAL:  color.New(color.BgHiRed, color.FgBlack),
		ALERT:     color.New(color.BgHiRed, color.FgHiWhite),
		EMERGENCY: color.New(color.BgHiRed, color.FgHiYellow),
	}

	logMap = map[severity]*log.Logger{}
)

func isRunOnGoTestViewMode() bool {
	return isArgGet("-test.v=true")
}

func isRunOnVSCode() bool {
	return strings.Contains(os.Args[0], "debug_bin") || isArgGet("-test.run") || isArgGet("-test.v=true")
}

func init() {
	if isRunOnVSCode() {
		for l, c := range colorMap {
			var flag int
			if isRunOnGoTestViewMode() {
				flag = log.Ltime | log.Lmsgprefix
			} else {
				c.EnableColor()
				flag = log.Ltime | log.Lmsgprefix
			}
			prefix := c.Sprintf("[%9s]", l) + " "
			switch l {
			case EMERGENCY, ALERT, CRITICAL, ERROR, WARNING:
				logMap[l] = log.New(os.Stderr, prefix, flag)
			default:
				logMap[l] = log.New(os.Stdout, prefix, flag)
			}
		}
	} else {
		write = sout.cloudWrite
	}
}

const (
	DEFAULT   severity = "DEFAULT"
	DEBUG     severity = "DEBUG"
	INFO      severity = "INFO"
	NOTICE    severity = "NOTICE"
	WARNING   severity = "WARNING"
	ERROR     severity = "ERROR"
	CRITICAL  severity = "CRITICAL"
	ALERT     severity = "ALERT"
	EMERGENCY severity = "EMERGENCY"
)

const (
	LabelMemberID = "MemberID"
	LabelTradeSN  = "TradeSN"
)
