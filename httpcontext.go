package httpserver

import (
	"github.com/gobestsdk/trace"
	"net/http"
	"strings"
	"time"
)

//HttpContext 作用
// 1 读写http双工
// 2 记录每个步骤耗费的时间
type HttpContext struct {
	Writer       http.ResponseWriter
	Realremoteip string
	Request      *http.Request
	//Query
	Query map[string]interface{}
	Trace string

	UsedTime trace.Span
	//记录上次时间
	t time.Time
	//Store 注意，线程并不安全，请勿跨线程使用
	Store map[string]interface{}
}

func CreateHttpContext(w http.ResponseWriter, r *http.Request, name string) (hctx *HttpContext) {
	hctx = new(HttpContext)
	hctx.Writer = w
	hctx.Request = r
	hctx.UsedTime = trace.NewStep(name)
	hctx.t = time.Now()
	hctx.Store = make(map[string]interface{})
	return
}

func (hctx *HttpContext) Reset() *HttpContext {
	hctx.Writer = nil
	hctx.Request = nil
	hctx.UsedTime = trace.NewStep("红包")
	hctx.t = time.Unix(0, 0)
	hctx.Store = make(map[string]interface{})
	return hctx
}
func (hctx *HttpContext) Name() (name string) {
	name = hctx.Request.Method + "_" + strings.ReplaceAll(hctx.Request.URL.String(), "/", "_")
	return
}
func (hctx *HttpContext) Step(name string) *HttpContext {
	hctx.UsedTime.NextStep(name, false)
	return hctx
}

func (hctx *HttpContext) HeaderIp(headeripkey string) *HttpContext {

	remote := hctx.Request.Header.Get(headeripkey)
	if len(remote) == 0 {
		remote = hctx.Request.Header.Get(headeripkey)
		if l := strings.Split(remote, ","); len(l) > 0 {
			remote = l[0]
		}
		if len(remote) == 0 {
			remote = hctx.Request.RemoteAddr
		}
	}
	hctx.Realremoteip = remote

	return hctx
}

func (hctx *HttpContext) HeaderTrace(headertracekey string) *HttpContext {
	traceid := hctx.Request.Header.Get(headertracekey)
	if len(traceid) == 0 {
		traceid = trace.NewtraceID(hctx.Realremoteip)
	}
	hctx.Trace = traceid
	return hctx
}
func (hctx *HttpContext) HeaderSpan(headerspankey string) *HttpContext {
	return hctx
}
