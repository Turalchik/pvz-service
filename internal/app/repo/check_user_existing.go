package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CheckUserExisting(ctx context.Context, login string) (bool, error) {
	sb := psql.Select("1").
		From("users").
		Where(sq.Eq{"login": login}).
		Prefix("SELECT EXISTS (").
		Suffix(") AS user_exists")

	query, args, err := sb.ToSql()
	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	var userExists bool
	err = repo.db.GetContext(ctx, &userExists, query, args...)

	if err != nil {
		return false, status.Error(codes.Internal, "Internal error")
	}

	return userExists, nil
}
