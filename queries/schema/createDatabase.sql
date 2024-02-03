-- {% func CreateDatabase() %}
create role "dbaudit" with login password 'dbaudit';

create database "dbaudit";
alter database "dbaudit" set timezone to 'utc';
grant all privileges on database "dbaudit" to "dbaudit";
-- {% endfunc %}
