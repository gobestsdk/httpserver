package httpserver

import (
	"encoding/json"
	"goapp/log"
	"goapp/types"
	"io/ioutil"
	"regexp"
	"time"

	"net/http"
	"strconv"
)

func Getfilter(req *http.Request) (filter map[string]interface{}) {
	filter = make(map[string]interface{})
	for k, vs := range req.URL.Query() {
		if len(vs) > 1 || len(vs) < 1 {
			log.Warn(log.Fields{
				"req.URL.Query()": "len(vs)>1",
				"values":          vs,
			})
			continue
		}

		v := vs[0]

		intv, err := strconv.Atoi(v)
		if err == nil {
			filter[k] = intv
			continue
		}

		if m, _ := regexp.MatchString("(-?)d+.(d+)", v); m {
			floatv, _ := strconv.ParseFloat(v, 64)
			filter[k] = floatv
			continue
		}

		boolv, er := strconv.ParseBool(v)
		if er == nil {
			filter[k] = boolv
			continue
		}

		t, te := time.Parse(types.DateFormat, v)
		if te == nil {
			filter[k] = t
			continue
		}
		ct, cte := time.Parse(types.CommonDatetime, v)
		if cte == nil {
			filter[k] = ct
			continue
		}
		st, ste := time.Parse(types.SimpleDate, v)
		if ste == nil {
			filter[k] = st
			continue
		}

		filter[k] = v
	}
	return
}

func Options(req *http.Request, resp http.ResponseWriter, contenttype, server, methods string) {
	Header(resp, contenttype, server, methods)
	resp.Write([]byte(""))
}

func Header(resp http.ResponseWriter, contenttype, server, methods string) {

	resp.Header().Set("Access-Control-Allow-Origin", "*")                   //允许访问所有域
	resp.Header().Add("Access-Control-Allow-Headers", "Content-Type,token") //header的类型
	resp.Header().Set("content-type", contenttype)
	resp.Header().Set("HttpServer", server)

	resp.Header().Set("Access-Control-Allow-Methods", methods)
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
}
func To(req *http.Request, s interface{}) (err error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return err
	}
	return
}
