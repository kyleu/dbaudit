// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"net/http"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/user"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/verror"
)

func Options(w http.ResponseWriter, _ *http.Request) {
	cutil.WriteCORS(w)
	w.WriteHeader(http.StatusOK)
}

func NotFoundAction(w http.ResponseWriter, r *http.Request) {
	Act("notfound", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return NotFoundResponse(w, r)(as, ps)
	})
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		cutil.WriteCORS(w)
		w.WriteHeader(http.StatusNotFound)
		ps.Logger.Warnf("%s %s returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
		if ps.Title == "" {
			ps.Title = "Page not found"
		}
		ps.Data = util.ValueMap{"status": "notfound", "statusCode": http.StatusNotFound, "message": ps.Title}
		bc := util.StringSplitAndTrim(r.URL.Path, "/")
		bc = append(bc, "Not Found")
		return Render(w, r, as, &verror.NotFound{Path: r.URL.Path}, ps, bc...)
	}
}

func Unauthorized(w http.ResponseWriter, r *http.Request, reason string, accounts user.Accounts) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		cutil.WriteCORS(w)
		w.WriteHeader(http.StatusUnauthorized)
		path := r.URL.Path
		ps.Logger.Warnf("%s %s returned [%d]", r.Method, path, http.StatusUnauthorized)
		bc := util.StringSplitAndTrim(r.URL.Path, "/")
		bc = append(bc, "Unauthorized")
		if ps.Title == "" {
			ps.Title = "Unauthorized"
		}
		if reason == "" {
			if len(accounts) == 0 {
				reason = "not signed in"
			} else {
				reason = "no access"
			}
		}
		ps.Data = util.ValueMap{"status": "unauthorized", "statusCode": http.StatusUnauthorized, "message": reason}
		return Render(w, r, as, &verror.Unauthorized{Path: path, Message: reason, Accounts: accounts}, ps, bc...)
	}
}
