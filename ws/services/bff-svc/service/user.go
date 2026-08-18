package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	userv1 "gocline.com/ws/gen/go/user/v1"
	"gocline.com/ws/gen/go/user/v1/userv1connect"
)

type userService struct {
	userv1connect.UnimplementedUserServiceHandler
	userClient userv1.UserServiceClient
}

func NewUserService(
	userClient userv1.UserServiceClient,
) *userService {
	return &userService{
		userClient: userClient,
	}
}

func (s *userService) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (
	*connect.Response[userv1.CreateUserResponse],
	error,
) {
	res, err := s.userClient.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("create user: %w", err),
		)
	}
	return connect.NewResponse(res), nil
}
