package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	productv1 "example.com/my_project/gen/go/product/v1"
	"example.com/my_project/gen/go/product/v1/productv1connect"
)

type ProductService struct {
	productClient productv1.ProductServiceClient
	productv1connect.UnimplementedProductServiceHandler
}

func NewProductService(
	productClient productv1.ProductServiceClient,
) *ProductService {
	return &ProductService{
		productClient: productClient,
	}
}

func (s *ProductService) CreateProduct(
	ctx context.Context,
	req *connect.Request[productv1.CreateProductRequest],
) (
	*connect.Response[productv1.CreateProductResponse],
	error,
) {
	res, err := s.productClient.CreateProduct(
		ctx,
		req.Msg,
	)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("create product: %w", err),
		)
	}
	return connect.NewResponse(res), nil
}
