package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CheckPVZExisting(ctx context.Context, pvzID string) (bool, error) {
	sb := psql.Select("1").
		From("pvzs").
		Where(sq.Eq{"id": pvzID}).
		Prefix("SELECT EXISTS (").
		Suffix(") AS pvz_exists")

	query, args, err := sb.ToSql()
	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	var pvzExists bool
	err = repo.db.GetContext(ctx, &pvzExists, query, args...)

	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	return pvzExists, nil
}
