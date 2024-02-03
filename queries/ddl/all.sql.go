// Code generated by qtc from "all.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/all.sql:2
package ddl

//line queries/ddl/all.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/all.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/all.sql:2
func StreamDropAll(qw422016 *qt422016.Writer) {
//line queries/ddl/all.sql:2
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:3
	StreamStatementDrop(qw422016)
//line queries/ddl/all.sql:3
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:4
}

//line queries/ddl/all.sql:4
func WriteDropAll(qq422016 qtio422016.Writer) {
//line queries/ddl/all.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/all.sql:4
	StreamDropAll(qw422016)
//line queries/ddl/all.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/all.sql:4
}

//line queries/ddl/all.sql:4
func DropAll() string {
//line queries/ddl/all.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/all.sql:4
	WriteDropAll(qb422016)
//line queries/ddl/all.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/all.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/all.sql:4
	return qs422016
//line queries/ddl/all.sql:4
}

// --

//line queries/ddl/all.sql:6
func StreamCreateAll(qw422016 *qt422016.Writer) {
//line queries/ddl/all.sql:6
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:7
	StreamTypesCreate(qw422016)
//line queries/ddl/all.sql:7
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:8
	StreamStatementCreate(qw422016)
//line queries/ddl/all.sql:8
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:9
}

//line queries/ddl/all.sql:9
func WriteCreateAll(qq422016 qtio422016.Writer) {
//line queries/ddl/all.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/all.sql:9
	StreamCreateAll(qw422016)
//line queries/ddl/all.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/all.sql:9
}

//line queries/ddl/all.sql:9
func CreateAll() string {
//line queries/ddl/all.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/all.sql:9
	WriteCreateAll(qb422016)
//line queries/ddl/all.sql:9
	qs422016 := string(qb422016.B)
//line queries/ddl/all.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/all.sql:9
	return qs422016
//line queries/ddl/all.sql:9
}
