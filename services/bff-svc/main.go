package main

import (
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	productv1 "example.com/my_project/gen/go/product/v1"
	"example.com/my_project/gen/go/product/v1/productv1connect"
	"example.com/my_project/pkg/core"
	"example.com/my_project/services/bff-svc/service"
)

func main() {
	cfg, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("error initializing config: %v", err)
	}

	mux := http.NewServeMux()

	productConn, err := core.NewClient(cfg.ProductSvcAddr)
	if err != nil {
		log.Fatalf("failed to connect to product-svc: %v", err)
	}
	defer productConn.Close()
	productGrpcClient := productv1.NewProductServiceClient(productConn)
	productService := service.NewProductService(productGrpcClient)
	path, handler := productv1connect.NewProductServiceHandler(
		productService,
		connect.WithInterceptors(validate.NewInterceptor()),
	)
	mux.Handle(path, handler)

	srv := core.NewHttp(cfg.BffSvcPort, mux, 5*time.Second)
	srv.Run()
}
