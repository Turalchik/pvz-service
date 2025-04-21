package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) DeleteProductByID(ctx context.Context, productID string) error {
	sb := psql.Delete("products").
		Where(squirrel.Eq{"id": productID})

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
