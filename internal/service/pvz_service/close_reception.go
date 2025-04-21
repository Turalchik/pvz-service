package pvz_service

import (
	"context"
	"database/sql"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (PVZService *PVZServiceAPI) CloseReception(ctx context.Context, req *desc.CloseReceptionRequest) (*emptypb.Empty, error) {
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
	if !hasActiveReception {
		return nil, status.Error(codes.InvalidArgument, "PVZ doesn't have active reception")
	}

	reception, err := PVZService.repo.GetReceptionByPVZID(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't get reception")
	}

	if err = PVZService.repo.UpdateActiveReceptionPVZ(ctx, req.GetId(), sql.NullString{}); err != nil {
		return nil, status.Error(codes.Internal, "Can't update active reception for PVZ")
	}

	if err = PVZService.repo.CloseReception(ctx, reception.ID, sql.NullTime{Time: PVZService.timerInterface.Now(), Valid: true}); err != nil {
		return nil, status.Error(codes.Internal, "Can't close reception")
	}

	return nil, nil
}
