package httpserver

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetfilter(t *testing.T) {
	req := http.Request{
		URL: &url.URL{Path: string("/?a=1&b=0.1&c=true&d=152123523235235325&e=2020-1-1&f=2021-03-12 12:31:31")},
	}
	fmt.Println(Getfilter(&req))
}
