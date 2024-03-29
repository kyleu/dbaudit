// Package cmd - Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return _logger, err
	}
	return _logger, nil
}
