package controller

import (
	"net/http"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/views/vparse"
)

func ParseForm(w http.ResponseWriter, r *http.Request) {
	Act("parse.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Render(r, as, &vparse.SQLServer{Path: "/tmp/*.sqlaudit", Task: "testbed", Limit: 10}, ps)
	})
}

func Parse(w http.ResponseWriter, r *http.Request) {
	Act("parse", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}

		path := frm.GetStringOpt("path")
		task := frm.GetStringOpt("task")
		limit := frm.GetIntOpt("limit")

		res, err := as.Services.Parse.Testbed(ps.Context, path, limit, ps.Logger)
		if err != nil {
			return "", err
		}

		return Render(r, as, &vparse.SQLServer{Path: path, Task: task, Limit: limit, Result: res}, ps)
	})
}
