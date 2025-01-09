package baseProtocol

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRetStatus_MarshalJSON(t *testing.T) {
	type XX struct {
		S         string
		RetStatus RetStatus `json:"retStatus"`
	}

	xx := &XX{
		S: "asdf",
	}

	b, err := json.Marshal(xx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", b)
}

func TestRetStatusFn(t *testing.T) {
	if err := retStatusFn(); err != nil {
		t.Logf("%v", err)
	}
}

func TestBaseResponseFn(t *testing.T) {
	if err := BaseResponseFn(); err != nil {
		t.Logf("%v", err)
	}
}

func retStatusFn() error {
	return RetStatus{
		StatusCode: 99111111,
		StatusMsg:  "testcode",
		SystemTime: time.Now().UnixMilli(),
		CheckCode:  0,
	}
}

func BaseResponseFn() error {
	return BaseResponse{RetStatus{
		StatusCode: 99111111,
		StatusMsg:  "testcode",
		SystemTime: time.Now().UnixMilli(),
		CheckCode:  0,
	}}
}
