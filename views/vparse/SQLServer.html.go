// Code generated by qtc from "SQLServer.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vparse/SQLServer.html:1
package vparse

//line views/vparse/SQLServer.html:1
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/parse"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/components/edit"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vparse/SQLServer.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vparse/SQLServer.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vparse/SQLServer.html:11
type SQLServer struct {
	layout.Basic
	Path   string
	Task   string
	Limit  int
	Result *parse.Result
}

//line views/vparse/SQLServer.html:19
func (p *SQLServer) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vparse/SQLServer.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vparse/SQLServer.html:21
	components.StreamSVGIcon(qw422016, `app`, ps)
//line views/vparse/SQLServer.html:21
	qw422016.N().S(`SQL Server Analysis</h3>
    <form action="" method="post">
      <table class="mt expanded min-200">
        <tbody>
          `)
//line views/vparse/SQLServer.html:25
	edit.StreamStringTable(qw422016, "path", "", "File Path", p.Path, 5, "Location of SQL Audit files")
//line views/vparse/SQLServer.html:25
	qw422016.N().S(`
          `)
//line views/vparse/SQLServer.html:26
	edit.StreamSelectTable(qw422016, "task", "", "Task", p.Task, []string{"testbed", "count", "top10"}, nil, 5, "Task to run")
//line views/vparse/SQLServer.html:26
	qw422016.N().S(`
          `)
//line views/vparse/SQLServer.html:27
	edit.StreamIntTable(qw422016, "limit", "", "Event Limit", p.Limit, 5, "Limit for returned events")
//line views/vparse/SQLServer.html:27
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Analyze</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vparse/SQLServer.html:33
	if p.Result != nil {
//line views/vparse/SQLServer.html:33
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vparse/SQLServer.html:35
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vparse/SQLServer.html:35
		qw422016.E().S(util.StringPlural(len(p.Result.Events), "Event"))
//line views/vparse/SQLServer.html:35
		qw422016.N().S(`</h3>
    <div class="mt overflow full-width">
      <table class="expanded">
        <thead>
          <tr>
            <th>ID</th>
            <th>Action</th>
            <th>Statement</th>
            <th>Arguments</th>
            <th>Rows</th>
            <th>Duration</th>
            <th>Occurred</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vparse/SQLServer.html:50
		for _, st := range p.Result.Statements {
//line views/vparse/SQLServer.html:50
			qw422016.N().S(`          <tr>
            <td>
              <a href="#modal-event-`)
//line views/vparse/SQLServer.html:53
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:53
			qw422016.N().S(`">`)
//line views/vparse/SQLServer.html:53
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:53
			qw422016.N().S(`</a>
              <div id="modal-event-`)
//line views/vparse/SQLServer.html:54
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:54
			qw422016.N().S(`" class="modal" style="display: none;">
                <a class="backdrop" href="#"></a>
                <div class="modal-content">
                  <div class="modal-header">
                    <a href="#" class="modal-close">×</a>
                    <h2>Event [`)
//line views/vparse/SQLServer.html:59
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:59
			qw422016.N().S(`]</h2>
                  </div>
                  <div class="modal-body">
                    <ul class="accordion">
                      <li>
                        <input id="accordion-event-`)
//line views/vparse/SQLServer.html:64
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:64
			qw422016.N().S(`-raw" type="checkbox" hidden="hidden" />
                        <label for="accordion-event-`)
//line views/vparse/SQLServer.html:65
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:65
			qw422016.N().S(`-raw">`)
//line views/vparse/SQLServer.html:65
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vparse/SQLServer.html:65
			qw422016.N().S(`Full Event</label>
                        <div class="bd"><div><div>
                          `)
//line views/vparse/SQLServer.html:67
			components.StreamJSON(qw422016, st)
//line views/vparse/SQLServer.html:67
			qw422016.N().S(`
                        </div></div></div>
                      </li>
                      <li>
                        <input id="accordion-event-`)
//line views/vparse/SQLServer.html:71
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:71
			qw422016.N().S(`-sql" type="checkbox" hidden="hidden" />
                        <label for="accordion-event-`)
//line views/vparse/SQLServer.html:72
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:72
			qw422016.N().S(`-sql">`)
//line views/vparse/SQLServer.html:72
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vparse/SQLServer.html:72
			qw422016.N().S(`SQL Statement</label>
                        <div class="bd"><div><div>
                          <pre style="white-space: pre-wrap;">`)
//line views/vparse/SQLServer.html:74
			qw422016.E().S(st.SQL)
//line views/vparse/SQLServer.html:74
			qw422016.N().S(`</pre>
                        </div></div></div>
                      </li>
                      <li>
                        <input id="accordion-event-`)
//line views/vparse/SQLServer.html:78
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:78
			qw422016.N().S(`-args" type="checkbox" hidden="hidden" />
                        <label for="accordion-event-`)
//line views/vparse/SQLServer.html:79
			qw422016.E().S(st.ID.String())
//line views/vparse/SQLServer.html:79
			qw422016.N().S(`-args">`)
//line views/vparse/SQLServer.html:79
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vparse/SQLServer.html:79
			qw422016.N().S(`Arguments</label>
                        <div class="bd"><div><div>
                          `)
//line views/vparse/SQLServer.html:81
			components.StreamJSON(qw422016, st.Values)
//line views/vparse/SQLServer.html:81
			qw422016.N().S(`
                        </div></div></div>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td>`)
//line views/vparse/SQLServer.html:89
			qw422016.E().S(st.Action.String())
//line views/vparse/SQLServer.html:89
			qw422016.N().S(`</td>
            <td>`)
//line views/vparse/SQLServer.html:90
			qw422016.E().S(util.ByteSizeSI(int64(len(st.SQL))))
//line views/vparse/SQLServer.html:90
			qw422016.N().S(`</td>
            <td>`)
//line views/vparse/SQLServer.html:91
			qw422016.N().D(len(st.Values))
//line views/vparse/SQLServer.html:91
			qw422016.N().S(`</td>
            <td>`)
//line views/vparse/SQLServer.html:92
			qw422016.N().D(st.RowsAffected)
//line views/vparse/SQLServer.html:92
			qw422016.N().S(`</td>
            <td>`)
//line views/vparse/SQLServer.html:93
			qw422016.E().S(util.MicrosToMillis(st.Duration * 1000))
//line views/vparse/SQLServer.html:93
			qw422016.N().S(`</td>
            <td>`)
//line views/vparse/SQLServer.html:94
			qw422016.E().S(util.TimeToFullMS(&st.Occurred))
//line views/vparse/SQLServer.html:94
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vparse/SQLServer.html:96
		}
//line views/vparse/SQLServer.html:96
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vparse/SQLServer.html:101
	}
//line views/vparse/SQLServer.html:102
}

//line views/vparse/SQLServer.html:102
func (p *SQLServer) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vparse/SQLServer.html:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vparse/SQLServer.html:102
	p.StreamBody(qw422016, as, ps)
//line views/vparse/SQLServer.html:102
	qt422016.ReleaseWriter(qw422016)
//line views/vparse/SQLServer.html:102
}

//line views/vparse/SQLServer.html:102
func (p *SQLServer) Body(as *app.State, ps *cutil.PageState) string {
//line views/vparse/SQLServer.html:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vparse/SQLServer.html:102
	p.WriteBody(qb422016, as, ps)
//line views/vparse/SQLServer.html:102
	qs422016 := string(qb422016.B)
//line views/vparse/SQLServer.html:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vparse/SQLServer.html:102
	return qs422016
//line views/vparse/SQLServer.html:102
}
