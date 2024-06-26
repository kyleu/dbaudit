// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Settings.html:2
package vadmin

//line views/vadmin/Settings.html:2
import (
	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/user"
	"github.com/kyleu/dbaudit/app/util"
	"github.com/kyleu/dbaudit/views/components"
	"github.com/kyleu/dbaudit/views/layout"
	"github.com/kyleu/dbaudit/views/vauth"
)

//line views/vadmin/Settings.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:12
type Settings struct {
	layout.Basic
	Perms user.Permissions
}

//line views/vadmin/Settings.html:17
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:19
	if util.AppSource != "" {
//line views/vadmin/Settings.html:19
		qw422016.N().S(`    <div class="right"><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:20
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:20
		qw422016.N().S(`"><button>GitHub</button></a></div>
`)
//line views/vadmin/Settings.html:21
	}
//line views/vadmin/Settings.html:21
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:22
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:22
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:22
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:22
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:22
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:22
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:22
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:23
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:23
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:24
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:25
	}
//line views/vadmin/Settings.html:26
	if util.AppURL != "" {
//line views/vadmin/Settings.html:26
		qw422016.N().S(`    <p><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:27
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:27
		qw422016.N().S(`">`)
//line views/vadmin/Settings.html:27
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:27
		qw422016.N().S(`</a></p>
`)
//line views/vadmin/Settings.html:28
	}
//line views/vadmin/Settings.html:28
	qw422016.N().S(`    <em>This page is for the settings of the application. To change your user preferences, such as themes, <a href="/profile">edit your profile</a>.</em>
  </div>

  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:33
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:33
	qw422016.N().S(`Admin Functions</h3>
    `)
//line views/vadmin/Settings.html:34
	streamsettingsLink(qw422016, "/admin/server", "archive", "App Information", "All sorts of info about the server and runtime", ps)
//line views/vadmin/Settings.html:34
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:35
	streamsettingsLink(qw422016, "/admin/modules", "archive", "Go Modules", "The Go modules used by "+util.AppName, ps)
//line views/vadmin/Settings.html:35
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:36
	streamsettingsLink(qw422016, "/theme", "archive", "Edit Themes", "Configure the design themes available to end users", ps)
//line views/vadmin/Settings.html:36
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:37
	streamsettingsLink(qw422016, "/admin/logs", "archive", "Recent Logs", "Displays the 100 most recent app logs", ps)
//line views/vadmin/Settings.html:37
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:41
	components.StreamSVGIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(`HTTP Methods</h3>
    `)
//line views/vadmin/Settings.html:42
	streamsettingsLink(qw422016, "/admin/sitemap", "bolt", "Sitemap", "Displays the HTTP actions that are available, with documentation", ps)
//line views/vadmin/Settings.html:42
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:43
	streamsettingsLink(qw422016, "/admin/routes", "bolt", "HTTP routes", "Enumerates all registered HTTP routes, by method", ps)
//line views/vadmin/Settings.html:43
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:44
	streamsettingsLink(qw422016, "/admin/session", "bolt", "User Session", "View the user session, including all cookies and settings", ps)
//line views/vadmin/Settings.html:44
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:45
	streamsettingsLink(qw422016, "/admin/request", "bolt", "Debug HTTP Request", "Full debug view of an HTTP request from your browser", ps)
//line views/vadmin/Settings.html:45
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:49
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:49
	qw422016.N().S(`App Profiling</h3>
    `)
//line views/vadmin/Settings.html:50
	streamsettingsLink(qw422016, "/admin/memusage", "cog", "Memory Usage", "Detailed memory usage statistics for this application", ps)
//line views/vadmin/Settings.html:50
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:51
	streamsettingsLink(qw422016, "/admin/gc", "cog", "Collect Garbage", "Runs the Go garbage collector", ps)
//line views/vadmin/Settings.html:51
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:52
	streamsettingsLink(qw422016, "/admin/heap", "cog", "Write Memory Dump", "Writes a memory dump to <em>./tmp/mem.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:52
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:53
	streamsettingsLink(qw422016, "/admin/cpu/start", "cog", "Start CPU Profile", "Profiles the CPU using <em>./tmp/cpu.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:53
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:54
	streamsettingsLink(qw422016, "/admin/cpu/stop", "cog", "Stop CPU Profile", "Stops the active CPU profile", ps)
//line views/vadmin/Settings.html:54
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:58
	components.StreamSVGIcon(qw422016, `database`, ps)
//line views/vadmin/Settings.html:58
	qw422016.N().S(`Database Management</h3>
    `)
//line views/vadmin/Settings.html:59
	streamsettingsLink(qw422016, "/admin/database", "database", "Database Management", "Tools for exploring and manipulating your database", ps)
//line views/vadmin/Settings.html:59
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:60
	streamsettingsLink(qw422016, "/admin/migrations", "archive", "Migrations", "Shows the full content of all database SQL migrations", ps)
//line views/vadmin/Settings.html:60
	qw422016.N().S(`
  </div>

  `)
//line views/vadmin/Settings.html:63
	vauth.StreamAuthentication(qw422016, as, ps)
//line views/vadmin/Settings.html:63
	qw422016.N().S(`

  `)
//line views/vadmin/Settings.html:65
	vauth.StreamPermissions(qw422016, p.Perms, as)
//line views/vadmin/Settings.html:65
	qw422016.N().S(`
`)
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:66
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:66
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:66
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:66
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:66
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:66
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:66
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:66
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:66
	return qs422016
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:68
func streamsettingsLink(qw422016 *qt422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:68
	qw422016.N().S(`<hr class="clear" /><div class="mts"><a href="`)
//line views/vadmin/Settings.html:71
	qw422016.E().S(href)
//line views/vadmin/Settings.html:71
	qw422016.N().S(`"><strong>`)
//line views/vadmin/Settings.html:71
	qw422016.E().S(title)
//line views/vadmin/Settings.html:71
	qw422016.N().S(`</strong></a><div><em>`)
//line views/vadmin/Settings.html:72
	qw422016.N().S(description)
//line views/vadmin/Settings.html:72
	qw422016.N().S(`</em></div></div>`)
//line views/vadmin/Settings.html:74
}

//line views/vadmin/Settings.html:74
func writesettingsLink(qq422016 qtio422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:74
	streamsettingsLink(qw422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:74
}

//line views/vadmin/Settings.html:74
func settingsLink(href string, icon string, title string, description string, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:74
	writesettingsLink(qb422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:74
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:74
	return qs422016
//line views/vadmin/Settings.html:74
}
