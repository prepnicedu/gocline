package main

import (
	"log"
	"net/http"
	"time"

	userv1 "gocline.com/ws/gen/go/user/v1"
	"gocline.com/ws/gen/go/user/v1/userv1connect"
	"gocline.com/ws/pkg/core/config"
	grpcclient "gocline.com/ws/pkg/core/grpc-client"
	httpserver "gocline.com/ws/pkg/core/http-server"
	"gocline.com/ws/services/bff-svc/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	mux := http.NewServeMux()

	userConn, err := grpcclient.NewClient(cfg.UserSvcAddr)
	if err != nil {
		log.Fatalf("failed to connect to user-svc: %v", err)
	}
	defer userConn.Close()

	userGrpcClient := userv1.NewUserServiceClient(userConn)
	userService := service.NewUserService(userGrpcClient)

	path, handler := userv1connect.NewUserServiceHandler(
		userService,
	)
	mux.Handle(path, handler)

	srv := httpserver.New(cfg.BffSvcPort, mux, 5*time.Second)
	srv.Run()
}
