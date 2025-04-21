package repo

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CreateReception(ctx context.Context, reception *receptions.Reception) error {
	sb := psql.Insert("receptions").Columns("id", "start_time", "pvz_id", "status")
	sb = sb.Values(reception.ID, reception.StartTime, reception.PVZID, reception.Status)

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
