-- {% func ConnectionSeedData() %}
insert into "connection" (
  "id", "name", "icon", "engine", "server", "port", "username", "password", "database", "schema", "conn_override"
) values (
  '10000000-0000-0000-0000-000000000001', 'Local SQL Server', 'star', 'mssql', 'localhost', 1433, 'sa', 'aStrongPassword123!', 'Groupmatics', '', ''
);
-- {% endfunc %}
