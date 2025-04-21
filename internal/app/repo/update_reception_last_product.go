package repo

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) UpdateReceptionLastProduct(ctx context.Context, receptionID string, productID sql.NullString) error {
	sb := psql.Update("receptions").
		Set("last_product_id", productID).
		Where(squirrel.Eq{"id": receptionID})

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
