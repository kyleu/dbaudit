// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Request.html:2
package vadmin

//line views/vadmin/Request.html:2
import (
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vadmin/Request.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Request.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Request.html:13
type Request struct {
	layout.Basic
	RC *fasthttp.RequestCtx
}

//line views/vadmin/Request.html:18
func (p *Request) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-ps"><button type="button">Page State</button></a></div>
    <h3>Request Debug</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>ID</td>
            <td>`)
//line views/vadmin/Request.html:33
	qw422016.N().DUL(p.RC.ID())
//line views/vadmin/Request.html:33
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>URL</td>
            <td>`)
//line views/vadmin/Request.html:37
	qw422016.E().S(p.RC.URI().String())
//line views/vadmin/Request.html:37
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Protocol</td>
            <td>`)
//line views/vadmin/Request.html:41
	qw422016.E().S(string(p.RC.Request.URI().Scheme()))
//line views/vadmin/Request.html:41
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Host</td>
            <td>`)
//line views/vadmin/Request.html:45
	qw422016.E().S(string(p.RC.Request.URI().Host()))
//line views/vadmin/Request.html:45
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Path</td>
            <td>`)
//line views/vadmin/Request.html:49
	qw422016.E().S(string(p.RC.Request.URI().Path()))
//line views/vadmin/Request.html:49
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Query String</td>
            <td>`)
//line views/vadmin/Request.html:53
	qw422016.E().S(string(p.RC.Request.URI().QueryString()))
//line views/vadmin/Request.html:53
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Body Size</td>
            <td>`)
//line views/vadmin/Request.html:57
	qw422016.N().D(len(p.RC.Request.Body()))
//line views/vadmin/Request.html:57
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Browser</td>
            <td>`)
//line views/vadmin/Request.html:61
	qw422016.E().S(ps.Browser)
//line views/vadmin/Request.html:61
	qw422016.N().S(` `)
//line views/vadmin/Request.html:61
	qw422016.E().S(ps.BrowserVersion)
//line views/vadmin/Request.html:61
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>OS</td>
            <td>`)
//line views/vadmin/Request.html:65
	qw422016.E().S(ps.OS)
//line views/vadmin/Request.html:65
	qw422016.N().S(` `)
//line views/vadmin/Request.html:65
	qw422016.E().S(ps.OSVersion)
//line views/vadmin/Request.html:65
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:71
	if p.RC.Request.Header.Len() > 0 {
//line views/vadmin/Request.html:73
		hd := cutil.RequestHeadersMap(p.RC)

//line views/vadmin/Request.html:74
		qw422016.N().S(`  <div class="card">
    <h3>Headers</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Request.html:86
		for _, k := range util.ArraySorted(lo.Keys(hd)) {
//line views/vadmin/Request.html:86
			qw422016.N().S(`          <tr>
            <td class="nowrap">`)
//line views/vadmin/Request.html:88
			qw422016.E().S(k)
//line views/vadmin/Request.html:88
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Request.html:89
			qw422016.E().S(hd.GetStringOpt(k))
//line views/vadmin/Request.html:89
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Request.html:91
		}
//line views/vadmin/Request.html:91
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:96
	}
//line views/vadmin/Request.html:96
	qw422016.N().S(`  `)
//line views/vadmin/Request.html:97
	components.StreamJSONModal(qw422016, "ps", "Page State", ps, 1)
//line views/vadmin/Request.html:97
	qw422016.N().S(`
`)
//line views/vadmin/Request.html:98
}

//line views/vadmin/Request.html:98
func (p *Request) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Request.html:98
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Request.html:98
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Request.html:98
}

//line views/vadmin/Request.html:98
func (p *Request) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Request.html:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Request.html:98
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Request.html:98
	qs422016 := string(qb422016.B)
//line views/vadmin/Request.html:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Request.html:98
	return qs422016
//line views/vadmin/Request.html:98
}
