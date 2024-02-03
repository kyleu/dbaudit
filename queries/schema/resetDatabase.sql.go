// Code generated by qtc from "resetDatabase.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/schema/resetDatabase.sql:1
package schema

//line queries/schema/resetDatabase.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/schema/resetDatabase.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/schema/resetDatabase.sql:1
func StreamResetDatabase(qw422016 *qt422016.Writer) {
//line queries/schema/resetDatabase.sql:1
	qw422016.N().S(`
drop table if exists "foo";
-- `)
//line queries/schema/resetDatabase.sql:3
}

//line queries/schema/resetDatabase.sql:3
func WriteResetDatabase(qq422016 qtio422016.Writer) {
//line queries/schema/resetDatabase.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/schema/resetDatabase.sql:3
	StreamResetDatabase(qw422016)
//line queries/schema/resetDatabase.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/schema/resetDatabase.sql:3
}

//line queries/schema/resetDatabase.sql:3
func ResetDatabase() string {
//line queries/schema/resetDatabase.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/schema/resetDatabase.sql:3
	WriteResetDatabase(qb422016)
//line queries/schema/resetDatabase.sql:3
	qs422016 := string(qb422016.B)
//line queries/schema/resetDatabase.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/schema/resetDatabase.sql:3
	return qs422016
//line queries/schema/resetDatabase.sql:3
}
