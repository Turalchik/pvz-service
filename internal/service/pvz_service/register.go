package pvz_service

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (PVZService *PVZServiceAPI) Register(ctx context.Context, req *desc.RegisterRequest) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	userExists, err := PVZService.repo.CheckUserExisting(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, status.Errorf(codes.InvalidArgument, "User exists")
	}

	newUser := &users.User{
		ID:       uuid.NewString(),
		Login:    req.Login,
		Password: req.Password,
		Role:     req.Role,
	}
	err = PVZService.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
