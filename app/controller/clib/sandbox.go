// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/sandbox"
	"github.com/kyleu/dbaudit/app/lib/telemetry"
	"github.com/kyleu/dbaudit/views/vsandbox"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	controller.Act("sandbox.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("Sandboxes", sandbox.AllSandboxes)
		return controller.Render(r, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	controller.Act("sandbox.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}

		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return controller.ERsp("no sandbox with key [%s]", key)
		}

		ctx, span, logger := telemetry.StartSpan(ps.Context, "sandbox."+key, ps.Logger)
		defer span.Complete()

		ret, err := sb.Run(ctx, as, logger.With("sandbox", key))
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(sb.Title, ret)
		if sb.Key == "testbed" {
			return controller.Render(r, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		return controller.Render(r, as, &vsandbox.Run{Key: key, Title: sb.Title, Icon: sb.Icon, Result: ret}, ps, "sandbox", sb.Key)
	})
}
