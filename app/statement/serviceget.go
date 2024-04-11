// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

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

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Statement, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "newid()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random statements")
	}
	return ret.ToStatement(), nil
}
