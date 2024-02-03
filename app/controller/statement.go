// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/vstatement"
)

func StatementList(rc *fasthttp.RequestCtx) {
	Act("statement.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("statement", nil, ps.Logger).Sanitize("statement")
		var ret statement.Statements
		var err error
		if q == "" {
			ret, err = as.Services.Statement.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Statement.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Statements", ret)
		page := &vstatement.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "statement")
	})
}

func StatementDetail(rc *fasthttp.RequestCtx) {
	Act("statement.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Statement)", ret)

		return Render(rc, as, &vstatement.Detail{Model: ret}, ps, "statement", ret.TitleString()+"**database")
	})
}

func StatementCreateForm(rc *fasthttp.RequestCtx) {
	Act("statement.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &statement.Statement{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = statement.Random()
		}
		ps.SetTitleAndData("Create [Statement]", ret)
		ps.Data = ret
		return Render(rc, as, &vstatement.Edit{Model: ret, IsNew: true}, ps, "statement", "Create")
	})
}

func StatementRandom(rc *fasthttp.RequestCtx) {
	Act("statement.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Statement.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Statement")
		}
		return ret.WebPath(), nil
	})
}

func StatementCreate(rc *fasthttp.RequestCtx) {
	Act("statement.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Statement from form")
		}
		err = as.Services.Statement.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Statement")
		}
		msg := fmt.Sprintf("Statement [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StatementEditForm(rc *fasthttp.RequestCtx) {
	Act("statement.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vstatement.Edit{Model: ret}, ps, "statement", ret.String())
	})
}

func StatementEdit(rc *fasthttp.RequestCtx) {
	Act("statement.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := statementFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Statement from form")
		}
		frm.ID = ret.ID
		err = as.Services.Statement.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Statement [%s]", frm.String())
		}
		msg := fmt.Sprintf("Statement [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StatementDelete(rc *fasthttp.RequestCtx) {
	Act("statement.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Statement.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete statement [%s]", ret.String())
		}
		msg := fmt.Sprintf("Statement [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/statement", rc, ps)
	})
}

func statementFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*statement.Statement, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Statement.Get(ps.Context, nil, idArg, ps.Logger)
}

func statementFromForm(rc *fasthttp.RequestCtx, setPK bool) (*statement.Statement, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return statement.FromMap(frm, setPK)
}
