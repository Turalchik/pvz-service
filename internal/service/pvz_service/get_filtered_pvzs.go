package pvz_service

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (PVZService *PVZServiceAPI) GetFilteredPVZs(ctx context.Context, req *desc.GetFilteredPVZsRequest) (*desc.GetFilteredPVZsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	claims, err := PVZService.tokenizerInterface.VerifyToken(req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	if claims.Role != "сотрудник ПВЗ" && claims.Role != "модератор" {
		return nil, status.Error(codes.PermissionDenied, "Insufficient permissions")
	}

	start := req.GetStart().AsTime()
	finish := req.GetFinish().AsTime()

	PVZs, err := PVZService.repo.GetFilteredPVZs(ctx, start, finish, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return makeGetFilteredPVZsResponse(PVZs), nil
}

func makeGetFilteredPVZsResponse(PVZs []*pvz.PVZ) *desc.GetFilteredPVZsResponse {
	return &desc.GetFilteredPVZsResponse{
		Pvzs: mapModelPVZsToPB(PVZs),
	}
}

func mapModelPVZsToPB(in []*pvz.PVZ) []*desc.PVZ {
	out := make([]*desc.PVZ, 0, len(in))
	for _, m := range in {

		out = append(out, &desc.PVZ{
			Id:               m.ID,
			RegistrationDate: timestamppb.New(m.RegistrationDate),
			City:             m.City,
		})
	}
	return out
}
