// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/lib/menu"
	"github.com/kyleu/dbaudit/app/lib/sandbox"
	"github.com/kyleu/dbaudit/app/lib/user"
	"github.com/kyleu/dbaudit/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger, //nolint:revive
) (menu.Items, any, error) {
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	if isAdmin {
		ret = append(ret, parseMenu()...)
		ret = append(ret, generatedMenu()...)
		ret = append(ret, menu.Separator)
	}
	// This is your menu, feel free to customize it
	admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
	ret = append(ret, sandbox.Menu(ctx), menu.Separator, admin, menu.Separator, docMenu(logger))
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
