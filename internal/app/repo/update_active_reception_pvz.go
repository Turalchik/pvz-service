package repo

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) UpdateActiveReceptionPVZ(ctx context.Context, pvzID string, receptionID sql.NullString) error {
	sb := psql.Update("pvzs").
		Set("active_reception", receptionID).
		Where(squirrel.Eq{"id": pvzID})

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
