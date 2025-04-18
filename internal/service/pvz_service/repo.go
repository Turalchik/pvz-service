package pvz_service

import (
	"context"
	"github.com/Turalchik/pvz-service/internal/entities/users"
)

type RepoInterface interface {
	CheckUserExisting(ctx context.Context, login string) (bool, error)
	CreateUser(ctx context.Context, user *users.User) error
}
