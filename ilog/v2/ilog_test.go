package ilog

import (
	"testing"

	"github.com/mike504110403/common-moduals/baseProtocol"
	"github.com/mike504110403/common-moduals/fiber/handler/tracehandler"
)

func TestLog(t *testing.T) {
	for _, l := range logMap {
		l.Printf("AAAA")
	}
}

func TestW(t *testing.T) {
	_trace, status := tracehandler.Conv("105445aa7843bc8bf206b12000100000/1;o=1")
	if !status {
		t.Error("trace status %t", status)
		return
	}
	Call().
		Trace(_trace).
		// Header(map[string]string{
		// 	"User-Agent":            "BestHTTP/2 v2.4.0",
		// 	"Host":                  "g1-token.gbaoonline.com",
		// 	"Content-Type":          "application/json",
		// 	"Content-Length":        "36",
		// 	"Accept-Encoding":       "gzip",
		// 	"Cf-Ipcountry":          "TW",
		// 	"Cf-Ray":                "6c108577cf1d6a98-TPE",
		// 	"X-Forwarded-Proto":     "https",
		// 	"Cf-Visitor":            "{\"scheme\":\"https\"}",
		// 	"Time":                  "1640083531452",
		// 	"Version":               "1.2.13",
		// 	"Authorization":         "Bearer eyJhbGciOiJIUzUxMiIsImtpZCI6IjEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJwdmlyOXptYyIsInN1YiI6IkdCNDM3NzI0MjQ2MiIsImV4cCI6MTY0MDA4NDMzNSwibmJmIjoxNjQwMDgzMDc1LCJpYXQiOjE2NDAwODMxMzUsImp0aSI6ImZkNzI2ZGVmLTAzMTItNDQ2Mi1iNzhkLTdmZjA2ODRjOWFjYiJ9.QBZ_ftwcZOobgXoPcIVvxkA2Jyn5-Yv5-j2HTyKMX9buWkFeIzvsUQbdGeIcHhhWEm7W16iN-tdJfn0xNJ9pig",
		// 	"Memberid":              "GB4377242462",
		// 	"Sign":                  "7/39dfe487c740e39d8f806c4904d1c6081f0cff8cee1ff50ddac964af84abeeac",
		// 	"Cf-Connecting-Ip":      "211.22.140.121",
		// 	"True-Client-Ip":        "211.22.140.121",
		// 	"Cdn-Loop":              "cloudflare",
		// 	"X-Cloud-Trace-Context": "0b4e8fb91ffd8367cc3ce32ac67baa70/3960617463050789398",
		// 	"Via":                   "1.1 google",
		// 	"X-Forwarded-For":       "211.22.140.121, 162.158.243.81, 35.201.79.194",
		// 	"Connection":            "Keep-Alive",
		// }).
		JSON([]byte(`{"machineId": "68682", "scatterNum": 1}`)).
		Warn("TEST")
}

func TestMsg(t *testing.T) {
	_log := Call().Ret(baseProtocol.TokenGuard_10101).Msg()
	_log.Default("AAAAAAAAAAAAAAA")
	_log.Debug("AAAAAAAAAAAAAAA")
	_log.Info("AAAAAAAAAAAAAAA")
	_log.Noti("AAAAAAAAAAAAAAA")
	_log.Warn("AAAAAAAAAAAAAAA")
	_log.Err("AAAAAAAAAAAAAAA")
	_log.Criti("AAAAAAAAAAAAAAA")
	_log.Alert("AAAAAAAAAAAAAAA")
	_log.Emergency("AAAAAAAAAAAAAAA")
}
