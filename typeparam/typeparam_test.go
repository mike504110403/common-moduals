package typeparam

import (
	"testing"
)

// TestGetReversedMap: 反轉鍵值對測試
func TestGetReversedMap(t *testing.T) {
	testcase := SubTypeMap{
		"Tron":    1,
		"POLYGON": 2,
	}
	result := GetReversedMap(testcase)

	if result[1] != "Tron" {
		t.Errorf("TestGetReversedMap error: 預期反轉後 map[1] 為 'Tron', 卻得到 %s", result[1])
	}
}
