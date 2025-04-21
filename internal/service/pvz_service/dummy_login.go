package pvz_service

import (
	"context"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (PVZService *PVZServiceAPI) DummyLogin(ctx context.Context, req *desc.DummyLoginRequest) (*desc.DummyLoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unknown role")
	}

	token, err := PVZService.tokenizerInterface.GenerateToken("someUUID", req.GetRole())
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't generate token")
	}

	return makeDummyLoginResponse(token), nil
}

func makeDummyLoginResponse(token string) *desc.DummyLoginResponse {
	return &desc.DummyLoginResponse{
		Token: token,
	}
}
