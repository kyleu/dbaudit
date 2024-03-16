// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

type Statements []*Statement

func (s Statements) Get(id uuid.UUID) *Statement {
	return lo.FindOrElse(s, nil, func(x *Statement) bool {
		return x.ID == id
	})
}

func (s Statements) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Statement, _ int) uuid.UUID {
		return xx.ID
	})
}

func (s Statements) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Statement, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Statements) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Statement, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Statements) GetByID(id uuid.UUID) Statements {
	return lo.Filter(s, func(xx *Statement, _ int) bool {
		return xx.ID == id
	})
}

func (s Statements) GetByIDs(ids ...uuid.UUID) Statements {
	return lo.Filter(s, func(xx *Statement, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Statements) Actions() []Action {
	return lo.Map(s, func(xx *Statement, _ int) Action {
		return xx.Action
	})
}

func (s Statements) GetByAction(action Action) Statements {
	return lo.Filter(s, func(xx *Statement, _ int) bool {
		return xx.Action == action
	})
}

func (s Statements) GetByActions(actions ...Action) Statements {
	return lo.Filter(s, func(xx *Statement, _ int) bool {
		return lo.Contains(actions, xx.Action)
	})
}

func (s Statements) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(s, func(x *Statement, _ int) []string {
		return x.Strings()
	})
}

func (s Statements) Random() *Statement {
	if len(s) == 0 {
		return nil
	}
	return s[util.RandomInt(len(s))]
}

func (s Statements) Clone() Statements {
	return slices.Clone(s)
}
