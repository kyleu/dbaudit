// Code generated by qtc from "About.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/About.html:2
package views

//line views/About.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/components/view"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/About.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/About.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/About.html:11
type About struct{ layout.Basic }

//line views/About.html:13
func (p *About) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/About.html:15
	components.StreamSVGIcon(qw422016, `app`, ps)
//line views/About.html:15
	qw422016.E().S(util.AppName)
//line views/About.html:15
	qw422016.N().S(`</h3>
    <em>v`)
//line views/About.html:16
	qw422016.E().S(as.BuildInfo.Version)
//line views/About.html:16
	qw422016.N().S(`, started `)
//line views/About.html:16
	view.StreamTimestampRelative(qw422016, &as.Started, false)
//line views/About.html:16
	qw422016.N().S(`</em>
  </div>
  <div class="card">
    <h3>About</h3>
`)
//line views/About.html:20
	qw422016.N().S(`    <p>Coming soon...</p>
`)
//line views/About.html:22
	qw422016.N().S(`  </div>
  `)
//line views/About.html:24
	StreamSourceCode(qw422016)
//line views/About.html:24
	qw422016.N().S(`
  `)
//line views/About.html:25
	StreamFeedback(qw422016)
//line views/About.html:25
	qw422016.N().S(`
`)
//line views/About.html:26
}

//line views/About.html:26
func (p *About) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:26
	p.StreamBody(qw422016, as, ps)
//line views/About.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:26
}

//line views/About.html:26
func (p *About) Body(as *app.State, ps *cutil.PageState) string {
//line views/About.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:26
	p.WriteBody(qb422016, as, ps)
//line views/About.html:26
	qs422016 := string(qb422016.B)
//line views/About.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:26
	return qs422016
//line views/About.html:26
}

//line views/About.html:28
func StreamSourceCode(qw422016 *qt422016.Writer) {
//line views/About.html:28
	qw422016.N().S(`
  <div class="card">
    <h3>Source Code</h3>
    <p>The project is available on <a href="https://github.com/kyleu/dbaudit" target="_blank" rel="noopener noreferrer">GitHub</a></p>
  </div>
`)
//line views/About.html:33
}

//line views/About.html:33
func WriteSourceCode(qq422016 qtio422016.Writer) {
//line views/About.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:33
	StreamSourceCode(qw422016)
//line views/About.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:33
}

//line views/About.html:33
func SourceCode() string {
//line views/About.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:33
	WriteSourceCode(qb422016)
//line views/About.html:33
	qs422016 := string(qb422016.B)
//line views/About.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:33
	return qs422016
//line views/About.html:33
}

//line views/About.html:35
func StreamFeedback(qw422016 *qt422016.Writer) {
//line views/About.html:35
	qw422016.N().S(`
  <div class="card">
    <h3>Feedback</h3>
    <p>For now, email <a href="mailto:kyle@kyleu.com">Kyle U</a></p>
  </div>
`)
//line views/About.html:40
}

//line views/About.html:40
func WriteFeedback(qq422016 qtio422016.Writer) {
//line views/About.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:40
	StreamFeedback(qw422016)
//line views/About.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:40
}

//line views/About.html:40
func Feedback() string {
//line views/About.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:40
	WriteFeedback(qb422016)
//line views/About.html:40
	qs422016 := string(qb422016.B)
//line views/About.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:40
	return qs422016
//line views/About.html:40
}
