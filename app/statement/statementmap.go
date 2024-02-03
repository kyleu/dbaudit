// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import "github.com/kyleu/dbaudit/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Statement, error) {
	ret := &Statement{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.SessionID, err = m.ParseInt("sessionID", true, true)
	if err != nil {
		return nil, err
	}
	retAction, err := m.ParseString("action", true, true)
	if err != nil {
		return nil, err
	}
	ret.Action = AllActions.Get(retAction, nil)
	ret.Succeeded, err = m.ParseBool("succeeded", true, true)
	if err != nil {
		return nil, err
	}
	ret.Principal, err = m.ParseString("principal", true, true)
	if err != nil {
		return nil, err
	}
	ret.Database, err = m.ParseString("database", true, true)
	if err != nil {
		return nil, err
	}
	ret.Filename, err = m.ParseString("filename", true, true)
	if err != nil {
		return nil, err
	}
	ret.Host, err = m.ParseString("host", true, true)
	if err != nil {
		return nil, err
	}
	ret.TransactionID, err = m.ParseInt("transactionID", true, true)
	if err != nil {
		return nil, err
	}
	ret.ClientIP, err = m.ParseString("clientIP", true, true)
	if err != nil {
		return nil, err
	}
	ret.Duration, err = m.ParseInt("duration", true, true)
	if err != nil {
		return nil, err
	}
	retConnectionID, e := m.ParseUUID("connectionID", true, true)
	if e != nil {
		return nil, e
	}
	if retConnectionID != nil {
		ret.ConnectionID = *retConnectionID
	}
	ret.RowsAffected, err = m.ParseInt("rowsAffected", true, true)
	if err != nil {
		return nil, err
	}
	ret.RowsReturned, err = m.ParseInt("rowsReturned", true, true)
	if err != nil {
		return nil, err
	}
	ret.SQL, err = m.ParseString("sql", true, true)
	if err != nil {
		return nil, err
	}
	ret.Types, err = m.ParseMap("types", true, true)
	if err != nil {
		return nil, err
	}
	ret.Values, err = m.ParseMap("values", true, true)
	if err != nil {
		return nil, err
	}
	retOccurred, e := m.ParseTime("occurred", true, true)
	if e != nil {
		return nil, e
	}
	if retOccurred != nil {
		ret.Occurred = *retOccurred
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
