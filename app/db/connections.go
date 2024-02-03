// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

type Connections []*Connection

func (c Connections) Get(id uuid.UUID) *Connection {
	return lo.FindOrElse(c, nil, func(x *Connection) bool {
		return x.ID == id
	})
}

func (c Connections) IDs() []uuid.UUID {
	return lo.Map(c, func(xx *Connection, _ int) uuid.UUID {
		return xx.ID
	})
}

func (c Connections) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(c, func(x *Connection, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (c Connections) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(c, func(x *Connection, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (c Connections) GetByID(id uuid.UUID) Connections {
	return lo.Filter(c, func(xx *Connection, _ int) bool {
		return xx.ID == id
	})
}

func (c Connections) GetByIDs(ids ...uuid.UUID) Connections {
	return lo.Filter(c, func(xx *Connection, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (c Connections) Engines() []Engine {
	return lo.Map(c, func(xx *Connection, _ int) Engine {
		return xx.Engine
	})
}

func (c Connections) GetByEngine(engine Engine) Connections {
	return lo.Filter(c, func(xx *Connection, _ int) bool {
		return xx.Engine == engine
	})
}

func (c Connections) GetByEngines(engines ...Engine) Connections {
	return lo.Filter(c, func(xx *Connection, _ int) bool {
		return lo.Contains(engines, xx.Engine)
	})
}

func (c Connections) Random() *Connection {
	if len(c) == 0 {
		return nil
	}
	return c[util.RandomInt(len(c))]
}

func (c Connections) Clone() Connections {
	return slices.Clone(c)
}
