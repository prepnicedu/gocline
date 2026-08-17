// ws/services/bff-svc/service/user.go
package service

import (
	"context"
	"fmt"
	userv1 "gen/go/user/v1"
	"gen/go/user/v1/userv1connect"

	"connectrpc.com/connect"
)

type UserService struct {
	userClient userv1.UserServiceClient
	userv1connect.UnimplementedUserServiceHandler
}

func NewUserService(userClient userv1.UserServiceClient) *UserService {
	return &UserService{
		userClient: userClient,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	res, err := s.userClient.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("user-svc error: %w", err),
		)
	}
	return connect.NewResponse(res), nil
}
