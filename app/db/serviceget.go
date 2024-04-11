// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Connection, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get connection by id [%v]", id)
	}
	return ret.ToConnection(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Connections, error) {
	if len(ids) == 0 {
		return Connections{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Connections for [%d] ids", len(ids))
	}
	return ret.ToConnections(), nil
}

func (s *Service) GetByEngine(ctx context.Context, tx *sqlx.Tx, engine Engine, params *filter.Params, logger util.Logger) (Connections, error) {
	params = filters(params)
	wc := "\"engine\" = @p1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, engine)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Connections by engine [%v]", engine)
	}
	return ret.ToConnections(), nil
}

func (s *Service) GetByEngines(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, engines ...Engine) (Connections, error) {
	if len(engines) == 0 {
		return Connections{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("engine", len(engines), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(engines)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Connections for [%d] engines", len(engines))
	}
	return ret.ToConnections(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Connection, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "newid()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random connections")
	}
	return ret.ToConnection(), nil
}
