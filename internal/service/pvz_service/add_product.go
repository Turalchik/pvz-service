package pvz_service

import (
	"context"
	"database/sql"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (PVZService *PVZServiceAPI) AddProduct(ctx context.Context, req *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unknown product type")
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

	newProduct := &products.Product{
		ID:                PVZService.uuidInterface.NewString(),
		ReceptionTime:     PVZService.timerInterface.Now(),
		Type:              req.GetType(),
		ReceptionID:       reception.ID,
		PreviousProductID: reception.LastProductID,
	}

	newLastProductID := sql.NullString{
		String: newProduct.ID,
		Valid:  true,
	}
	if err = PVZService.repo.UpdateReceptionLastProduct(ctx, reception.ID, newLastProductID); err != nil {
		return nil, status.Error(codes.Internal, "Can't update reception last product")
	}
	if err = PVZService.repo.CreateProduct(ctx, newProduct); err != nil {
		return nil, status.Error(codes.Internal, "Can't create new product")
	}

	return makeProductResponse(newProduct.ID), nil
}

func makeProductResponse(productID string) *desc.AddProductResponse {
	return &desc.AddProductResponse{
		IdProduct: productID,
	}
}
