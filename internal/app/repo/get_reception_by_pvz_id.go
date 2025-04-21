package repo

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Turalchik/pvz-service/internal/entities/receptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (repo *Repo) GetReceptionByPVZID(ctx context.Context, pvzID string) (*receptions.Reception, error) {
	sb := psql.Select("receptions.*").
		From("receptions").
		Join("pvzs ON receptions.id = pvzs.active_reception").
		Where(sq.Eq{"pvzs.id": pvzID})

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	var reception = &receptions.Reception{}
	err = repo.db.GetContext(ctx, reception, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Reception doesn't exist")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return reception, nil
}
