package repo

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) GetProductByID(ctx context.Context, productID string) (*products.Product, error) {
	sb := psql.Select("*").
		From("products").
		Where(sq.Eq{"id": productID})

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	var product = &products.Product{}
	err = repo.db.GetContext(ctx, product, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Reception doesn't exist")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return product, nil
}
