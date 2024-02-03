// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"fmt"
	"strings"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

var (
	table         = "statement"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "session_id", "action", "succeeded", "principal", "database", "filename", "host", "transaction_id", "client_ip", "duration", "connection_id", "rows_affected", "rows_returned", "sql", "types", "values", "occurred"} //nolint:lll
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID            mssql.UniqueIdentifier `db:"id" json:"id"`
	SessionID     int                    `db:"session_id" json:"session_id"`
	Action        Action                 `db:"action" json:"action"`
	Succeeded     bool                   `db:"succeeded" json:"succeeded"`
	Principal     string                 `db:"principal" json:"principal"`
	Database      string                 `db:"database" json:"database"`
	Filename      string                 `db:"filename" json:"filename"`
	Host          string                 `db:"host" json:"host"`
	TransactionID int                    `db:"transaction_id" json:"transaction_id"`
	ClientIP      string                 `db:"client_ip" json:"client_ip"`
	Duration      int                    `db:"duration" json:"duration"`
	ConnectionID  mssql.UniqueIdentifier `db:"connection_id" json:"connection_id"`
	RowsAffected  int                    `db:"rows_affected" json:"rows_affected"`
	RowsReturned  int                    `db:"rows_returned" json:"rows_returned"`
	SQL           string                 `db:"sql" json:"sql"`
	Types         string                 `db:"types" json:"types"`
	Values        string                 `db:"values" json:"values"`
	Occurred      time.Time              `db:"occurred" json:"occurred"`
}

func (r *row) ToStatement() *Statement {
	if r == nil {
		return nil
	}
	typesArg, _ := util.FromJSONMap([]byte(r.Types))
	valuesArg, _ := util.FromJSONMap([]byte(r.Values))
	return &Statement{
		ID:            util.UUIDFromStringOK(r.ID.String()),
		SessionID:     r.SessionID,
		Action:        r.Action,
		Succeeded:     r.Succeeded,
		Principal:     r.Principal,
		Database:      r.Database,
		Filename:      r.Filename,
		Host:          r.Host,
		TransactionID: r.TransactionID,
		ClientIP:      r.ClientIP,
		Duration:      r.Duration,
		ConnectionID:  util.UUIDFromStringOK(r.ConnectionID.String()),
		RowsAffected:  r.RowsAffected,
		RowsReturned:  r.RowsReturned,
		SQL:           r.SQL,
		Types:         typesArg,
		Values:        valuesArg,
		Occurred:      r.Occurred,
	}
}

type rows []*row

func (x rows) ToStatements() Statements {
	return lo.Map(x, func(d *row, _ int) *Statement {
		return d.ToStatement()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = @p%d", idx+1)
}
