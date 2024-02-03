package parse

const sqlQuery = `select top %d
  event_time,
  session_id,
  action_id,
  sequence_group_id,
  sequence_number,
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
  sys.fn_get_audit_file('%s', default, default);`
