package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	productv1 "example.com/my_project/gen/go/product/v1"
	"example.com/my_project/pkg/core"
	"example.com/my_project/services/product-svc/cmd"
	"example.com/my_project/services/product-svc/internal/repository"
	"example.com/my_project/services/product-svc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dbSvc, err := cmd.Connect(cfg)
	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}

	repository := repository.NewRepository(dbSvc.Database, cfg.MongoProductCollection)
	productService, err := service.NewService(repository)
	if err != nil {
		log.Fatalf("error initializing product service: %v", err)
	}

	grpcServer := grpc.NewServer()
	productv1.RegisterProductServiceServer(grpcServer, productService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":"+cfg.ProductSvcPort)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.ProductSvcPort, err)
	}

	go func() {
		log.Printf("application started on port: %v", cfg.ProductSvcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	grpcServer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()
	if err := dbSvc.Disconnect(ctx); err != nil {
		log.Fatalf("error closing db connection: %v", err)
	}
	log.Println("application exited cleanly")

}
