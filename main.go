// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/dbaudit

import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/cmd"
)

var (
	version = "0.0.1" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
