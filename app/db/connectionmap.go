// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import "github.com/kyleu/dbaudit/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Connection, error) {
	ret := &Connection{}
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
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	ret.Icon, err = m.ParseString("icon", true, true)
	if err != nil {
		return nil, err
	}
	retEngine, err := m.ParseString("engine", true, true)
	if err != nil {
		return nil, err
	}
	ret.Engine = AllEngines.Get(retEngine, nil)
	ret.Server, err = m.ParseString("server", true, true)
	if err != nil {
		return nil, err
	}
	ret.Port, err = m.ParseInt("port", true, true)
	if err != nil {
		return nil, err
	}
	ret.Username, err = m.ParseString("username", true, true)
	if err != nil {
		return nil, err
	}
	ret.Password, err = m.ParseString("password", true, true)
	if err != nil {
		return nil, err
	}
	ret.Database, err = m.ParseString("database", true, true)
	if err != nil {
		return nil, err
	}
	ret.Schema, err = m.ParseString("schema", true, true)
	if err != nil {
		return nil, err
	}
	ret.ConnOverride, err = m.ParseString("connOverride", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
