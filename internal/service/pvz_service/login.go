package pvz_service

import (
	"context"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (PVZService *PVZServiceAPI) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	user, err := PVZService.repo.GetUserByLogin(ctx, req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not exist")
	}

	if user.Password != req.GetPassword() {
		return nil, status.Error(codes.InvalidArgument, "Wrong password")
	}

	token, err := PVZService.tokenizerInterface.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't generate token")
	}

	return makeLoginResponse(token), nil
}

func makeLoginResponse(token string) *desc.LoginResponse {
	return &desc.LoginResponse{
		Token: token,
	}
}
