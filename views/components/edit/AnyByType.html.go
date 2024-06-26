// Code generated by qtc from "AnyByType.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/AnyByType.html:2
package edit

//line views/components/edit/AnyByType.html:2
import (
	"fmt"

	"github.com/kyleu/dbaudit/app/lib/types"
)

//line views/components/edit/AnyByType.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/AnyByType.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/AnyByType.html:8
func StreamAnyByType(qw422016 *qt422016.Writer, key string, id string, x any, t *types.Wrapped) {
//line views/components/edit/AnyByType.html:9
	switch t.Key() {
//line views/components/edit/AnyByType.html:10
	case types.KeyAny:
//line views/components/edit/AnyByType.html:11
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:12
	case types.KeyBit:
//line views/components/edit/AnyByType.html:13
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:14
	case types.KeyBool:
//line views/components/edit/AnyByType.html:15
		StreamBool(qw422016, key, id, x, false)
//line views/components/edit/AnyByType.html:16
	case types.KeyByte:
//line views/components/edit/AnyByType.html:17
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:18
	case types.KeyChar:
//line views/components/edit/AnyByType.html:19
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:20
	case types.KeyDate:
//line views/components/edit/AnyByType.html:21
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:22
	case types.KeyEnum:
//line views/components/edit/AnyByType.html:22
		qw422016.N().S(`<span title="enum:`)
//line views/components/edit/AnyByType.html:23
		qw422016.E().S(t.T.(*types.Enum).Ref)
//line views/components/edit/AnyByType.html:23
		qw422016.N().S(`">`)
//line views/components/edit/AnyByType.html:23
		qw422016.E().V(x)
//line views/components/edit/AnyByType.html:23
		qw422016.N().S(`</span>`)
//line views/components/edit/AnyByType.html:24
	case types.KeyEnumValue:
//line views/components/edit/AnyByType.html:25
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:26
	case types.KeyError:
//line views/components/edit/AnyByType.html:27
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:28
	case types.KeyFloat:
//line views/components/edit/AnyByType.html:29
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:30
	case types.KeyInt:
//line views/components/edit/AnyByType.html:31
		StreamInt(qw422016, key, id, x)
//line views/components/edit/AnyByType.html:32
	case types.KeyJSON:
//line views/components/edit/AnyByType.html:33
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:34
	case types.KeyList:
//line views/components/edit/AnyByType.html:35
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:36
	case types.KeyMap:
//line views/components/edit/AnyByType.html:37
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:38
	case types.KeyMethod:
//line views/components/edit/AnyByType.html:39
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:40
	case types.KeyNil:
//line views/components/edit/AnyByType.html:41
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:42
	case types.KeyOption:
//line views/components/edit/AnyByType.html:43
		StreamOption(qw422016, key, id, x, t.T.(*types.Option))
//line views/components/edit/AnyByType.html:44
	case types.KeyRange:
//line views/components/edit/AnyByType.html:45
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:46
	case types.KeyReference:
//line views/components/edit/AnyByType.html:47
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:48
	case types.KeySet:
//line views/components/edit/AnyByType.html:49
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:50
	case types.KeyString:
//line views/components/edit/AnyByType.html:51
		StreamString(qw422016, key, id, fmt.Sprint(x))
//line views/components/edit/AnyByType.html:52
	case types.KeyTime:
//line views/components/edit/AnyByType.html:53
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:54
	case types.KeyTimestamp:
//line views/components/edit/AnyByType.html:55
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:56
	case types.KeyTimestampZoned:
//line views/components/edit/AnyByType.html:57
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:58
	case types.KeyUnknown:
//line views/components/edit/AnyByType.html:59
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:60
	case types.KeyUUID:
//line views/components/edit/AnyByType.html:61
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:62
	case types.KeyValueMap:
//line views/components/edit/AnyByType.html:63
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:64
	case types.KeyXML:
//line views/components/edit/AnyByType.html:65
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:66
	default:
//line views/components/edit/AnyByType.html:67
		StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:68
	}
//line views/components/edit/AnyByType.html:69
}

//line views/components/edit/AnyByType.html:69
func WriteAnyByType(qq422016 qtio422016.Writer, key string, id string, x any, t *types.Wrapped) {
//line views/components/edit/AnyByType.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/AnyByType.html:69
	StreamAnyByType(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/AnyByType.html:69
}

//line views/components/edit/AnyByType.html:69
func AnyByType(key string, id string, x any, t *types.Wrapped) string {
//line views/components/edit/AnyByType.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/AnyByType.html:69
	WriteAnyByType(qb422016, key, id, x, t)
//line views/components/edit/AnyByType.html:69
	qs422016 := string(qb422016.B)
//line views/components/edit/AnyByType.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/AnyByType.html:69
	return qs422016
//line views/components/edit/AnyByType.html:69
}

//line views/components/edit/AnyByType.html:71
func StreamDefault(qw422016 *qt422016.Writer, key string, id string, x any, t types.Type) {
//line views/components/edit/AnyByType.html:73
	msg := fmt.Sprintf("unhandled type: %s (%T)", t.String(), x)
	if x == nil {
		x = "∅"
	}

//line views/components/edit/AnyByType.html:77
	qw422016.N().S(`<input title="`)
//line views/components/edit/AnyByType.html:78
	qw422016.E().S(msg)
//line views/components/edit/AnyByType.html:78
	qw422016.N().S(`" value="`)
//line views/components/edit/AnyByType.html:78
	qw422016.E().V(x)
//line views/components/edit/AnyByType.html:78
	qw422016.N().S(`" name="`)
//line views/components/edit/AnyByType.html:78
	qw422016.E().S(key)
//line views/components/edit/AnyByType.html:78
	qw422016.N().S(`" id="`)
//line views/components/edit/AnyByType.html:78
	qw422016.E().S(id)
//line views/components/edit/AnyByType.html:78
	qw422016.N().S(`" />`)
//line views/components/edit/AnyByType.html:79
}

//line views/components/edit/AnyByType.html:79
func WriteDefault(qq422016 qtio422016.Writer, key string, id string, x any, t types.Type) {
//line views/components/edit/AnyByType.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/AnyByType.html:79
	StreamDefault(qw422016, key, id, x, t)
//line views/components/edit/AnyByType.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/AnyByType.html:79
}

//line views/components/edit/AnyByType.html:79
func Default(key string, id string, x any, t types.Type) string {
//line views/components/edit/AnyByType.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/AnyByType.html:79
	WriteDefault(qb422016, key, id, x, t)
//line views/components/edit/AnyByType.html:79
	qs422016 := string(qb422016.B)
//line views/components/edit/AnyByType.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/AnyByType.html:79
	return qs422016
//line views/components/edit/AnyByType.html:79
}
