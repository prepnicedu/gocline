///workspaces/gocline/ws/services/user-svc/internal/service.go
package internal

import (
	"context"
	"fmt"

	"buf.build/go/protovalidate"
	userv1 "gocline.com/ws/gen/go/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo      *Repository
	validator protovalidate.Validator
	userv1.UnimplementedUserServiceServer
}

func NewService(repo *Repository) (*Service, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize protovalidate: %w", err)
	}

	return &Service{
		repo:      repo,
		validator: v,
	}, nil
}

func (s *Service) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid request: %v",
			err,
		)
	}
	savedUser, err := s.repo.Create(ctx, User{Name: req.GetName()})
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &userv1.CreateUserResponse{
		Id:   savedUser.ID.Hex(),
		Name: savedUser.Name,
	}, nil
}
