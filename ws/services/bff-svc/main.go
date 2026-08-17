// ws/services/bff-svc/main.go
package main

import (
	userv1 "gen/go/user/v1"
	"gen/go/user/v1/userv1connect"
	"log"
	"net/http"
	grpcclient "pkg/grpc-client"
	"services/bff-svc/config"
	"services/bff-svc/service"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	mux := http.NewServeMux()

	// dial user grpc connection
	userConn, err := grpcclient.NewClient(cfg.UserSvcAddr)
	if err != nil {
		log.Fatalf("failed to connect to user-svc: %v", err)
	}
	defer userConn.Close()
	userGrpcClient := userv1.NewUserServiceClient(userConn)
	userService := service.NewUserService(userGrpcClient)
	path, handler := userv1connect.NewUserServiceHandler(userService)
	mux.Handle(path, handler)

	srv := grpcclient.New(cfg.Port, mux, 10*time.Second)
	srv.Run()
}
