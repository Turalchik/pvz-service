package repo

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Turalchik/pvz-service/internal/entities/pvz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (repo *Repo) GetFilteredPVZs(ctx context.Context, startDate time.Time, endDate time.Time, limit uint64, offset uint64) ([]*pvz.PVZ, error) {
	sb := sq.Select("DISTINCT pvzs.*").
		From("pvzs").
		Join("receptions ON receptions.pvz_id = pvzs.id").
		Where(sq.LtOrEq{"receptions.start_time": endDate}).
		Where(
			sq.Or{
				sq.Expr("receptions.end_time IS NULL"),
				sq.GtOrEq{"receptions.end_time": startDate},
			},
		).
		OrderBy("pvzs.registration_date DESC").
		Limit(limit).
		Offset(offset)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal corrupt")
	}

	var pvzs []*pvz.PVZ
	err = repo.db.SelectContext(ctx, &pvzs, query, args...)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.Internal, "failed to fetch PVZ list")
	}

	return pvzs, nil
}
