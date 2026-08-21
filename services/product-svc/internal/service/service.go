package service

import (
	"context"
	"fmt"

	"buf.build/go/protovalidate"
	productv1 "example.com/my_project/gen/go/product/v1"
	"example.com/my_project/services/product-svc/internal/domain"
	"example.com/my_project/services/product-svc/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo      *repository.Repository
	validator protovalidate.Validator
	productv1.UnimplementedProductServiceServer
}

func NewService(repo *repository.Repository) (*Service, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, fmt.Errorf("error initializing protovalidate: %w", err)
	}
	return &Service{
		repo:      repo,
		validator: v,
	}, nil
}

func (s *Service) CreateProduct(
	ctx context.Context,
	req *productv1.CreateProductRequest,
) (
	*productv1.CreateProductResponse,
	error,
) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"create product: %v",
			err,
		)
	}

	res, err := s.repo.Create(ctx, domain.Product{
		Title: req.GetTitle(),
		Description: req.GetDescription(),
		ProductId: req.GetProductId(),
		Qty: req.GetQty(),
	})
	if err != nil {
		return nil, fmt.Errorf("create product failed: %w", err)
	}

	return &productv1.CreateProductResponse{
		Id: res.ID.Hex(),
		Product: &productv1.Product{
			Title:       res.Title,
			Description: res.Description,
			ProductId:   res.ProductId,
			Qty:         res.Qty,
		},
	}, nil
}
