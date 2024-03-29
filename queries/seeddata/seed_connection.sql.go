// Code generated by qtc from "seed_connection.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_connection.sql:1
package seeddata

//line queries/seeddata/seed_connection.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_connection.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_connection.sql:1
func StreamConnectionSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_connection.sql:1
	qw422016.N().S(`
insert into "connection" (
  "id", "name", "icon", "engine", "server", "port", "username", "password", "database", "schema", "conn_override"
) values (
  '10000000-0000-0000-0000-000000000001', 'Local SQL Server', 'star', 'mssql', 'localhost', 1433, 'sa', 'aStrongPassword123!', 'Groupmatics', '', ''
);
-- `)
//line queries/seeddata/seed_connection.sql:7
}

//line queries/seeddata/seed_connection.sql:7
func WriteConnectionSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_connection.sql:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_connection.sql:7
	StreamConnectionSeedData(qw422016)
//line queries/seeddata/seed_connection.sql:7
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_connection.sql:7
}

//line queries/seeddata/seed_connection.sql:7
func ConnectionSeedData() string {
//line queries/seeddata/seed_connection.sql:7
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_connection.sql:7
	WriteConnectionSeedData(qb422016)
//line queries/seeddata/seed_connection.sql:7
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_connection.sql:7
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_connection.sql:7
	return qs422016
//line queries/seeddata/seed_connection.sql:7
}
