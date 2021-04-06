package httpserver

import (
	"testing"
	"time"
)

func TestServer_Run(t *testing.T) {
	s := New("test")
	s.SetPort(":8090")
	go s.Run()

	s.Waitstop()
	time.Sleep(time.Second * 5)
}
