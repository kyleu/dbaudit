// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"fmt"

	"github.com/kyleu/dbaudit/app/util"
)

func (s *Statement) Diff(sx *Statement) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.SessionID != sx.SessionID {
		diffs = append(diffs, util.NewDiff("sessionID", fmt.Sprint(s.SessionID), fmt.Sprint(sx.SessionID)))
	}
	if s.Action != sx.Action {
		diffs = append(diffs, util.NewDiff("action", s.Action.Key, sx.Action.Key))
	}
	if s.Succeeded != sx.Succeeded {
		diffs = append(diffs, util.NewDiff("succeeded", fmt.Sprint(s.Succeeded), fmt.Sprint(sx.Succeeded)))
	}
	if s.Principal != sx.Principal {
		diffs = append(diffs, util.NewDiff("principal", s.Principal, sx.Principal))
	}
	if s.Database != sx.Database {
		diffs = append(diffs, util.NewDiff("database", s.Database, sx.Database))
	}
	if s.Filename != sx.Filename {
		diffs = append(diffs, util.NewDiff("filename", s.Filename, sx.Filename))
	}
	if s.Host != sx.Host {
		diffs = append(diffs, util.NewDiff("host", s.Host, sx.Host))
	}
	if s.TransactionID != sx.TransactionID {
		diffs = append(diffs, util.NewDiff("transactionID", fmt.Sprint(s.TransactionID), fmt.Sprint(sx.TransactionID)))
	}
	if s.ClientIP != sx.ClientIP {
		diffs = append(diffs, util.NewDiff("clientIP", s.ClientIP, sx.ClientIP))
	}
	if s.Duration != sx.Duration {
		diffs = append(diffs, util.NewDiff("duration", fmt.Sprint(s.Duration), fmt.Sprint(sx.Duration)))
	}
	if s.ConnectionID != sx.ConnectionID {
		diffs = append(diffs, util.NewDiff("connectionID", s.ConnectionID.String(), sx.ConnectionID.String()))
	}
	if s.RowsAffected != sx.RowsAffected {
		diffs = append(diffs, util.NewDiff("rowsAffected", fmt.Sprint(s.RowsAffected), fmt.Sprint(sx.RowsAffected)))
	}
	if s.RowsReturned != sx.RowsReturned {
		diffs = append(diffs, util.NewDiff("rowsReturned", fmt.Sprint(s.RowsReturned), fmt.Sprint(sx.RowsReturned)))
	}
	if s.SQL != sx.SQL {
		diffs = append(diffs, util.NewDiff("sql", s.SQL, sx.SQL))
	}
	diffs = append(diffs, util.DiffObjects(s.Types, sx.Types, "types")...)
	diffs = append(diffs, util.DiffObjects(s.Values, sx.Values, "values")...)
	if s.Occurred != sx.Occurred {
		diffs = append(diffs, util.NewDiff("occurred", s.Occurred.String(), sx.Occurred.String()))
	}
	return diffs
}
