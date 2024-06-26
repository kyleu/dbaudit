// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstatement/List.html:2
package vstatement

//line views/vstatement/List.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/filter"
	"github.com/kyleu/dbaudit/app/statement"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/components/edit"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vstatement/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstatement/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstatement/List.html:12
type List struct {
	layout.Basic
	Models      statement.Statements
	Params      filter.ParamSet
	SearchQuery string
}

//line views/vstatement/List.html:19
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstatement/List.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vstatement/List.html:21
	edit.StreamSearchForm(qw422016, "", "q", "Search Statements", p.SearchQuery, ps)
//line views/vstatement/List.html:21
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vstatement/List.html:23
	if len(p.Models) > 0 {
//line views/vstatement/List.html:23
		qw422016.N().S(`<a href="/statement/_random"><button>`)
//line views/vstatement/List.html:23
		components.StreamSVGRef(qw422016, "gift", 15, 15, "icon", ps)
//line views/vstatement/List.html:23
		qw422016.N().S(`Random</button></a>`)
//line views/vstatement/List.html:23
	}
//line views/vstatement/List.html:23
	qw422016.N().S(`      <a href="/statement/_new"><button>`)
//line views/vstatement/List.html:24
	components.StreamSVGRef(qw422016, "plus", 15, 15, "icon", ps)
//line views/vstatement/List.html:24
	qw422016.N().S(`New</button></a>
    </div>
    <h3>`)
//line views/vstatement/List.html:26
	components.StreamSVGIcon(qw422016, `database`, ps)
//line views/vstatement/List.html:26
	qw422016.E().S(ps.Title)
//line views/vstatement/List.html:26
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vstatement/List.html:28
	if p.SearchQuery != "" {
//line views/vstatement/List.html:28
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vstatement/List.html:30
		qw422016.E().S(p.SearchQuery)
//line views/vstatement/List.html:30
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vstatement/List.html:31
	}
//line views/vstatement/List.html:32
	if len(p.Models) == 0 {
//line views/vstatement/List.html:32
		qw422016.N().S(`    <div class="mt"><em>No statements available</em></div>
`)
//line views/vstatement/List.html:34
	} else {
//line views/vstatement/List.html:34
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vstatement/List.html:36
		StreamTable(qw422016, p.Models, p.Params, as, ps)
//line views/vstatement/List.html:36
		qw422016.N().S(`
    </div>
`)
//line views/vstatement/List.html:38
	}
//line views/vstatement/List.html:38
	qw422016.N().S(`  </div>
`)
//line views/vstatement/List.html:40
}

//line views/vstatement/List.html:40
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstatement/List.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstatement/List.html:40
	p.StreamBody(qw422016, as, ps)
//line views/vstatement/List.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/vstatement/List.html:40
}

//line views/vstatement/List.html:40
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstatement/List.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstatement/List.html:40
	p.WriteBody(qb422016, as, ps)
//line views/vstatement/List.html:40
	qs422016 := string(qb422016.B)
//line views/vstatement/List.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstatement/List.html:40
	return qs422016
//line views/vstatement/List.html:40
}
