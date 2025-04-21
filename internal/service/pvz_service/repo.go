package pvz_service

import (
	"context"
	"database/sql"
	"github.com/Turalchik/pvz-service/internal/entities/products"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	"github.com/Turalchik/pvz-service/internal/entities/users"
	"time"
)

type RepoInterface interface {
	CheckUserExisting(ctx context.Context, login string) (bool, error)
	CreateUser(ctx context.Context, user *users.User) error
	GetUserByLogin(ctx context.Context, login string) (*users.User, error)
	CreatePVZ(ctx context.Context, pvz *pvz.PVZ) error
	CreateReception(ctx context.Context, reception *receptions.Reception) error
	CheckReceptionActive(ctx context.Context, pvzID string) (bool, error)
	CheckPVZExisting(ctx context.Context, pvzID string) (bool, error)
	UpdateActiveReceptionPVZ(ctx context.Context, pvzID string, receptionID sql.NullString) error
	GetReceptionByPVZID(ctx context.Context, pvzID string) (*receptions.Reception, error)
	UpdateReceptionLastProduct(ctx context.Context, receptionID string, productID sql.NullString) error
	CreateProduct(ctx context.Context, product *products.Product) error
	GetProductByID(ctx context.Context, productID string) (*products.Product, error)
	DeleteProductByID(ctx context.Context, productID string) error
	CloseReception(ctx context.Context, receptionID string, endTime sql.NullTime) error
	GetFilteredPVZs(ctx context.Context, startDate, endDate time.Time, limit, offset uint64) ([]*pvz.PVZ, error)
}
