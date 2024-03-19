// Code generated by qtc from "Search.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Search.html:2
package edit

//line views/components/edit/Search.html:2
import (
	"strings"

	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/views/components"
)

//line views/components/edit/Search.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Search.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Search.html:9
func StreamSearchForm(qw422016 *qt422016.Writer, action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) {
//line views/components/edit/Search.html:11
	if fieldKey == "" {
		fieldKey = "q"
	}
	m := map[string]string{}
	for k, v := range ps.URI.Query() {
		if k == fieldKey || strings.HasSuffix(k, ".x") {
			return
		}
		curr, ok := m[k]
		if ok {
			curr += ","
		}
		curr += strings.Join(v, ",")
		m[k] = curr
	}

//line views/components/edit/Search.html:26
	qw422016.N().S(`<form action="`)
//line views/components/edit/Search.html:27
	qw422016.E().S(action)
//line views/components/edit/Search.html:27
	qw422016.N().S(`" method="get" class="">`)
//line views/components/edit/Search.html:28
	for k, v := range m {
//line views/components/edit/Search.html:28
		qw422016.N().S(`<input type="hidden" name="`)
//line views/components/edit/Search.html:29
		qw422016.E().S(k)
//line views/components/edit/Search.html:29
		qw422016.N().S(`" value="`)
//line views/components/edit/Search.html:29
		qw422016.E().S(v)
//line views/components/edit/Search.html:29
		qw422016.N().S(`" />`)
//line views/components/edit/Search.html:30
	}
//line views/components/edit/Search.html:30
	qw422016.N().S(`<button class="right" type="submit">`)
//line views/components/edit/Search.html:31
	components.StreamSVGRef(qw422016, "search", 22, 22, `icon`, ps)
//line views/components/edit/Search.html:31
	qw422016.N().S(`</button><input class="right br0" type="search" placeholder="`)
//line views/components/edit/Search.html:32
	qw422016.E().S(placeholder)
//line views/components/edit/Search.html:32
	qw422016.N().S(`" name="`)
//line views/components/edit/Search.html:32
	qw422016.E().S(fieldKey)
//line views/components/edit/Search.html:32
	qw422016.N().S(`" value="`)
//line views/components/edit/Search.html:32
	qw422016.E().S(value)
//line views/components/edit/Search.html:32
	qw422016.N().S(`"><div class="clear"></div></form>`)
//line views/components/edit/Search.html:35
}

//line views/components/edit/Search.html:35
func WriteSearchForm(qq422016 qtio422016.Writer, action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) {
//line views/components/edit/Search.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Search.html:35
	StreamSearchForm(qw422016, action, fieldKey, placeholder, value, ps)
//line views/components/edit/Search.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Search.html:35
}

//line views/components/edit/Search.html:35
func SearchForm(action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) string {
//line views/components/edit/Search.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Search.html:35
	WriteSearchForm(qb422016, action, fieldKey, placeholder, value, ps)
//line views/components/edit/Search.html:35
	qs422016 := string(qb422016.B)
//line views/components/edit/Search.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Search.html:35
	return qs422016
//line views/components/edit/Search.html:35
}
