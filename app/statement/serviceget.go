// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

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

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Statements, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get statements")
	}
	return ret.ToStatements(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of statements")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Statement, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get statement by id [%v]", id)
	}
	return ret.ToStatement(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Statements, error) {
	if len(ids) == 0 {
		return Statements{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Statements for [%d] ids", len(ids))
	}
	return ret.ToStatements(), nil
}

func (s *Service) GetByAction(ctx context.Context, tx *sqlx.Tx, action Action, params *filter.Params, logger util.Logger) (Statements, error) {
	params = filters(params)
	wc := "\"action\" = @p1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, action)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Statements by action [%v]", action)
	}
	return ret.ToStatements(), nil
}

func (s *Service) GetByActions(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, actions ...Action) (Statements, error) {
	if len(actions) == 0 {
		return Statements{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("action", len(actions), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(actions)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Statements for [%d] actions", len(actions))
	}
	return ret.ToStatements(), nil
}

const searchClause = "(lower(cast(id as nvarchar(2048))) like @p1 or lower(cast(action as nvarchar(2048))) like @p1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Statements, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToStatements(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Statements, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get statements using custom SQL")
	}
	return ret.ToStatements(), nil
}

func (s *Service) ListWhere(ctx context.Context, tx *sqlx.Tx, where string, params *filter.Params, logger util.Logger, values ...any) (Statements, error) {
	params = filters(params)
	sql := database.SQLSelect(columnsString, tableQuoted, where, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	return s.ListSQL(ctx, tx, sql, logger, values...)
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Statement, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "newid()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random statements")
	}
	return ret.ToStatement(), nil
}
