// Code generated by qtc from "auditfiles.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/auditfiles.sql:1
package queries

//line queries/auditfiles.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/auditfiles.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/auditfiles.sql:1
func StreamAuditFiles(qw422016 *qt422016.Writer, path string, limit int) {
//line queries/auditfiles.sql:1
	qw422016.N().S(`
select top `)
//line queries/auditfiles.sql:2
	qw422016.N().D(limit)
//line queries/auditfiles.sql:2
	qw422016.N().S(`
  event_time,
  session_id,
  action_id,
  sequence_number,
  sequence_group_id,
  succeeded,
  object_id,
  class_type,
  session_server_principal_name,
  database_name,
  object_name,
  statement,
  additional_information,
  file_name,
  audit_file_offset,
  transaction_id,
  client_ip,
  application_name,
  duration_milliseconds,
  response_rows,
  affected_rows,
  connection_id,
  host_name
from
  sys.fn_get_audit_file('`)
//line queries/auditfiles.sql:27
	qw422016.E().S(path)
//line queries/auditfiles.sql:27
	qw422016.N().S(`', default, default)
-- `)
//line queries/auditfiles.sql:28
}

//line queries/auditfiles.sql:28
func WriteAuditFiles(qq422016 qtio422016.Writer, path string, limit int) {
//line queries/auditfiles.sql:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/auditfiles.sql:28
	StreamAuditFiles(qw422016, path, limit)
//line queries/auditfiles.sql:28
	qt422016.ReleaseWriter(qw422016)
//line queries/auditfiles.sql:28
}

//line queries/auditfiles.sql:28
func AuditFiles(path string, limit int) string {
//line queries/auditfiles.sql:28
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/auditfiles.sql:28
	WriteAuditFiles(qb422016, path, limit)
//line queries/auditfiles.sql:28
	qs422016 := string(qb422016.B)
//line queries/auditfiles.sql:28
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/auditfiles.sql:28
	return qs422016
//line queries/auditfiles.sql:28
}
