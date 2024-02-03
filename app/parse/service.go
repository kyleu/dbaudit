package parse

import (
	"context"
	"github.com/kyleu/dbaudit/queries"

	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/app/util"
)

type Service struct {
	DB         *database.Service
	Statements *statement.Service
}

func NewService(db *database.Service, s *statement.Service) *Service {
	return &Service{DB: db, Statements: s}
}

func (s *Service) Testbed(ctx context.Context, path string, limit int, logger util.Logger) (*Result, error) {
	q := queries.AuditFiles(path, limit)
	var res Events
	err := s.DB.Select(ctx, &res, q, nil, logger, path)
	if err != nil {
		return nil, err
	}
	evts := res.Collapse()
	stmts := evts.ToStatements()
	commits := stmts.GetByAction(statement.ActionProcComplete)

	err = s.Statements.SaveChunked(ctx, nil, 100, logger, commits...)
	if err != nil {
		return nil, err
	}

	return &Result{Events: evts, Statements: commits}, nil
}
