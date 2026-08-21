package core

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
	httpServer *http.Server
	timeout    time.Duration
}

func NewHttp(
	addr string,
	handler http.Handler,
	timeout time.Duration,
) *Server {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	return &Server{
		httpServer: &http.Server{
			Addr:         ":"+addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		timeout: timeout,
	}
}

func (s *Server) Run() {
	go func() {
		log.Printf("server started on port: %v", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server closed forcefully: %v", err)
	}
	log.Println("server exited cleanly")

}
