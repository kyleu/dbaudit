// Package database - Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

const whereSpaces = " where "

func SQLSelect(columns string, tables string, where string, orderBy string, limit int, offset int, placeholder string) string {
	return SQLSelectGrouped(columns, tables, where, "", orderBy, limit, offset, placeholder)
}

func SQLSelectSimple(columns string, tables string, placeholder string, where ...string) string {
	return SQLSelectGrouped(columns, tables, strings.Join(where, " and "), "", "", 0, 0, placeholder)
}

func SQLSelectGrouped(columns string, tables string, where string, groupBy string, orderBy string, limit int, offset int, placeholder string) string {
	if columns == "" {
		columns = "*"
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	groupByClause := ""
	if len(groupBy) > 0 {
		groupByClause = " group by " + groupBy
	}

	orderByClause := ""
	if len(orderBy) > 0 {
		orderByClause = " order by " + orderBy
	}

	limitClause := ""
	offsetClause := ""
	if placeholder == "@" {
		prefix := ""
		if limit > 0 {
			if orderBy == "" {
				prefix = fmt.Sprintf("top %d ", limit)
			} else {
				limitClause = fmt.Sprintf(" fetch next %d rows only", limit)
			}
		}
		if len(orderBy) > 0 && (offset > 0 || limit > 0) {
			offsetClause = fmt.Sprintf(" offset %d rows", offset)
		}
		return "select " + prefix + columns + " from " + tables + whereClause + groupByClause + orderByClause + offsetClause + limitClause
	}

	if limit > 0 {
		limitClause = fmt.Sprintf(" limit %d", limit)
	}
	if offset > 0 {
		offsetClause = fmt.Sprintf(" offset %d", offset)
	}
	return "select " + columns + " from " + tables + whereClause + groupByClause + orderByClause + limitClause + offsetClause
}

func SQLInsert(table string, columns []string, rows int, placeholder string) string {
	if rows <= 0 {
		rows = 1
	}
	colString := strings.Join(columns, ", ")
	var placeholders []string
	lo.Times(rows, func(i int) struct{} {
		ph := lo.FilterMap(columns, func(_ string, idx int) (string, bool) {
			switch placeholder {
			case "$", "":
				return fmt.Sprintf("$%d", (i*len(columns))+idx+1), true
			case "?":
				return "?", true
			case "@":
				return fmt.Sprintf("@p%d", (i*len(columns))+idx+1), true
			default:
				return "", false
			}
		})
		placeholders = append(placeholders, "("+strings.Join(ph, ", ")+")")
		return struct{}{}
	})
	return fmt.Sprintf("insert into %s (%s) values %s", table, colString, strings.Join(placeholders, ", "))
}

func SQLInsertReturning(table string, columns []string, rows int, returning []string, placeholder string) string {
	q := SQLInsert(table, columns, rows, placeholder)
	if len(returning) > 0 {
		q += " returning " + strings.Join(returning, ", ")
	}
	return q
}

func SQLUpdate(table string, columns []string, where string, placeholder string) string {
	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	stmts := lo.FilterMap(columns, func(col string, i int) (string, bool) {
		switch placeholder {
		case "$", "":
			return fmt.Sprintf("%s = $%d", col, i+1), true
		case "?":
			return fmt.Sprintf("%s = ?", col), true
		case "@":
			return fmt.Sprintf("%s = @p%d", col, i+1), true
		default:
			return "", false
		}
	})
	return fmt.Sprintf("update %s set %s%s", table, strings.Join(stmts, ", "), whereClause)
}

func SQLUpdateReturning(table string, columns []string, where string, returned []string, placeholder string) string {
	q := SQLUpdate(table, columns, where, placeholder)
	if len(returned) > 0 {
		q += " returning " + strings.Join(returned, ", ")
	}
	return q
}

func SQLUpsert(table string, columns []string, rows int, conflicts []string, updates []string, placeholder string) string {
	colNames := strings.Join(columns, ", ")
	params := make([]string, 0, rows)
	for rowIdx := range lo.Range(rows) {
		cols := make([]string, 0, len(columns))
		for colIdx := range columns {
			cols = append(cols, fmt.Sprintf("@p%d", (rowIdx*len(columns))+colIdx+1))
		}
		params = append(params, fmt.Sprintf("(%s)", strings.Join(cols, ", ")))
	}
	atSymbols := strings.Join(params, ", ")
	sourceJoin := strings.Join(lo.Map(conflicts, func(pk string, _ int) string {
		return fmt.Sprintf("%s.%s = source.%s", table, pk, pk)
	}), " and ")
	assignments := strings.Join(lo.Map(updates, func(c string, _ int) string {
		return fmt.Sprintf("%s = \"source\".%s", c, c)
	}), ", ")
	vals := strings.Join(lo.Map(updates, func(c string, _ int) string {
		return fmt.Sprintf("source.%s", c)
	}), ", ")
	sql := `merge into %s using (values %s) as source (%s) on %s when matched then update set %s when not matched then insert (%s) values (%s);`
	return fmt.Sprintf(sql, table, atSymbols, colNames, sourceJoin, assignments, colNames, vals)
}

func SQLDelete(table string, where string, placeholder string) string {
	if strings.TrimSpace(where) == "" {
		return fmt.Sprintf("attempt to delete from [%s] using placeholder [%s] with empty where clause", table, placeholder)
	}
	return "delete from " + table + whereSpaces + where
}

func SQLInClause(column string, numParams int, offset int, placeholder string) string {
	resBuilder := strings.Builder{}
	for index := 0; index < numParams; index++ {
		if index == 0 {
			resBuilder.WriteString(column + " in (")
		} else {
			resBuilder.WriteString(", ")
		}
		switch placeholder {
		case "$", "":
			resBuilder.WriteString(fmt.Sprintf("$%d", index+offset+1))
		case "?":
			resBuilder.WriteString("?")
		case "@":
			resBuilder.WriteString(fmt.Sprintf("@p%d", index+offset+1))
		}
	}
	resBuilder.WriteString(")")
	return resBuilder.String()
}
