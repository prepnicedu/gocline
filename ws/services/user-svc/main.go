// /workspaces/gocline/ws/services/user-svc/main.go
package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	userv1 "gocline.com/ws/gen/go/user/v1"
	"gocline.com/ws/pkg/core/config"
	database "gocline.com/ws/services/user-svc/database"
	"gocline.com/ws/services/user-svc/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	dbSvc, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}

	repository := internal.NewRepository(dbSvc.Database, cfg.MongoUserCollection)
	service, err := internal.NewService(repository)
	if err != nil {
		log.Fatalf("failed to initialized service: %v", err)
	}

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":"+cfg.UserSvcPort)
	if err != nil {
		log.Fatalf("error starting grpc [user-svc] server: %v", err)
	}

	go func() {
		log.Printf("user-svc[grpc] is listening on port: %v", cfg.UserSvcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc server [user-svc] failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("server stopped gracefully")

	grpcServer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeOut)
	defer cancel()
	if err := dbSvc.Disconnect(ctx); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}

}
