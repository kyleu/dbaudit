// Code generated by qtc from "Add.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Add.html:2
package vtheme

//line views/vtheme/Add.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/theme"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/layout"
)

//line views/vtheme/Add.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Add.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Add.html:11
type Add struct {
	layout.Basic
	Palette string
	Themes  theme.Themes
}

//line views/vtheme/Add.html:17
func (p *Add) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Add.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vtheme/Add.html:19
	components.StreamSVGRefIcon(qw422016, `eye`, ps)
//line views/vtheme/Add.html:19
	qw422016.N().S(`Choose Theme</h3>
    <form action="/theme" method="post">
      <input type="hidden" name="palette" value="`)
//line views/vtheme/Add.html:21
	qw422016.E().S(p.Palette)
//line views/vtheme/Add.html:21
	qw422016.N().S(`" />
      <div class="overflow full-width">
        <table class="mt">
          <tbody>
`)
//line views/vtheme/Add.html:25
	for _, t := range p.Themes {
//line views/vtheme/Add.html:25
		qw422016.N().S(`            <tr>
              <th><a href="/theme/palette/`)
//line views/vtheme/Add.html:27
		qw422016.E().S(p.Palette)
//line views/vtheme/Add.html:27
		qw422016.N().S(`/`)
//line views/vtheme/Add.html:27
		qw422016.E().S(t.Key)
//line views/vtheme/Add.html:27
		qw422016.N().S(`">`)
//line views/vtheme/Add.html:27
		qw422016.E().S(t.Key)
//line views/vtheme/Add.html:27
		qw422016.N().S(`</a></th>
              <th class="shrink" style="background-color: #ffffff; padding: 12px 36px;">`)
//line views/vtheme/Add.html:28
		StreamMockupColors(qw422016, util.AppName, "", t.Light, true, "app", 5, ps)
//line views/vtheme/Add.html:28
		qw422016.N().S(`</th>
              <th class="shrink" style="background-color: #121212; padding: 12px 36px;">`)
//line views/vtheme/Add.html:29
		StreamMockupColors(qw422016, util.AppName, "", t.Dark, true, "app", 5, ps)
//line views/vtheme/Add.html:29
		qw422016.N().S(`</th>
            </tr>
`)
//line views/vtheme/Add.html:31
	}
//line views/vtheme/Add.html:31
	qw422016.N().S(`          </tbody>
        </table>
      </div>
    </form>
  </div>
`)
//line views/vtheme/Add.html:37
}

//line views/vtheme/Add.html:37
func (p *Add) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Add.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Add.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/Add.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Add.html:37
}

//line views/vtheme/Add.html:37
func (p *Add) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/Add.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Add.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/Add.html:37
	qs422016 := string(qb422016.B)
//line views/vtheme/Add.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Add.html:37
	return qs422016
//line views/vtheme/Add.html:37
}
