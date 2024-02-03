package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/views/vstatement"
)

func StatementRun(rc *fasthttp.RequestCtx) {
	Act("statement.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		stmt, err := statementFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(stmt.TitleString()+" (Statement)", stmt)

		params := database.SQLServerParamsFromEnv("", "", "")
		params.Database = stmt.Database
		db, err := database.OpenSQLServerDatabase(ps.Context, "test", params, ps.Logger)
		if err != nil {
			return "", err
		}

		tx, err := db.StartTransaction(ps.Logger)
		if err != nil {
			return "", err
		}
		_, err = db.Exec(ps.Context, "set showplan_all on", tx, -1, ps.Logger)
		if err != nil {
			return "", err
		}
		ret, err := db.QueryRows(ps.Context, stmt.PlainSQL(), tx, ps.Logger)
		if err != nil {
			return "", err
		}
		_, err = db.Exec(ps.Context, "set showplan_all off", tx, -1, ps.Logger)
		if err != nil {
			return "", err
		}

		page := &vstatement.Result{Statement: stmt, Result: ret}
		return Render(rc, as, page, ps, "statement", stmt.TitleString()+"**database", "Run**play")
	})
}
