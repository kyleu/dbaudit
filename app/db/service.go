// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"github.com/google/uuid"

	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/lib/svc"
)

var (
	_ svc.ServiceID[*Connection, Connections, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Connections]                     = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["db"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("db")
}
