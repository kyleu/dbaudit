package app // import github.com/kyleu/dbaudit

import (
	"context"
	"github.com/kyleu/dbaudit/app/db"

	"github.com/pkg/errors"

	"github.com/kyleu/dbaudit/app/lib/database/migrate"
	"github.com/kyleu/dbaudit/app/parse"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/queries/migrations"
)

type Services struct {
	Connection *db.Service
	Statement  *statement.Service
	Parse      *parse.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run database migrations")
	}
	stmtSvc := statement.NewService(st.DB)
	return &Services{
		Connection: db.NewService(st.DB),
		Statement:  stmtSvc,
		Parse:      parse.NewService(st.DB, stmtSvc),
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
