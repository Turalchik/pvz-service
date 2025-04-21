package pvz_service

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (PVZService *PVZServiceAPI) CreatePVZ(ctx context.Context, req *desc.CreatePVZRequest) (*desc.CreatePVZResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unknown city")
	}

	claims, err := PVZService.tokenizerInterface.VerifyToken(req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	if claims.Role != "модератор" {
		return nil, status.Error(codes.PermissionDenied, "Insufficient permissions")
	}

	newPVZ := &pvz.PVZ{
		ID:               PVZService.uuidInterface.NewString(),
		RegistrationDate: PVZService.timerInterface.Now(),
		City:             req.GetCity(),
	}

	if err := PVZService.repo.CreatePVZ(ctx, newPVZ); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return makeCreatePVZResponse(newPVZ), nil
}

func makeCreatePVZResponse(pvz *pvz.PVZ) *desc.CreatePVZResponse {
	return &desc.CreatePVZResponse{
		Id: pvz.ID,
	}
}
