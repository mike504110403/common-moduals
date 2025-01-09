package tools

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var testString = "A"

func BenchmarkSHA256Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA256Hash(CreateRandomID_Any(5))
	}
}

func BenchmarkCRC32Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CRC32Hash(CreateRandomID_Any(5))
	}
}

func TestCRC32Hash(t *testing.T) {
	t.Log(CRC32Hash("JCG3809814421"))
}

func BenchmarkMD5hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MD5hash(CreateRandomID_Any(5))
	}
}

func TestNumOnly(t *testing.T) {
	str := "Hello X42 I'm a Y-32.35 string Z30"

	re := regexp.MustCompile("[0-9]+")

	submatchall := re.FindAllString(str, -1)
	submatchall2 := strings.Join(submatchall, "")
	fmt.Println(submatchall2)
}

func BenchmarkNumOnly_1(b *testing.B) {
	str := "Hello X42 I'm a Y-32.35 string Z30"
	for i := 0; i < b.N; i++ {
		i, err := NumOnly(str)
		if err != nil {
			b.Error(err)
		}
		fmt.Println(i)
	}
}

func Test_dotNetURLEncode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `-`, args: args{str: `-`}, want: `-`},
		{name: `_`, args: args{str: `_`}, want: `_`},
		{name: `.`, args: args{str: `.`}, want: `.`},
		{name: `!`, args: args{str: `!`}, want: `!`},
		{name: `~`, args: args{str: `~`}, want: `%7E`},
		{name: `*`, args: args{str: `*`}, want: `*`},
		{name: `(`, args: args{str: `(`}, want: `(`},
		{name: `)`, args: args{str: `)`}, want: `)`},
		{name: ` `, args: args{str: ` `}, want: `+`},
		{name: `@`, args: args{str: `@`}, want: `%40`},
		{name: `#`, args: args{str: `#`}, want: `%23`},
		{name: `$`, args: args{str: `$`}, want: `%24`},
		{name: `%`, args: args{str: `%`}, want: `%25`},
		{name: `^`, args: args{str: `^`}, want: `%5E`},
		{name: `&`, args: args{str: `&`}, want: `%26`},
		{name: `=`, args: args{str: `=`}, want: `%3D`},
		{name: `+`, args: args{str: `+`}, want: `%2B`},
		{name: `;`, args: args{str: `;`}, want: `%3B`},
		{name: `?`, args: args{str: `?`}, want: `%3F`},
		{name: `/`, args: args{str: `/`}, want: `%2F`},
		{name: `\`, args: args{str: `\`}, want: `%5C`},
		{name: `>`, args: args{str: `>`}, want: `%3E`},
		{name: `<`, args: args{str: `<`}, want: `%3C`},
		{name: "`", args: args{str: "`"}, want: `%60`},
		{name: `[`, args: args{str: `[`}, want: `%5B`},
		{name: `]`, args: args{str: `]`}, want: `%5D`},
		{name: `{`, args: args{str: `{`}, want: `%7B`},
		{name: `}`, args: args{str: `}`}, want: `%7D`},
		{name: `:`, args: args{str: `:`}, want: `%3A`},
		{name: `'`, args: args{str: `'`}, want: `%27`},
		{name: `"`, args: args{str: `"`}, want: `%22`},
		{name: `,`, args: args{str: `,`}, want: `%2C`},
		{name: `|`, args: args{str: `|`}, want: `%7C`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DotNetQueryEscape(tt.args.str); got != tt.want {
				t.Errorf("DotNetQueryEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_DotNetQueryEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DotNetQueryEscape(`-_.!~*() @#$%^&=+;?/\><` + "`" + `[]{}:'",|`)
	}
}

