package cmenu

import "github.com/kyleu/dbaudit/app/lib/menu"

func parseMenu() menu.Items {
	return menu.Items{{Key: "parse", Title: "Parse Audits", Description: "Parses and save [sqlaudit] files", Icon: "star", Route: "/parse"}}
}
