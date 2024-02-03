// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/dbaudit/app/lib/menu"

//nolint:lll
var menuItemStatement = &menu.Item{Key: "statement", Title: "Statements", Description: "A representation of a SQL execution", Icon: "database", Route: "/statement"}

//nolint:unused
func generatedMenu() menu.Items {
	return menu.Items{
		menuItemStatement,
	}
}
