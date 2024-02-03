// Code generated by qtc from "File.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/File.html:2
package edit

//line views/components/edit/File.html:2
import "github.com/kyleu/dbaudit/views/components"

//line views/components/edit/File.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/File.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/File.html:4
func StreamFile(qw422016 *qt422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:5
	if id == "" {
//line views/components/edit/File.html:5
		qw422016.N().S(`<label><input type="file" name="`)
//line views/components/edit/File.html:6
		qw422016.E().S(key)
//line views/components/edit/File.html:6
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:6
		qw422016.E().S(value)
//line views/components/edit/File.html:6
		qw422016.N().S(`"`)
//line views/components/edit/File.html:6
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:6
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:6
		qw422016.E().S(label)
//line views/components/edit/File.html:6
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:7
	} else {
//line views/components/edit/File.html:7
		qw422016.N().S(`<label><input id="`)
//line views/components/edit/File.html:8
		qw422016.E().S(id)
//line views/components/edit/File.html:8
		qw422016.N().S(`" type="file" name="`)
//line views/components/edit/File.html:8
		qw422016.E().S(key)
//line views/components/edit/File.html:8
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:8
		qw422016.E().S(value)
//line views/components/edit/File.html:8
		qw422016.N().S(`"`)
//line views/components/edit/File.html:8
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:8
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:8
		qw422016.E().S(label)
//line views/components/edit/File.html:8
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:9
	}
//line views/components/edit/File.html:10
}

//line views/components/edit/File.html:10
func WriteFile(qq422016 qtio422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:10
	StreamFile(qw422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:10
}

//line views/components/edit/File.html:10
func File(key string, id string, label string, value string, placeholder ...string) string {
//line views/components/edit/File.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:10
	WriteFile(qb422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:10
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:10
	return qs422016
//line views/components/edit/File.html:10
}

//line views/components/edit/File.html:12
func StreamFileTable(qw422016 *qt422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:12
	qw422016.N().S(`<tr><th class="shrink"><label for="`)
//line views/components/edit/File.html:14
	qw422016.E().S(id)
//line views/components/edit/File.html:14
	qw422016.N().S(`">`)
//line views/components/edit/File.html:14
	qw422016.E().S(title)
//line views/components/edit/File.html:14
	qw422016.N().S(`</label></th><td>`)
//line views/components/edit/File.html:16
	StreamFile(qw422016, key, id, label, value)
//line views/components/edit/File.html:16
	qw422016.N().S(`</td></tr>`)
//line views/components/edit/File.html:19
}

//line views/components/edit/File.html:19
func WriteFileTable(qq422016 qtio422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:19
	StreamFileTable(qw422016, key, id, title, label, value)
//line views/components/edit/File.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:19
}

//line views/components/edit/File.html:19
func FileTable(key string, id string, title string, label string, value string) string {
//line views/components/edit/File.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:19
	WriteFileTable(qb422016, key, id, title, label, value)
//line views/components/edit/File.html:19
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:19
	return qs422016
//line views/components/edit/File.html:19
}

//line views/components/edit/File.html:21
func StreamFileMultiple(qw422016 *qt422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:22
	if id == "" {
//line views/components/edit/File.html:22
		qw422016.N().S(`<label><input type="file" name="`)
//line views/components/edit/File.html:23
		qw422016.E().S(key)
//line views/components/edit/File.html:23
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:23
		qw422016.E().S(value)
//line views/components/edit/File.html:23
		qw422016.N().S(`" multiple="multiple"`)
//line views/components/edit/File.html:23
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:23
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:23
		qw422016.E().S(label)
//line views/components/edit/File.html:23
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:24
	} else {
//line views/components/edit/File.html:24
		qw422016.N().S(`<label><input id="`)
//line views/components/edit/File.html:25
		qw422016.E().S(id)
//line views/components/edit/File.html:25
		qw422016.N().S(`" type="file" name="`)
//line views/components/edit/File.html:25
		qw422016.E().S(key)
//line views/components/edit/File.html:25
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:25
		qw422016.E().S(value)
//line views/components/edit/File.html:25
		qw422016.N().S(`" multiple="multiple"`)
//line views/components/edit/File.html:25
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:25
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:25
		qw422016.E().S(label)
//line views/components/edit/File.html:25
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:26
	}
//line views/components/edit/File.html:27
}

//line views/components/edit/File.html:27
func WriteFileMultiple(qq422016 qtio422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:27
	StreamFileMultiple(qw422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:27
}

//line views/components/edit/File.html:27
func FileMultiple(key string, id string, label string, value string, placeholder ...string) string {
//line views/components/edit/File.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:27
	WriteFileMultiple(qb422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:27
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:27
	return qs422016
//line views/components/edit/File.html:27
}

//line views/components/edit/File.html:29
func StreamFileMultipleTable(qw422016 *qt422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:29
	qw422016.N().S(`<tr><th class="shrink"><label for="`)
//line views/components/edit/File.html:31
	qw422016.E().S(id)
//line views/components/edit/File.html:31
	qw422016.N().S(`">`)
//line views/components/edit/File.html:31
	qw422016.E().S(title)
//line views/components/edit/File.html:31
	qw422016.N().S(`</label></th><td>`)
//line views/components/edit/File.html:33
	StreamFileMultiple(qw422016, key, id, label, value)
//line views/components/edit/File.html:33
	qw422016.N().S(`</td></tr>`)
//line views/components/edit/File.html:36
}

//line views/components/edit/File.html:36
func WriteFileMultipleTable(qq422016 qtio422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:36
	StreamFileMultipleTable(qw422016, key, id, title, label, value)
//line views/components/edit/File.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:36
}

//line views/components/edit/File.html:36
func FileMultipleTable(key string, id string, title string, label string, value string) string {
//line views/components/edit/File.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:36
	WriteFileMultipleTable(qb422016, key, id, title, label, value)
//line views/components/edit/File.html:36
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:36
	return qs422016
//line views/components/edit/File.html:36
}
