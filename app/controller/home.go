package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views"
)

var homeContent = util.ValueMap{
	util.AppKey: util.AppName,
	"urls": map[string]string{
		"home": "/",
	},
}

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = homeContent
		return Render(rc, as, &views.Home{}, ps)
	})
}
