-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SizeInfo(dbType string) %}
-- {% switch dbType %}
-- {% case "sqlserver" %}
select 'TODO';
-- {% default %}
select 'unhandled database type [{%s dbType %}]';
-- {% endswitch %}
-- {% endfunc %}
