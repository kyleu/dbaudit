package controller

import (
	"net/http"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/views/vstatement"
)

func StatementRun(w http.ResponseWriter, r *http.Request) {
	Act("statement.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		stmt, err := statementFromPath(r, as, ps)
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
		return Render(r, as, page, ps, "statement", stmt.TitleString()+"**database", "Run**play")
	})
}
