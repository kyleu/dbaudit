// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kyleu/dbaudit/app/lib/auth"
	"github.com/kyleu/dbaudit/app/lib/database"
	"github.com/kyleu/dbaudit/app/lib/filesystem"
	"github.com/kyleu/dbaudit/app/lib/telemetry"
	"github.com/kyleu/dbaudit/app/lib/theme"
	"github.com/kyleu/dbaudit/app/util"
)

var once sync.Once

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == util.KeyUnknown {
		return b.Version
	}
	d, _ := util.TimeFromJS(b.Date)
	return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Auth      *auth.Service
	DB        *database.Service
	Themes    *theme.Service
	Services  *Services
	Started   time.Time
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, port uint16, logger util.Logger) (*State, error) {
	var loadLocationError error
	once.Do(func() {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			loadLocationError = err
			return
		}
		time.Local = loc
	})
	if loadLocationError != nil {
		return nil, loadLocationError
	}

	_ = telemetry.InitializeIfNeeded(enableTelemetry, bi.Version, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      auth.NewService("", port, logger),
		Themes:    theme.NewService(f),
		Started:   util.TimeCurrent(),
	}, nil
}

func (s State) Close(ctx context.Context, logger util.Logger) error {
	defer func() { _ = telemetry.Close(ctx) }()
	if err := s.DB.Close(); err != nil {
		logger.Errorf("error closing database: %+v", err)
	}
	return s.Services.Close(ctx, logger)
}
