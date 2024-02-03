package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/views/vparse"
)

func ParseForm(rc *fasthttp.RequestCtx) {
	Act("parse.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Render(rc, as, &vparse.SQLServer{Path: "/tmp/*.sqlaudit", Task: "testbed", Limit: 10}, ps)
	})
}

func Parse(rc *fasthttp.RequestCtx) {
	Act("parse", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
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

		return Render(rc, as, &vparse.SQLServer{Path: path, Task: task, Limit: limit, Result: res}, ps)
	})
}
