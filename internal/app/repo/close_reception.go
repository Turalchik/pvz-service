package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CloseReception(ctx context.Context, receptionID string, endTime sql.NullTime) error {
	sb := psql.Update("receptions").
		Set("end_time", endTime).
		Set("status", "close").
		Where(sq.Eq{"id": receptionID})

	query, args, err := sb.ToSql()
	if err != nil {
		return status.Error(codes.Internal, "Internal corrupt")
	}

	_, err = repo.db.ExecContext(ctx, query, args...)
	if err != nil {
		return status.Error(codes.Internal, "Internal corrupt")
	}
	return nil
}
