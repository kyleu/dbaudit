// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/dbaudit/app/util"
)

type Statement struct {
	ID            uuid.UUID     `json:"id,omitempty"`
	SessionID     int           `json:"sessionID,omitempty"`
	Action        Action        `json:"action,omitempty"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	Principal     string        `json:"principal,omitempty"`
	Database      string        `json:"database,omitempty"`
	Filename      string        `json:"filename,omitempty"`
	Host          string        `json:"host,omitempty"`
	TransactionID int           `json:"transactionID,omitempty"`
	ClientIP      string        `json:"clientIP,omitempty"`
	Duration      int           `json:"duration,omitempty"`
	ConnectionID  uuid.UUID     `json:"connectionID,omitempty"`
	RowsAffected  int           `json:"rowsAffected,omitempty"`
	RowsReturned  int           `json:"rowsReturned,omitempty"`
	SQL           string        `json:"sql,omitempty"`
	Types         util.ValueMap `json:"types,omitempty"`
	Values        util.ValueMap `json:"values,omitempty"`
	Occurred      time.Time     `json:"occurred,omitempty"`
}

func New(id uuid.UUID) *Statement {
	return &Statement{ID: id}
}

func (s *Statement) Clone() *Statement {
	return &Statement{
		s.ID, s.SessionID, s.Action, s.Succeeded, s.Principal, s.Database, s.Filename, s.Host, s.TransactionID, s.ClientIP,
		s.Duration, s.ConnectionID, s.RowsAffected, s.RowsReturned, s.SQL, s.Types.Clone(), s.Values.Clone(), s.Occurred,
	}
}

func (s *Statement) String() string {
	return s.ID.String()
}

func (s *Statement) TitleString() string {
	return s.String()
}

func Random() *Statement {
	return &Statement{
		ID:            util.UUID(),
		SessionID:     util.RandomInt(10000),
		Action:        AllActions.Random(),
		Succeeded:     util.RandomBool(),
		Principal:     util.RandomString(12),
		Database:      util.RandomString(12),
		Filename:      util.RandomString(12),
		Host:          util.RandomString(12),
		TransactionID: util.RandomInt(10000),
		ClientIP:      util.RandomString(12),
		Duration:      util.RandomInt(10000),
		ConnectionID:  util.UUID(),
		RowsAffected:  util.RandomInt(10000),
		RowsReturned:  util.RandomInt(10000),
		SQL:           util.RandomString(12),
		Types:         util.RandomValueMap(4),
		Values:        util.RandomValueMap(4),
		Occurred:      util.TimeCurrent(),
	}
}

func (s *Statement) WebPath() string {
	return "/statement/" + s.ID.String()
}

func (s *Statement) ToData() []any {
	return []any{
		s.ID, s.SessionID, s.Action, s.Succeeded, s.Principal, s.Database, s.Filename, s.Host, s.TransactionID, s.ClientIP,
		s.Duration, s.ConnectionID, s.RowsAffected, s.RowsReturned, s.SQL, util.ToJSON(s.Types), util.ToJSON(s.Values),
		s.Occurred,
	}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "sessionID", Title: "Session ID", Description: "", Type: "int"},
	{Key: "action", Title: "Action", Description: "", Type: "enum(action)"},
	{Key: "succeeded", Title: "Succeeded", Description: "", Type: "bool"},
	{Key: "principal", Title: "Principal", Description: "", Type: "string"},
	{Key: "database", Title: "Database", Description: "", Type: "string"},
	{Key: "filename", Title: "Filename", Description: "", Type: "string"},
	{Key: "host", Title: "Host", Description: "", Type: "string"},
	{Key: "transactionID", Title: "Transaction ID", Description: "", Type: "int"},
	{Key: "clientIP", Title: "Client IP", Description: "", Type: "string"},
	{Key: "duration", Title: "Duration", Description: "", Type: "int"},
	{Key: "connectionID", Title: "Connection ID", Description: "", Type: "uuid"},
	{Key: "rowsAffected", Title: "Rows Affected", Description: "", Type: "int"},
	{Key: "rowsReturned", Title: "Rows Returned", Description: "", Type: "int"},
	{Key: "sql", Title: "SQL", Description: "", Type: "string"},
	{Key: "types", Title: "Types", Description: "", Type: "map"},
	{Key: "values", Title: "Values", Description: "", Type: "map"},
	{Key: "occurred", Title: "Occurred", Description: "", Type: "timestamp"},
}
