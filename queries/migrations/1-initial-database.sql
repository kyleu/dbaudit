-- {% import "github.com/kyleu/dbaudit/queries/ddl" %}

-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {% endfunc %}
