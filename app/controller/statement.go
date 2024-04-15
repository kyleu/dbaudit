// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/vstatement"
)

func StatementList(w http.ResponseWriter, r *http.Request) {
	Act("statement.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("statement", ps.Logger)
		var ret statement.Statements
		var err error
		if q == "" {
			ret, err = as.Services.Statement.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Statement.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Statements", ret)
		page := &vstatement.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "statement")
	})
}

func StatementDetail(w http.ResponseWriter, r *http.Request) {
	Act("statement.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Statement)", ret)

		return Render(r, as, &vstatement.Detail{Model: ret}, ps, "statement", ret.TitleString()+"**database")
	})
}

func StatementCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("statement.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &statement.Statement{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = statement.Random()
		}
		ps.SetTitleAndData("Create [Statement]", ret)
		ps.Data = ret
		return Render(r, as, &vstatement.Edit{Model: ret, IsNew: true}, ps, "statement", "Create")
	})
}

func StatementRandom(w http.ResponseWriter, r *http.Request) {
	Act("statement.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Statement.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Statement")
		}
		return ret.WebPath(), nil
	})
}

func StatementCreate(w http.ResponseWriter, r *http.Request) {
	Act("statement.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Statement from form")
		}
		err = as.Services.Statement.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Statement")
		}
		msg := fmt.Sprintf("Statement [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func StatementEditForm(w http.ResponseWriter, r *http.Request) {
	Act("statement.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vstatement.Edit{Model: ret}, ps, "statement", ret.String())
	})
}

func StatementEdit(w http.ResponseWriter, r *http.Request) {
	Act("statement.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := statementFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Statement from form")
		}
		frm.ID = ret.ID
		err = as.Services.Statement.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Statement [%s]", frm.String())
		}
		msg := fmt.Sprintf("Statement [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func StatementDelete(w http.ResponseWriter, r *http.Request) {
	Act("statement.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := statementFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Statement.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete statement [%s]", ret.String())
		}
		msg := fmt.Sprintf("Statement [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/statement", ps)
	})
}

func statementFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*statement.Statement, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func statementFromForm(r *http.Request, b []byte, setPK bool) (*statement.Statement, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return statement.FromMap(frm, setPK)
}
