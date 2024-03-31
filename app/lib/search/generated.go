// Package search - Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"

	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/db"
	"github.com/kyleu/dbaudit/app/lib/search/result"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
)

func generatedSearch() []Provider {
	dbFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("db", logger).WithLimit(5)
		models, err := as.Services.Connection.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *db.Connection, _ int) *result.Result {
			return result.NewResult("db", m.String(), m.WebPath(), m.TitleString(), "database", m, m, params.Q)
		}), nil
	}
	statementFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("statement", logger).WithLimit(5)
		models, err := as.Services.Statement.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *statement.Statement, _ int) *result.Result {
			return result.NewResult("statement", m.String(), m.WebPath(), m.TitleString(), "database", m, m, params.Q)
		}), nil
	}
	return []Provider{dbFunc, statementFunc}
}
