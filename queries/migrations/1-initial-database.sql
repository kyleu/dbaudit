-- {% import "github.com/kyleu/dbaudit/queries/ddl" %}
-- {% import "github.com/kyleu/dbaudit/queries/seeddata" %}

-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {%= seeddata.SeedDataAll() %}
-- {% endfunc %}
