// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"

	"github.com/kyleu/dbaudit/app/db"
	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
)

type GeneratedServices struct {
	Connection *db.Service
	Statement  *statement.Service
}

func initGeneratedServices(ctx context.Context, dbSvc *database.Service, logger util.Logger) GeneratedServices {
	return GeneratedServices{
		Connection: db.NewService(dbSvc),
		Statement:  statement.NewService(dbSvc),
	}
}
