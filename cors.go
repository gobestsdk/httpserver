package httpserver

import "net/http"

func Cros(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")                            //允许访问所有域
	writer.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,POST,PUT,DELETE") //允许访问所有域
	writer.Header().Add("Access-Control-Allow-Headers", "*")                           //header的类型
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
}
