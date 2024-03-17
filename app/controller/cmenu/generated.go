// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/dbaudit/app/lib/menu"

//nolint:lll
var (
	menuItemConnection = &menu.Item{Key: "db", Title: "Connections", Description: "A connection to a database", Icon: "database", Route: "/db"}
	menuItemStatement  = &menu.Item{Key: "statement", Title: "Statements", Description: "A representation of a SQL execution", Icon: "database", Route: "/statement"}
)

func generatedMenu() menu.Items {
	return menu.Items{
		menuItemConnection,
		menuItemStatement,
	}
}
