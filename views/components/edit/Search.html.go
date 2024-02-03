// Code generated by qtc from "Search.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Search.html:2
package edit

//line views/components/edit/Search.html:2
import (
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/views/components"
)

//line views/components/edit/Search.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Search.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Search.html:7
func StreamSearchForm(qw422016 *qt422016.Writer, action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) {
//line views/components/edit/Search.html:9
	if fieldKey == "" {
		fieldKey = "q"
	}
	m := map[string]string{}
	ps.URI.QueryArgs().VisitAll(func(key []byte, value []byte) {
		k, v := string(key), string(value)
		if k == fieldKey {
			return
		}
		curr, ok := m[k]
		if ok {
			curr += ","
		}
		curr += v
		m[k] = curr
	})

//line views/components/edit/Search.html:25
	qw422016.N().S(`<form action="`)
//line views/components/edit/Search.html:26
	qw422016.E().S(action)
//line views/components/edit/Search.html:26
	qw422016.N().S(`" method="get" class="">`)
//line views/components/edit/Search.html:27
	for k, v := range m {
//line views/components/edit/Search.html:27
		qw422016.N().S(`<input type="hidden" name="`)
//line views/components/edit/Search.html:28
		qw422016.E().S(k)
//line views/components/edit/Search.html:28
		qw422016.N().S(`" value="`)
//line views/components/edit/Search.html:28
		qw422016.E().S(v)
//line views/components/edit/Search.html:28
		qw422016.N().S(`" />`)
//line views/components/edit/Search.html:29
	}
//line views/components/edit/Search.html:29
	qw422016.N().S(`<button class="right" type="submit">`)
//line views/components/edit/Search.html:30
	components.StreamSVGRef(qw422016, "search", 22, 22, `icon`, ps)
//line views/components/edit/Search.html:30
	qw422016.N().S(`</button><input class="right br0" type="search" placeholder="`)
//line views/components/edit/Search.html:31
	qw422016.E().S(placeholder)
//line views/components/edit/Search.html:31
	qw422016.N().S(`" name="`)
//line views/components/edit/Search.html:31
	qw422016.E().S(fieldKey)
//line views/components/edit/Search.html:31
	qw422016.N().S(`" value="`)
//line views/components/edit/Search.html:31
	qw422016.E().S(value)
//line views/components/edit/Search.html:31
	qw422016.N().S(`"><div class="clear"></div></form>`)
//line views/components/edit/Search.html:34
}

//line views/components/edit/Search.html:34
func WriteSearchForm(qq422016 qtio422016.Writer, action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) {
//line views/components/edit/Search.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Search.html:34
	StreamSearchForm(qw422016, action, fieldKey, placeholder, value, ps)
//line views/components/edit/Search.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Search.html:34
}

//line views/components/edit/Search.html:34
func SearchForm(action string, fieldKey string, placeholder string, value string, ps *cutil.PageState) string {
//line views/components/edit/Search.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Search.html:34
	WriteSearchForm(qb422016, action, fieldKey, placeholder, value, ps)
//line views/components/edit/Search.html:34
	qs422016 := string(qb422016.B)
//line views/components/edit/Search.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Search.html:34
	return qs422016
//line views/components/edit/Search.html:34
}
