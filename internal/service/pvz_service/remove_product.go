package pvz_service

import (
	"context"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (PVZService *PVZServiceAPI) RemoveProduct(ctx context.Context, req *desc.RemoveProductRequest) (*emptypb.Empty, error) {
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
		return nil, status.Error(codes.Internal, "Can't get reception by pvzID")
	}

	if !reception.LastProductID.Valid {
		return nil, status.Error(codes.InvalidArgument, "PVZ doesn't have any product")
	}

	product, err := PVZService.repo.GetProductByID(ctx, reception.LastProductID.String)
	if err != nil {
		return nil, status.Error(codes.Internal, "Can't get last product by id")
	}

	if err = PVZService.repo.UpdateReceptionLastProduct(ctx, reception.ID, product.PreviousProductID); err != nil {
		return nil, status.Error(codes.Internal, "Can't get last product by id")
	}
	if err = PVZService.repo.DeleteProductByID(ctx, product.ID); err != nil {
		return nil, status.Error(codes.Internal, "Can't delete last product by id")
	}

	return nil, nil
}
