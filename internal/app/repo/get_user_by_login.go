package repo

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) GetUserByLogin(ctx context.Context, login string) (*users.User, error) {
	sb := psql.Select("*").
		From("users").
		Where(sq.Eq{"login": login})

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	var user = &users.User{}
	err = repo.db.GetContext(ctx, user, query, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, "Cannot take user by login")
	}
	return user, nil
}
