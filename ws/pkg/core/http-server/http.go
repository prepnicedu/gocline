package httpserver

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
	timeout    time.Duration
}

func New(
	addr string,
	handler http.Handler,
	shuutDownTimeout time.Duration,
) *Server {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	return &Server{
		httpServer: http.Server{
			Addr:         ":" + addr,
			Handler:      handler,
			Protocols:    p,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		timeout: shuutDownTimeout,
	}
}

func (s *Server) Run() {
	go func() {
		log.Printf("server started on port: %v", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("bff-svc stopped gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("forcefully closing server: %v", err)
	}
	log.Println("bff-svc exited cleanly")

}