func TestOnlyQueryEscapeToLower(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `-`, args: args{str: `-`}, want: `-`},
		{name: `_`, args: args{str: `_`}, want: `_`},
		{name: `.`, args: args{str: `.`}, want: `.`},
		{name: `!`, args: args{str: `!`}, want: `!`},
		{name: `~`, args: args{str: `~`}, want: `%7e`},
		{name: `*`, args: args{str: `*`}, want: `*`},
		{name: `(`, args: args{str: `(`}, want: `(`},
		{name: `)`, args: args{str: `)`}, want: `)`},
		{name: ` `, args: args{str: ` `}, want: `+`},
		{name: `@`, args: args{str: `@`}, want: `%40`},
		{name: `#`, args: args{str: `#`}, want: `%23`},
		{name: `$`, args: args{str: `$`}, want: `%24`},
		{name: `%`, args: args{str: `%`}, want: `%25`},
		{name: `^`, args: args{str: `^`}, want: `%5e`},
		{name: `&`, args: args{str: `&`}, want: `%26`},
		{name: `=`, args: args{str: `=`}, want: `%3d`},
		{name: `+`, args: args{str: `+`}, want: `%2b`},
		{name: `;`, args: args{str: `;`}, want: `%3b`},
		{name: `?`, args: args{str: `?`}, want: `%3f`},
		{name: `/`, args: args{str: `/`}, want: `%2f`},
		{name: `\`, args: args{str: `\`}, want: `%5c`},
		{name: `>`, args: args{str: `>`}, want: `%3e`},
		{name: `<`, args: args{str: `<`}, want: `%3c`},
		{name: "`", args: args{str: "`"}, want: `%60`},
		{name: `[`, args: args{str: `[`}, want: `%5b`},
		{name: `]`, args: args{str: `]`}, want: `%5d`},
		{name: `{`, args: args{str: `{`}, want: `%7b`},
		{name: `}`, args: args{str: `}`}, want: `%7d`},
		{name: `:`, args: args{str: `:`}, want: `%3a`},
		{name: `'`, args: args{str: `'`}, want: `%27`},
		{name: `"`, args: args{str: `"`}, want: `%22`},
		{name: `,`, args: args{str: `,`}, want: `%2c`},
		{name: `|`, args: args{str: `|`}, want: `%7c`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnlyQueryEscapeToLower(DotNetQueryEscape(tt.args.str)); got != tt.want {
				t.Errorf("OnlyQueryEscapeToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAA(t *testing.T) {

	a := &struct {
		C string
		D string
	}{
		C: "ZZZ",
		D: "FFFF",
	}

	b, _ := json.Marshal(a)

	t.Logf("TEST: %s", b)
	fmt.Printf("TEST: %s", b)

}

func TestIsJSONString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str: `[
					{
						"mWeight": "225000",
						"mCategory": "1885",
						"mQuantity": "1",
						"mAttachment": "1"
					},
					{
						"mWeight": "60000",
						"mCategory": "1886",
						"mQuantity": "1",
						"mAttachment": "1"
					}
				]`,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `{"mWeight": "225000", "mCategory": "1885", "mQuantity": "1", "mAttachment": "1"}`,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": "225000"`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: `00000`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "[{\"mWeight\": \"225000\", \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}, {\"mWeight\": \"60000\", \"mCategory\": \"1886\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}]",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "{\"mWeight\": \"225000\", \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": "999999999999999"`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": 999999999999999`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "{\"mWeight\": 999999999999999999, \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "{99999999999999999999999}",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "999999999999999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSONString(tt.args.str); got != tt.want {
				t.Errorf("IsJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSONStringV2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str: `[
					{
						"mWeight": "225000",
						"mCategory": "1885",
						"mQuantity": "1",
						"mAttachment": "1"
					},
					{
						"mWeight": "60000",
						"mCategory": "1886",
						"mQuantity": "1",
						"mAttachment": "1"
					}
				]`,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `{"mWeight": "225000", "mCategory": "1885", "mQuantity": "1", "mAttachment": "1"}`,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": "225000"`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: `00000`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "[{\"mWeight\": \"225000\", \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}, {\"mWeight\": \"60000\", \"mCategory\": \"1886\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}]",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "{\"mWeight\": \"225000\", \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": "999999999999999"`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: `"mWeight": 999999999999999`,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "{\"mWeight\": 999999999999999999, \"mCategory\": \"1885\", \"mQuantity\": \"1\", \"mAttachment\": \"1\"}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "{99999999999999999999999}",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				str: "999999999999999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSONStringV2(tt.args.str); got != tt.want {
				t.Errorf("IsJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestIsTimeout_FastHTTPCli(t *testing.T) {
// 	c := &fasthttp.Client{
// 		MaxConnWaitTimeout: 5 * time.Second,
// 		ReadTimeout:        5 * time.Second,
// 		WriteTimeout:       5 * time.Second,
// 	}
// 	tt := map[string]string{}
// 	err := fasthttpClient.JSON(c, "http://localhost:8080/version", fasthttpClient.GET, nil, &tt)
// 	t.Log(err)
// 	t.Logf("IsTemporary : %t", IsTimeout(err))
// }

// func TestIsTimeout_HTTPCli(t *testing.T) {
// 	c := &http.Client{
// 		Timeout: 5 * time.Second,
// 	}
// 	tt := map[string]string{}
// 	err := httpClientService.JSON(c, "http://localhost:8080/version", httpClientService.GET, nil, &tt)
// 	t.Log(err)
// 	t.Logf("IsTimeout : %t", IsTimeout(err))
// }
