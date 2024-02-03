// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/Detail.html:2
package vfile

//line views/vfile/Detail.html:2
import (
	"encoding/base64"
	"path/filepath"
	"unicode/utf8"

	"github.com/kyleu/dbaudit/app"
	"github.com/kyleu/dbaudit/app/controller/cutil"
	"github.com/kyleu/dbaudit/app/lib/filesystem"
	"github.com/kyleu/dbaudit/app/util"
)

//line views/vfile/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Detail.html:13
func StreamDetail(qw422016 *qt422016.Writer, path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:13
	qw422016.N().S(`
  <div class="right"></div>
  <h3>
`)
//line views/vfile/Detail.html:16
	if additionalLinks != nil && len(additionalLinks) > 0 {
//line views/vfile/Detail.html:16
		qw422016.N().S(`    <div class="right">
`)
//line views/vfile/Detail.html:18
		for k, v := range additionalLinks {
//line views/vfile/Detail.html:18
			qw422016.N().S(`      <a href="`)
//line views/vfile/Detail.html:19
			qw422016.E().S(v)
//line views/vfile/Detail.html:19
			qw422016.N().S(`"><button>`)
//line views/vfile/Detail.html:19
			qw422016.E().S(k)
//line views/vfile/Detail.html:19
			qw422016.N().S(`</button></a>
`)
//line views/vfile/Detail.html:20
		}
//line views/vfile/Detail.html:20
		qw422016.N().S(`    </div>
`)
//line views/vfile/Detail.html:22
	}
//line views/vfile/Detail.html:23
	for idx, p := range path {
//line views/vfile/Detail.html:23
		qw422016.N().S(`/<a href="`)
//line views/vfile/Detail.html:23
		qw422016.E().S(urlPrefix)
//line views/vfile/Detail.html:23
		qw422016.N().S(`/`)
//line views/vfile/Detail.html:23
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/Detail.html:23
		qw422016.N().S(`">`)
//line views/vfile/Detail.html:23
		qw422016.E().S(p)
//line views/vfile/Detail.html:23
		qw422016.N().S(`</a>`)
//line views/vfile/Detail.html:23
	}
//line views/vfile/Detail.html:23
	qw422016.N().S(`    <em>(`)
//line views/vfile/Detail.html:24
	qw422016.E().S(util.ByteSizeSI(int64(len(b))))
//line views/vfile/Detail.html:24
	qw422016.N().S(`)</em>
  </h3>
  <div class="mt">
`)
//line views/vfile/Detail.html:27
	if len(b) > (1024 * 128) {
//line views/vfile/Detail.html:27
		qw422016.N().S(`    <em>File is `)
//line views/vfile/Detail.html:28
		qw422016.N().D(len(b))
//line views/vfile/Detail.html:28
		qw422016.N().S(` bytes, which is too large for the file viewer</em>
`)
//line views/vfile/Detail.html:29
	} else if utf8.Valid(b) {
//line views/vfile/Detail.html:30
		out, _ := cutil.FormatFilename(string(b), path[len(path)-1])

//line views/vfile/Detail.html:30
		qw422016.N().S(`    `)
//line views/vfile/Detail.html:31
		qw422016.N().S(out)
//line views/vfile/Detail.html:31
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:32
	} else {
//line views/vfile/Detail.html:32
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:34
		if imgType := filesystem.ImageType(path...); imgType != "" {
//line views/vfile/Detail.html:34
			qw422016.N().S(`    <img alt="Image of type [`)
//line views/vfile/Detail.html:35
			qw422016.E().S(imgType)
//line views/vfile/Detail.html:35
			qw422016.N().S(`]" src="data:image/`)
//line views/vfile/Detail.html:35
			qw422016.E().S(imgType)
//line views/vfile/Detail.html:35
			qw422016.N().S(`;base64,`)
//line views/vfile/Detail.html:35
			qw422016.E().S(base64.StdEncoding.EncodeToString(b))
//line views/vfile/Detail.html:35
			qw422016.N().S(`" />
    <hr />
`)
//line views/vfile/Detail.html:37
		}
//line views/vfile/Detail.html:37
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:39
		exif, err := filesystem.ExifExtract(b)

//line views/vfile/Detail.html:40
		if err == nil {
//line views/vfile/Detail.html:40
			qw422016.N().S(`    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>EXIF Name</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vfile/Detail.html:50
			for k, v := range exif {
//line views/vfile/Detail.html:50
				qw422016.N().S(`          <tr>
            <td>`)
//line views/vfile/Detail.html:52
				qw422016.E().S(k)
//line views/vfile/Detail.html:52
				qw422016.N().S(`</td>
            <td>`)
//line views/vfile/Detail.html:53
				qw422016.E().V(v)
//line views/vfile/Detail.html:53
				qw422016.N().S(`</td>
          </tr>
`)
//line views/vfile/Detail.html:55
			}
//line views/vfile/Detail.html:55
			qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vfile/Detail.html:59
		} else {
//line views/vfile/Detail.html:59
			qw422016.N().S(`    <em>File is binary and contains no exif header</em>
`)
//line views/vfile/Detail.html:61
		}
//line views/vfile/Detail.html:62
	}
//line views/vfile/Detail.html:62
	qw422016.N().S(`  </div>
`)
//line views/vfile/Detail.html:64
}

//line views/vfile/Detail.html:64
func WriteDetail(qq422016 qtio422016.Writer, path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Detail.html:64
	StreamDetail(qw422016, path, b, urlPrefix, additionalLinks, as, ps)
//line views/vfile/Detail.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Detail.html:64
}

//line views/vfile/Detail.html:64
func Detail(path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Detail.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Detail.html:64
	WriteDetail(qb422016, path, b, urlPrefix, additionalLinks, as, ps)
//line views/vfile/Detail.html:64
	qs422016 := string(qb422016.B)
//line views/vfile/Detail.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Detail.html:64
	return qs422016
//line views/vfile/Detail.html:64
}
