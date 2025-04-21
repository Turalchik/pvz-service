package repo

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CreatePVZ(ctx context.Context, pvz *pvz.PVZ) error {
	sb := psql.Insert("pvzs")
	sb = sb.Values(pvz.ID, pvz.RegistrationDate, pvz.City)

	query, args, err := sb.ToSql()
	if err != nil {
		return status.Error(codes.Internal, "Internal corrupt")
	}

	_, err = repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return status.Error(codes.Internal, "Internal corrupt")
	}
	return nil
}
