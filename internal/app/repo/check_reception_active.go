package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CheckReceptionActive(ctx context.Context, pvzID string) (bool, error) {
	sb := psql.Select("1").
		From("pvzs").
		Where(sq.And{
			sq.Eq{"id": pvzID},
			sq.NotEq{"active_reception": nil},
		}).
		Prefix("SELECT EXISTS (").
		Suffix(") AS has_active_reception")

	query, args, err := sb.ToSql()
	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	var hasActiveReception bool
	err = repo.db.GetContext(ctx, &hasActiveReception, query, args...)
	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	return hasActiveReception, nil
}
