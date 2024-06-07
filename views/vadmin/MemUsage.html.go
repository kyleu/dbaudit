// Code generated by qtc from "MemUsage.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/MemUsage.html:2
package vadmin

//line views/vadmin/MemUsage.html:2
import (
	"runtime"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vadmin/MemUsage.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/MemUsage.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/MemUsage.html:11
type MemUsage struct {
	layout.Basic
	Mem *runtime.MemStats
}

//line views/vadmin/MemUsage.html:16
func (p *MemUsage) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/MemUsage.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/MemUsage.html:18
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vadmin/MemUsage.html:18
	qw422016.N().S(`Memory Usage</h3>
    <em>Better formatting is coming soon</em>
    `)
//line views/vadmin/MemUsage.html:20
	qw422016.N().S(components.JSON(p.Mem))
//line views/vadmin/MemUsage.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/MemUsage.html:22
}

//line views/vadmin/MemUsage.html:22
func (p *MemUsage) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/MemUsage.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/MemUsage.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/MemUsage.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/MemUsage.html:22
}

//line views/vadmin/MemUsage.html:22
func (p *MemUsage) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/MemUsage.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/MemUsage.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/MemUsage.html:22
	qs422016 := string(qb422016.B)
//line views/vadmin/MemUsage.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/MemUsage.html:22
	return qs422016
//line views/vadmin/MemUsage.html:22
}
