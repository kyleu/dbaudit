// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"fmt"

	"github.com/kyleu/dbaudit/app/util"
)

func (c *Connection) Diff(cx *Connection) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID.String(), cx.ID.String()))
	}
	if c.Name != cx.Name {
		diffs = append(diffs, util.NewDiff("name", c.Name, cx.Name))
	}
	if c.Icon != cx.Icon {
		diffs = append(diffs, util.NewDiff("icon", c.Icon, cx.Icon))
	}
	if c.Engine != cx.Engine {
		diffs = append(diffs, util.NewDiff("engine", c.Engine.Key, cx.Engine.Key))
	}
	if c.Server != cx.Server {
		diffs = append(diffs, util.NewDiff("server", c.Server, cx.Server))
	}
	if c.Port != cx.Port {
		diffs = append(diffs, util.NewDiff("port", fmt.Sprint(c.Port), fmt.Sprint(cx.Port)))
	}
	if c.Username != cx.Username {
		diffs = append(diffs, util.NewDiff("username", c.Username, cx.Username))
	}
	if c.Password != cx.Password {
		diffs = append(diffs, util.NewDiff("password", c.Password, cx.Password))
	}
	if c.Database != cx.Database {
		diffs = append(diffs, util.NewDiff("database", c.Database, cx.Database))
	}
	if c.Schema != cx.Schema {
		diffs = append(diffs, util.NewDiff("schema", c.Schema, cx.Schema))
	}
	if c.ConnOverride != cx.ConnOverride {
		diffs = append(diffs, util.NewDiff("connOverride", c.ConnOverride, cx.ConnOverride))
	}
	return diffs
}
