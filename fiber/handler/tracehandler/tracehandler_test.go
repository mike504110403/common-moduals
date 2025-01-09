package tracehandler

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	//"git-golang.yile808.com/common/mod-common/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

var app *fiber.App

func TestMain(m *testing.M) {
	app = fiber.New()
	m.Run()
}

// func TestGet(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want Trace
// 	}{
// 		{
// 			name: "",
// 			want: Trace{
// 				TraceID:   "AAA",
// 				SpanID:    "BBB",
// 				TraceTrue: false,
// 			},
// 		},
// 		{
// 			name: "",
// 			want: Trace{
// 				TraceID:   "dv3BVPyTSi4JTOtUAi6b3rTfEUZD9a",
// 				SpanID:    "Yw1zYLcdLECSmSXoJfxm0c6GoZCMdW",
// 				TraceTrue: true,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
// 			ctx.Locals(KeyTrace, tt.want)
// 			if got := Get(ctx); !reflect.DeepEqual(got, tt.want) {
// 				tools.StructPrint(got)
// 				t.Errorf("Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Benchmark_UUID(b *testing.B) {
	var res string
	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = utils.UUID()
		}
		utils.AssertEqual(b, 36, len(res))
	})
	b.Run("default", func(b *testing.B) {
		rnd := make([]byte, 16)
		_, _ = rand.Read(rnd)
		for n := 0; n < b.N; n++ {
			res = fmt.Sprintf("%x-%x-%x-%x-%x", rnd[0:4], rnd[4:6], rnd[6:8], rnd[8:10], rnd[10:])
		}
		utils.AssertEqual(b, 36, len(res))
	})
	b.Run("google", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = utils.UUIDv4()
		}
		utils.AssertEqual(b, 36, len(res))
	})
}

func Test_auto(t *testing.T) {
	type args struct {
		c     *fiber.Ctx
		trace string
	}
	tests := []struct {
		name string
		args args
		want Trace
	}{
		{
			name: "",
			args: args{
				c:     app.AcquireCtx(&fasthttp.RequestCtx{}),
				trace: "105445aa7843bc8bf206b12000100000/1;o=1",
			},
			want: Trace{
				TraceID:   "105445aa7843bc8bf206b12000100000",
				SpanID:    "1",
				TraceTrue: true,
				Raw:       "105445aa7843bc8bf206b12000100000/1;o=1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.c.Request().Header.Add(XCloudTraceContext, tt.args.trace)
			if got := auto(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auto() = %v, want %v", got, tt.want)
			}
		})
	}
}
