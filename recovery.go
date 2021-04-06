package httpserver

import (
	"encoding/json"
	"goapp/log"
	"net"
	"os"
	"strings"
)

//Recovery 捕获panic
func Recovery(ctx *HttpContext) {

	err := recover()
	if err != nil {
		var brokenPipe bool
		if ne, ok := err.(*net.OpError); ok {
			if se, ok := ne.Err.(*os.SyscallError); ok {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}
		if brokenPipe {
			//连接 已断开
			log.Warn(log.Fields{
				"http":     "connect broke",
				"remoteip": ctx.Realremoteip,
				"trace":    ctx.Trace,
				"span":     ctx.UsedTime,
			})
			return
		} else {
			//服务器内部错误
			ctx.Writer.WriteHeader(500)
			respdata := make(map[string]interface{})
			respdata["msg"] = "服务器遇到了panic"
			respdata["panic"] = err

			log.Error(log.Fields{
				"panic":    "connect broke",
				"remoteip": ctx.Realremoteip,
				"trace":    ctx.Trace,
				"usedtime": ctx.UsedTime,
			})
			bs, _ := json.Marshal(respdata)
			ctx.Writer.Write(bs)
		}
		return
	}
}
