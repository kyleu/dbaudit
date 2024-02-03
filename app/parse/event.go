package parse

import (
	"strings"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"

	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
)

type Event struct {
	EventTime      time.Time              `db:"event_time,omitempty" json:"eventTime,omitempty"`
	SessionID      int                    `db:"session_id,omitempty" json:"sessionID,omitempty"`
	ActionID       string                 `db:"action_id,omitempty" json:"actionID,omitempty"`
	SequenceGroup  mssql.UniqueIdentifier `db:"sequence_group_id" json:"sequenceGroup,omitempty"`
	SequenceNumber int                    `db:"sequence_number,omitempty" json:"sequenceNumber,omitempty"`
	Succeeded      bool                   `db:"succeeded,omitempty" json:"succeeded,omitempty"`
	ObjectID       int                    `db:"object_id,omitempty" json:"objectID,omitempty"`
	ClassType      string                 `db:"class_type,omitempty" json:"classType,omitempty"`
	PrincipalName  string                 `db:"session_server_principal_name,omitempty" json:"principalName,omitempty"`
	DatabaseName   string                 `db:"database_name,omitempty" json:"databaseName,omitempty"`
	ObjectName     string                 `db:"object_name,omitempty" json:"objectName,omitempty"`
	Statement      string                 `db:"statement,omitempty" json:"statement,omitempty"`
	AdditionalInfo string                 `db:"additional_information,omitempty" json:"additionalInfo,omitempty"`
	Filename       string                 `db:"file_name,omitempty" json:"filename,omitempty"`
	Offset         int                    `db:"audit_file_offset,omitempty" json:"offset,omitempty"`
	TransactionID  int                    `db:"transaction_id,omitempty" json:"transactionID,omitempty"`
	ClientIP       string                 `db:"client_ip,omitempty" json:"clientIP,omitempty"`
	Application    string                 `db:"application_name,omitempty" json:"application,omitempty"`
	DurationMS     int                    `db:"duration_milliseconds,omitempty" json:"durationMS,omitempty"`
	ResponseRows   int                    `db:"response_rows,omitempty" json:"responseRows,omitempty"`
	AffectedRows   int                    `db:"affected_rows,omitempty" json:"affectedRows,omitempty"`
	ConnectionID   mssql.UniqueIdentifier `db:"connection_id,omitempty" json:"connectionID,omitempty"`
	HostName       string                 `db:"host_name,omitempty" json:"hostName,omitempty"`
}

func (e *Event) ID() uuid.UUID {
	return util.UUIDFromStringOK(e.SequenceGroup.String())
}

func (e *Event) Type() string {
	if e.Statement == "select 1;" {
		return "healthcheck"
	}
	switch strings.ToLower(strings.TrimSpace(e.ActionID)) {
	case "bst":
		return "batch_start"
	case "bcm":
		return "batch_complete"
	case "rst":
		return "proc_start"
	case "rcm":
		return "proc_complete"
	case "trbs":
		return "tx_begin_start"
	case "trbc":
		return "tx_begin_complete"
	case "trcs":
		return "tx_commit_start"
	case "trcc":
		return "tx_commit_complete"
	default:
		return "unknown:" + e.ActionID
	}
}

func (e *Event) Append(x *Event) {
	e.Statement += x.Statement
}

func (e *Event) Parsed() (string, util.ValueMap, util.ValueMap, error) {
	return parseSQL(e.Statement)
}

func (e *Event) ToStatement() *statement.Statement {
	t := statement.AllActions.Get(e.Type(), nil)
	conn := util.UUIDFromStringOK(e.ConnectionID.String())
	sql, typs, vals, err := e.Parsed()
	if err != nil {
		vals = util.ValueMap{"parse_error": err.Error()}
	}
	return &statement.Statement{
		ID: e.ID(), SessionID: e.SessionID, Action: t, Succeeded: e.Succeeded, Principal: e.PrincipalName,
		Database: e.DatabaseName, Filename: e.Filename, Host: e.HostName, TransactionID: e.TransactionID,
		ClientIP: e.ClientIP, Duration: e.DurationMS, ConnectionID: conn, RowsAffected: e.AffectedRows,
		RowsReturned: e.ResponseRows, SQL: sql, Types: typs, Values: vals, Occurred: e.EventTime,
	}
}

type Events []*Event

func (e Events) Collapse() Events {
	ret := make(Events, 0, len(e))
	for _, x := range e {
		if x.SequenceNumber == 1 {
			ret = append(ret, x)
		} else {
			ret[len(ret)-1].Append(x)
		}
	}
	return ret
}

func (e Events) ToStatements() statement.Statements {
	ret := make(statement.Statements, 0, len(e))
	for _, x := range e {
		curr := x.ToStatement()
		switch x.Type() {
		default:
			ret = append(ret, curr)
		}
	}
	return ret
}
