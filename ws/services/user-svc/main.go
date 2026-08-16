// ws/services/user-svc/main.go
package main

import (
	"context"
	userv1 "gen/go/user/v1"
	"log"
	"net"
	"os"
	"os/signal"
	"services/user-svc/config"
	"services/user-svc/internal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dbSvc, err := config.Connect(cfg)
	if err != nil {
		log.Fatalf("error starting db: %v", err)
	}

	repository := internal.NewRepository(dbSvc.Database, cfg.MongoCollection)
	service, err := internal.NewService(repository)
	if err != nil {
		log.Fatalf("failed to initialize service: %v", err)
	}

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.Port, err)
	}

	go func() {
		log.Printf("application started on port: %v", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	grpcServer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := dbSvc.Disconnect(ctx); err != nil {
		log.Fatalf("error closing db: %v", err)

	}
	log.Println("Server gracefully stopped.")

}
