package repo

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CreateProduct(ctx context.Context, product *products.Product) error {
	sb := psql.Insert("products")
	sb = sb.Values(product.ID, product.ReceptionTime, product.Type, product.ReceptionID, product.PreviousProductID)

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
