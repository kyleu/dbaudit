// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/db"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/vdb"
)

func ConnectionList(w http.ResponseWriter, r *http.Request) {
	Act("db.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("db", ps.Logger)
		var ret db.Connections
		var err error
		if q == "" {
			ret, err = as.Services.Connection.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Connection.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), w, ps)
			}
		}
		ps.SetTitleAndData("Connections", ret)
		page := &vdb.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(w, r, as, page, ps, "db")
	})
}

func ConnectionDetail(w http.ResponseWriter, r *http.Request) {
	Act("db.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Connection)", ret)

		return Render(w, r, as, &vdb.Detail{Model: ret}, ps, "db", ret.TitleString()+"**database")
	})
}

func ConnectionCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("db.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &db.Connection{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = db.Random()
		}
		ps.SetTitleAndData("Create [Connection]", ret)
		ps.Data = ret
		return Render(w, r, as, &vdb.Edit{Model: ret, IsNew: true}, ps, "db", "Create")
	})
}

func ConnectionRandom(w http.ResponseWriter, r *http.Request) {
	Act("db.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Connection.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Connection")
		}
		return ret.WebPath(), nil
	})
}

func ConnectionCreate(w http.ResponseWriter, r *http.Request) {
	Act("db.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Connection from form")
		}
		err = as.Services.Connection.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Connection")
		}
		msg := fmt.Sprintf("Connection [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func ConnectionEditForm(w http.ResponseWriter, r *http.Request) {
	Act("db.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(w, r, as, &vdb.Edit{Model: ret}, ps, "db", ret.String())
	})
}

func ConnectionEdit(w http.ResponseWriter, r *http.Request) {
	Act("db.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := dbFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Connection from form")
		}
		frm.ID = ret.ID
		err = as.Services.Connection.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Connection [%s]", frm.String())
		}
		msg := fmt.Sprintf("Connection [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func ConnectionDelete(w http.ResponseWriter, r *http.Request) {
	Act("db.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Connection.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete connection [%s]", ret.String())
		}
		msg := fmt.Sprintf("Connection [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/db", w, ps)
	})
}

func dbFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*db.Connection, error) {
	idArgStr, err := cutil.RCRequiredString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Connection.Get(ps.Context, nil, idArg, ps.Logger)
}

func dbFromForm(r *http.Request, b []byte, setPK bool) (*db.Connection, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return db.FromMap(frm, setPK)
}
