// ws/pkg/grpc-client/http-server.go
package grpcclient

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer http.Server
	timeOut    time.Duration
}

func New(
	addr string,
	handler http.Handler,
	shutDownTimeOut time.Duration,
) *Server {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	return &Server{
		httpServer: http.Server{
			Addr:         ":"+addr,
			Handler:      handler,
			Protocols:    p,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
		timeOut: shutDownTimeOut,
	}
}

func (s *Server) Run() {

	go func() {
		log.Printf("server listening on port: %v", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("server shutdown gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("server stopped cleanly")
}
