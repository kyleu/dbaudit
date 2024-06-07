// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Request.html:2
package vadmin

//line views/vadmin/Request.html:2
import (
	"net/http"

	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vadmin/Request.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Request.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Request.html:14
type Request struct {
	layout.Basic
	Rsp http.ResponseWriter
	Req *http.Request
}

//line views/vadmin/Request.html:20
func (p *Request) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-ps"><button type="button">Page State</button></a></div>
    <h3>`)
//line views/vadmin/Request.html:23
	components.StreamSVGIcon(qw422016, `link`, ps)
//line views/vadmin/Request.html:23
	qw422016.N().S(`Request Debug</h3>
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
            <td>URL</td>
            <td>`)
//line views/vadmin/Request.html:35
	qw422016.E().S(p.Req.URL.String())
//line views/vadmin/Request.html:35
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Protocol</td>
            <td>`)
//line views/vadmin/Request.html:39
	qw422016.E().S(p.Req.URL.Scheme)
//line views/vadmin/Request.html:39
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Host</td>
            <td>`)
//line views/vadmin/Request.html:43
	qw422016.E().S(p.Req.URL.Host)
//line views/vadmin/Request.html:43
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Path</td>
            <td>`)
//line views/vadmin/Request.html:47
	qw422016.E().S(p.Req.URL.Path)
//line views/vadmin/Request.html:47
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Query String</td>
            <td>`)
//line views/vadmin/Request.html:51
	qw422016.E().S(p.Req.URL.RawQuery)
//line views/vadmin/Request.html:51
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Body Size</td>
            <td>`)
//line views/vadmin/Request.html:55
	qw422016.E().S(util.ByteSizeSI(int64(len(ps.RequestBody))))
//line views/vadmin/Request.html:55
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Browser</td>
            <td>`)
//line views/vadmin/Request.html:59
	qw422016.E().S(ps.Browser)
//line views/vadmin/Request.html:59
	qw422016.N().S(` `)
//line views/vadmin/Request.html:59
	qw422016.E().S(ps.BrowserVersion)
//line views/vadmin/Request.html:59
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>OS</td>
            <td>`)
//line views/vadmin/Request.html:63
	qw422016.E().S(ps.OS)
//line views/vadmin/Request.html:63
	qw422016.N().S(` `)
//line views/vadmin/Request.html:63
	qw422016.E().S(ps.OSVersion)
//line views/vadmin/Request.html:63
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:69
	if len(p.Req.Header) > 0 {
//line views/vadmin/Request.html:70
		hd := cutil.RequestHeadersMap(p.Req)

//line views/vadmin/Request.html:70
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Request.html:72
		components.StreamSVGIcon(qw422016, `code`, ps)
//line views/vadmin/Request.html:72
		qw422016.N().S(`Headers</h3>
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
//line views/vadmin/Request.html:82
		for _, k := range util.ArraySorted(lo.Keys(hd)) {
//line views/vadmin/Request.html:82
			qw422016.N().S(`          <tr>
            <td class="nowrap">`)
//line views/vadmin/Request.html:84
			qw422016.E().S(k)
//line views/vadmin/Request.html:84
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Request.html:85
			qw422016.E().S(hd.GetStringOpt(k))
//line views/vadmin/Request.html:85
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Request.html:87
		}
//line views/vadmin/Request.html:87
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:92
	}
//line views/vadmin/Request.html:92
	qw422016.N().S(`  `)
//line views/vadmin/Request.html:93
	components.StreamJSONModal(qw422016, "ps", "Page State", ps, 1)
//line views/vadmin/Request.html:93
	qw422016.N().S(`
`)
//line views/vadmin/Request.html:94
}

//line views/vadmin/Request.html:94
func (p *Request) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Request.html:94
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Request.html:94
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Request.html:94
}

//line views/vadmin/Request.html:94
func (p *Request) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Request.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Request.html:94
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Request.html:94
	qs422016 := string(qb422016.B)
//line views/vadmin/Request.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Request.html:94
	return qs422016
//line views/vadmin/Request.html:94
}
