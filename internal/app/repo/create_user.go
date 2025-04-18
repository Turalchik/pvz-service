package repo

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) CreateUser(ctx context.Context, user *users.User) error {
	sb := psql.Insert("users")
	sb = sb.Values(user.ID, user.Login, user.Password, user.Role)

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
