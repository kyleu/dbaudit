// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Connections, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get connections")
	}
	return ret.ToConnections(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of connections")
	}
	return int(ret), nil
}

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

const searchClause = "(lower(cast(id as nvarchar(2048))) like @p1 or lower(cast(engine as nvarchar(2048))) like @p1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Connections, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToConnections(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Connections, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get connections using custom SQL")
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
