-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func ConnectionDrop() %}
if exists (select * from sysobjects where name='connection' and xtype='U')
drop table "connection";
-- {% endfunc %}

-- {% func ConnectionCreate() %}
if not exists (select * from sysobjects where name='connection' and xtype='U')
create table "connection" (
  "id" uniqueidentifier not null,
  "name" nvarchar(max) not null,
  "icon" nvarchar(max) not null,
  "engine" nvarchar(255) not null,
  "server" nvarchar(max) not null,
  "port" int not null,
  "username" nvarchar(max) not null,
  "password" nvarchar(max) not null,
  "database" nvarchar(max) not null,
  "schema" nvarchar(max) not null,
  "conn_override" nvarchar(max) not null,
  primary key ("id")
);

if not exists (select * from sys.indexes where name='connection' and object_id=object_id('connection__engine_idx'))
create index "connection__engine_idx" on "connection" ("engine");
-- {% endfunc %}
