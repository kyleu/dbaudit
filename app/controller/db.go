// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/db"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/vdb"
)

func ConnectionList(rc *fasthttp.RequestCtx) {
	Act("db.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("db", nil, ps.Logger).Sanitize("db")
		var ret db.Connections
		var err error
		if q == "" {
			ret, err = as.Services.Connection.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Connection.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Connections", ret)
		page := &vdb.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "db")
	})
}

func ConnectionDetail(rc *fasthttp.RequestCtx) {
	Act("db.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Connection)", ret)

		return Render(rc, as, &vdb.Detail{Model: ret}, ps, "db", ret.TitleString()+"**database")
	})
}

func ConnectionCreateForm(rc *fasthttp.RequestCtx) {
	Act("db.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &db.Connection{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = db.Random()
		}
		ps.SetTitleAndData("Create [Connection]", ret)
		ps.Data = ret
		return Render(rc, as, &vdb.Edit{Model: ret, IsNew: true}, ps, "db", "Create")
	})
}

func ConnectionRandom(rc *fasthttp.RequestCtx) {
	Act("db.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Connection.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Connection")
		}
		return ret.WebPath(), nil
	})
}

func ConnectionCreate(rc *fasthttp.RequestCtx) {
	Act("db.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Connection from form")
		}
		err = as.Services.Connection.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Connection")
		}
		msg := fmt.Sprintf("Connection [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func ConnectionEditForm(rc *fasthttp.RequestCtx) {
	Act("db.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vdb.Edit{Model: ret}, ps, "db", ret.String())
	})
}

func ConnectionEdit(rc *fasthttp.RequestCtx) {
	Act("db.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := dbFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Connection from form")
		}
		frm.ID = ret.ID
		err = as.Services.Connection.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Connection [%s]", frm.String())
		}
		msg := fmt.Sprintf("Connection [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func ConnectionDelete(rc *fasthttp.RequestCtx) {
	Act("db.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := dbFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Connection.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete connection [%s]", ret.String())
		}
		msg := fmt.Sprintf("Connection [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/db", rc, ps)
	})
}

func dbFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*db.Connection, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func dbFromForm(rc *fasthttp.RequestCtx, setPK bool) (*db.Connection, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return db.FromMap(frm, setPK)
}
