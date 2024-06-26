// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views"
)

func About(w http.ResponseWriter, r *http.Request) {
	controller.Act("about", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("About "+util.AppName, util.AppName+" v"+as.BuildInfo.Version)
		return controller.Render(r, as, &views.About{}, ps, "about")
	})
}
