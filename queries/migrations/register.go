package migrations

import (
	"github.com/kyleu/dbaudit/app/lib/database/migrate"
)

func LoadMigrations(debug bool) {
	migrate.RegisterMigration("create initial database", Migration1InitialDatabase(debug))
}
