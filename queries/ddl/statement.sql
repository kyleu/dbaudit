-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StatementDrop() %}
if exists (select * from sysobjects where name='statement' and xtype='U')
drop table "statement";
-- {% endfunc %}

-- {% func StatementCreate() %}
if not exists (select * from sysobjects where name='statement' and xtype='U')
create table "statement" (
  "id" uniqueidentifier not null,
  "session_id" int not null,
  "action" nvarchar(255) not null,
  "succeeded" bit not null,
  "principal" nvarchar(max) not null,
  "database" nvarchar(max) not null,
  "filename" nvarchar(max) not null,
  "host" nvarchar(max) not null,
  "transaction_id" int not null,
  "client_ip" nvarchar(max) not null,
  "duration" int not null,
  "connection_id" uniqueidentifier not null,
  "rows_affected" int not null,
  "rows_returned" int not null,
  "sql" nvarchar(max) not null,
  "types" nvarchar(max) not null,
  "values" nvarchar(max) not null,
  "occurred" datetime not null,
  primary key ("id")
);

if not exists (select * from sys.indexes where name='statement' and object_id=object_id('statement__action_idx'))
create index "statement__action_idx" on "statement" ("action");
-- {% endfunc %}
