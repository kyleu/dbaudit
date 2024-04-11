// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"github.com/google/uuid"

	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/lib/svc"
)

var (
	_ svc.ServiceID[*Statement, Statements, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Statements]                    = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["statement"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("statement")
}
