package httpserver

import (
	"context"
	"github.com/gobestsdk/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
)

// HttpServer http的server

type HttpServer struct {
	name string //服务名称

	Server http.Server
	Mux    *http.ServeMux

	quitChan    chan os.Signal
	Runflag     bool
	quitTimeout time.Duration
}

// New 生产Server实例
func New(serverName string) *HttpServer {

	s := &HttpServer{
		name:        serverName,
		Server:      http.Server{},
		Runflag:     true,
		quitChan:    make(chan os.Signal),
		quitTimeout: 5 * time.Second,
	}
	s.Mux = new(http.ServeMux)
	s.Server.Handler = s.Mux
	return s
}

// SetPort 设置服务端口
func (s *HttpServer) SetPort(port string) {
	s.Server.Addr = port
}

// Run server on addr
func (s *HttpServer) Run() {

	go s.httpServer()
	<-s.quitChan
}

func (s *HttpServer) httpServer() {
	log.Info(log.Fields{"server " + s.name + "ListenAndServe addr:": s.Server.Addr})
	err := s.Server.ListenAndServe()
	if err != nil {
		log.Error(log.Fields{"app": "ListenAndServe failed", "error": err.Error()})
	}
}

// Waitstop 阻塞主线程,直到进程结束
func (s *HttpServer) Waitstop() {
	s.Runflag = false
s:
	signal.Notify(s.quitChan,
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)
	sig := <-s.quitChan
	log.Info(log.Fields{
		"signal": sig,
	})
	switch sig {
	case syscall.SIGUSR1:
		fallthrough
	case syscall.SIGUSR2:
		log.Loggoid = 0
		goto s
	default:

	}

	s.Stop()
}

// Waitstop 停止server
func (s *HttpServer) Stop() {
	s.Runflag = false
	ctx, cancel := context.WithTimeout(context.Background(), s.quitTimeout)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Error(log.Fields{
			"app": "Shutdown HttpServer",
			"err": err.Error(),
		})
	}
	log.Info(log.Fields{
		"msg": "exiting...",
	})
	close(s.quitChan)
}
