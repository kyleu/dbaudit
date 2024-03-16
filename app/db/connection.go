// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kyleu/dbaudit/app/util"
)

type Connection struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Icon         string    `json:"icon,omitempty"`
	Engine       Engine    `json:"engine,omitempty"`
	Server       string    `json:"server,omitempty"`
	Port         int       `json:"port,omitempty"`
	Username     string    `json:"username,omitempty"`
	Password     string    `json:"password,omitempty"`
	Database     string    `json:"database,omitempty"`
	Schema       string    `json:"schema,omitempty"`
	ConnOverride string    `json:"connOverride,omitempty"`
}

func New(id uuid.UUID) *Connection {
	return &Connection{ID: id}
}

func (c *Connection) Clone() *Connection {
	return &Connection{
		c.ID, c.Name, c.Icon, c.Engine, c.Server, c.Port, c.Username, c.Password, c.Database, c.Schema, c.ConnOverride,
	}
}

func (c *Connection) String() string {
	return c.ID.String()
}

func (c *Connection) TitleString() string {
	return c.Name
}

func Random() *Connection {
	return &Connection{
		ID:           util.UUID(),
		Name:         util.RandomString(12),
		Icon:         util.RandomString(12),
		Engine:       AllEngines.Random(),
		Server:       util.RandomString(12),
		Port:         util.RandomInt(10000),
		Username:     util.RandomString(12),
		Password:     util.RandomString(12),
		Database:     util.RandomString(12),
		Schema:       util.RandomString(12),
		ConnOverride: util.RandomString(12),
	}
}

func (c *Connection) Strings() []string {
	return []string{c.ID.String(), c.Name, c.Icon, c.Engine.String(), c.Server, fmt.Sprint(c.Port), c.Username, c.Password, c.Database, c.Schema, c.ConnOverride}
}

func (c *Connection) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Connection) WebPath() string {
	return "/db/" + c.ID.String()
}

func (c *Connection) ToData() []any {
	return []any{
		c.ID, c.Name, c.Icon, c.Engine, c.Server, c.Port, c.Username, c.Password, c.Database, c.Schema, c.ConnOverride,
	}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "engine", Title: "Engine", Description: "", Type: "enum(engine)"},
	{Key: "server", Title: "Server", Description: "", Type: "string"},
	{Key: "port", Title: "Port", Description: "", Type: "int"},
	{Key: "username", Title: "Username", Description: "", Type: "string"},
	{Key: "password", Title: "Password", Description: "", Type: "string"},
	{Key: "database", Title: "Database", Description: "", Type: "string"},
	{Key: "schema", Title: "Schema", Description: "", Type: "string"},
	{Key: "connOverride", Title: "Conn Override", Description: "", Type: "string"},
}
