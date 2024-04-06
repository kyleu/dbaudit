// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstatement/Table.html:2
package vstatement

//line views/vstatement/Table.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/components/view"
)

//line views/vstatement/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstatement/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstatement/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models statement.Statements, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstatement/Table.html:11
	qw422016.N().S(`
`)
//line views/vstatement/Table.html:12
	prms := params.Sanitized("statement", ps.Logger)

//line views/vstatement/Table.html:12
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vstatement/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "statement", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstatement/Table.html:17
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "statement", "session_id", "Session ID", "Integer", prms, ps.URI, ps)
//line views/vstatement/Table.html:18
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "statement", "action", "Action", statement.AllActions.Help(), prms, ps.URI, ps)
//line views/vstatement/Table.html:19
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "statement", "succeeded", "Succeeded", "Value [true] or [false]", prms, ps.URI, ps)
//line views/vstatement/Table.html:20
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "statement", "principal", "Principal", "String text", prms, ps.URI, ps)
//line views/vstatement/Table.html:21
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "statement", "database", "Database", "String text", prms, ps.URI, ps)
//line views/vstatement/Table.html:22
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "statement", "transaction_id", "Transaction ID", "Integer", prms, ps.URI, ps)
//line views/vstatement/Table.html:23
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "statement", "duration", "Duration", "Integer", prms, ps.URI, ps)
//line views/vstatement/Table.html:24
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "statement", "rows_affected", "Rows Affected", "Integer", prms, ps.URI, ps)
//line views/vstatement/Table.html:25
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "statement", "rows_returned", "Rows Returned", "Integer", prms, ps.URI, ps)
//line views/vstatement/Table.html:26
	qw422016.N().S(`
          `)
//line views/vstatement/Table.html:27
	components.StreamTableHeaderSimple(qw422016, "statement", "occurred", "Occurred", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstatement/Table.html:27
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vstatement/Table.html:31
	for _, model := range models {
//line views/vstatement/Table.html:31
		qw422016.N().S(`        <tr>
          <td><a href="/statement/`)
//line views/vstatement/Table.html:33
		view.StreamUUID(qw422016, &model.ID)
//line views/vstatement/Table.html:33
		qw422016.N().S(`">`)
//line views/vstatement/Table.html:33
		view.StreamUUID(qw422016, &model.ID)
//line views/vstatement/Table.html:33
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vstatement/Table.html:34
		qw422016.N().D(model.SessionID)
//line views/vstatement/Table.html:34
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:35
		qw422016.E().S(model.Action.String())
//line views/vstatement/Table.html:35
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:36
		qw422016.E().V(model.Succeeded)
//line views/vstatement/Table.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:37
		view.StreamString(qw422016, model.Principal)
//line views/vstatement/Table.html:37
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:38
		view.StreamString(qw422016, model.Database)
//line views/vstatement/Table.html:38
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:39
		qw422016.N().D(model.TransactionID)
//line views/vstatement/Table.html:39
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:40
		qw422016.N().D(model.Duration)
//line views/vstatement/Table.html:40
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:41
		qw422016.N().D(model.RowsAffected)
//line views/vstatement/Table.html:41
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:42
		qw422016.N().D(model.RowsReturned)
//line views/vstatement/Table.html:42
		qw422016.N().S(`</td>
          <td>`)
//line views/vstatement/Table.html:43
		view.StreamTimestamp(qw422016, &model.Occurred)
//line views/vstatement/Table.html:43
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vstatement/Table.html:45
	}
//line views/vstatement/Table.html:45
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vstatement/Table.html:49
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstatement/Table.html:49
		qw422016.N().S(`  <hr />
  `)
//line views/vstatement/Table.html:51
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstatement/Table.html:51
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vstatement/Table.html:53
	}
//line views/vstatement/Table.html:54
}

//line views/vstatement/Table.html:54
func WriteTable(qq422016 qtio422016.Writer, models statement.Statements, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstatement/Table.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstatement/Table.html:54
	StreamTable(qw422016, models, params, as, ps)
//line views/vstatement/Table.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vstatement/Table.html:54
}

//line views/vstatement/Table.html:54
func Table(models statement.Statements, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vstatement/Table.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstatement/Table.html:54
	WriteTable(qb422016, models, params, as, ps)
//line views/vstatement/Table.html:54
	qs422016 := string(qb422016.B)
//line views/vstatement/Table.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstatement/Table.html:54
	return qs422016
//line views/vstatement/Table.html:54
}
