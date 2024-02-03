// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"fmt"
	"strings"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

var (
	table         = "connection"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "icon", "engine", "server", "port", "username", "password", "database", "schema", "conn_override"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID           mssql.UniqueIdentifier `db:"id" json:"id"`
	Name         string                 `db:"name" json:"name"`
	Icon         string                 `db:"icon" json:"icon"`
	Engine       Engine                 `db:"engine" json:"engine"`
	Server       string                 `db:"server" json:"server"`
	Port         int                    `db:"port" json:"port"`
	Username     string                 `db:"username" json:"username"`
	Password     string                 `db:"password" json:"password"`
	Database     string                 `db:"database" json:"database"`
	Schema       string                 `db:"schema" json:"schema"`
	ConnOverride string                 `db:"conn_override" json:"conn_override"`
}

func (r *row) ToConnection() *Connection {
	if r == nil {
		return nil
	}
	return &Connection{
		ID:           util.UUIDFromStringOK(r.ID.String()),
		Name:         r.Name,
		Icon:         r.Icon,
		Engine:       r.Engine,
		Server:       r.Server,
		Port:         r.Port,
		Username:     r.Username,
		Password:     r.Password,
		Database:     r.Database,
		Schema:       r.Schema,
		ConnOverride: r.ConnOverride,
	}
}

type rows []*row

func (x rows) ToConnections() Connections {
	return lo.Map(x, func(d *row, _ int) *Connection {
		return d.ToConnection()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = @p%d", idx+1)
}
