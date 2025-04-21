package pvz_service

import (
	"context"
	"database/sql"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (PVZService *PVZServiceAPI) OpenReception(ctx context.Context, req *desc.OpenReceptionRequest) (*desc.OpenReceptionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	claims, err := PVZService.tokenizerInterface.VerifyToken(req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	if claims.Role != "сотрудник ПВЗ" {
		return nil, status.Error(codes.PermissionDenied, "Insufficient permissions")
	}

	pvzExists, err := PVZService.repo.CheckPVZExisting(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't check existing PVZ")
	}
	if !pvzExists {
		return nil, status.Error(codes.InvalidArgument, "PVZ doesn't exist")
	}

	hasActiveReception, err := PVZService.repo.CheckReceptionActive(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't check PVZ on active reception")
	}
	if hasActiveReception {
		return nil, status.Error(codes.InvalidArgument, "PVZ has active reception")
	}

	newReception := &receptions.Reception{
		ID:        PVZService.uuidInterface.NewString(),
		StartTime: PVZService.timerInterface.Now(),
		PVZID:     req.Id,
		Status:    "in_progress",
	}

	newActiveReceptionID := sql.NullString{String: newReception.ID, Valid: true}
	if err = PVZService.repo.UpdateActiveReceptionPVZ(ctx, req.GetId(), newActiveReceptionID); err != nil {
		return nil, status.Error(codes.Internal, "Can't update active reception for PVZ")
	}
	if err = PVZService.repo.CreateReception(ctx, newReception); err != nil {
		return nil, status.Error(codes.Internal, "Can't open reception")
	}

	return makeOpenReceptionResponse(newReception.ID), nil
}

func makeOpenReceptionResponse(receptionID string) *desc.OpenReceptionResponse {
	return &desc.OpenReceptionResponse{
		IdReception: receptionID,
	}
}
