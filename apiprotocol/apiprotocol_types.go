package apiprotocol

import "cloud.google.com/go/logging"

// BaseResponse : 基礎回應資料結構
type BaseResponse struct {
	RetStatus RetStatus `json:"retStatus"`
}

// RetStatus : 回傳狀態
type RetStatus struct {
	Code  Code             `json:"code"`
	Msg   string           `json:"msg,omitempty"`
	Error string           `json:"error,omitempty"`
	Level logging.Severity `json:"level,omitempty"`
}

// RetStatusContent : 回傳狀態內容
type RetStatusContent struct {
	Msg   string           `json:"msg"`
	Level logging.Severity `json:"level"`
}

// APIResponse : API統一回傳結構
type APIResponse struct {
	Data      any       `json:"data"`
	RetStatus RetStatus `json:"retStatus"`
}

// StatusResponse : 狀態回傳結構
type StatusResponse struct {
	RetStatus RetStatus `json:"retStatus"`
}

type Code int

// Success10000 : 正常回應
const Success10000 Code = 10000

var retStatusList = map[Code]RetStatusContent{
	Success10000: {Msg: "Success"},
}
