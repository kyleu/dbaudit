// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vdb/Edit.html:2
package vdb

//line views/vdb/Edit.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/db"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/components/edit"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vdb/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdb/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdb/Edit.html:11
type Edit struct {
	layout.Basic
	Model *db.Connection
	IsNew bool
}

//line views/vdb/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdb/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vdb/Edit.html:19
	if p.IsNew {
//line views/vdb/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vdb/Edit.html:21
		components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vdb/Edit.html:21
		qw422016.N().S(` New Connection</h3>
    <form action="/db/_new" class="mt" method="post">
`)
//line views/vdb/Edit.html:23
	} else {
//line views/vdb/Edit.html:23
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vdb/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vdb/Edit.html:24
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete connection [`)
//line views/vdb/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vdb/Edit.html:24
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vdb/Edit.html:25
		components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vdb/Edit.html:25
		qw422016.N().S(` Edit Connection [`)
//line views/vdb/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vdb/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vdb/Edit.html:27
	}
//line views/vdb/Edit.html:27
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vdb/Edit.html:30
	if p.IsNew {
//line views/vdb/Edit.html:30
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vdb/Edit.html:30
	}
//line views/vdb/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:31
	edit.StreamStringTable(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vdb/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:32
	edit.StreamStringTable(qw422016, "icon", "", "Icon", p.Model.Icon, 5, "String text")
//line views/vdb/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:33
	edit.StreamSelectTable(qw422016, "engine", "", "Engine", p.Model.Engine.Key, db.AllEngines.Keys(), db.AllEngines.Strings(), 5, db.AllEngines.Help())
//line views/vdb/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:34
	edit.StreamStringTable(qw422016, "server", "", "Server", p.Model.Server, 5, "String text")
//line views/vdb/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:35
	edit.StreamIntTable(qw422016, "port", "", "Port", p.Model.Port, 5, "Integer")
//line views/vdb/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:36
	edit.StreamStringTable(qw422016, "username", "", "Username", p.Model.Username, 5, "String text")
//line views/vdb/Edit.html:36
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:37
	edit.StreamStringTable(qw422016, "password", "", "Password", p.Model.Password, 5, "String text")
//line views/vdb/Edit.html:37
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:38
	edit.StreamStringTable(qw422016, "database", "", "Database", p.Model.Database, 5, "String text")
//line views/vdb/Edit.html:38
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:39
	edit.StreamStringTable(qw422016, "schema", "", "Schema", p.Model.Schema, 5, "String text")
//line views/vdb/Edit.html:39
	qw422016.N().S(`
          `)
//line views/vdb/Edit.html:40
	edit.StreamStringTable(qw422016, "connOverride", "", "Conn Override", p.Model.ConnOverride, 5, "String text")
//line views/vdb/Edit.html:40
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vdb/Edit.html:46
}

//line views/vdb/Edit.html:46
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdb/Edit.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdb/Edit.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vdb/Edit.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vdb/Edit.html:46
}

//line views/vdb/Edit.html:46
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdb/Edit.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdb/Edit.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vdb/Edit.html:46
	qs422016 := string(qb422016.B)
//line views/vdb/Edit.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdb/Edit.html:46
	return qs422016
//line views/vdb/Edit.html:46
}
