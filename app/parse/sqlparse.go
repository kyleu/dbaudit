package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/dbaudit/app/util"
)

var (
	partsRegex = regexp.MustCompile(`(?s)N'(.*?)',\s*(N'.*?'),\s*(.*)`)
)

func parseSQL(input string) (string, util.ValueMap, util.ValueMap, error) {
	if !strings.HasPrefix(input, "exec sp_executesql") {
		return input, nil, nil, nil
	}

	matches := partsRegex.FindStringSubmatch(input)
	if len(matches) != 4 {
		return input, nil, nil, errors.New("input does not match the expected format")
	}

	query, argTypesStr, argValuesStr := matches[1], matches[2], matches[3]

	typesParts := util.StringSplitAndTrim(argTypesStr, "@")
	types := make(util.ValueMap, len(typesParts))
	for _, part := range typesParts {
		if part == "N'" {
			continue
		}
		part = strings.TrimSuffix(part, ",")
		part = strings.TrimSuffix(part, "'")
		num, t := util.StringSplit(part, ' ', true)
		types["@"+num] = t
	}

	valsParts := util.StringSplitAndTrim(argValuesStr, "@")
	vals := make(util.ValueMap, len(valsParts))
	for _, part := range valsParts {
		part = strings.TrimSuffix(part, ",")
		num, vStr := util.StringSplit(part, '=', true)
		if !types.HasKey("@" + num) {
			return input, nil, nil, errors.Errorf("missing type specifier for column [%s]", num)
		}
		t := types.GetStringOpt("@" + num)

		var v any = vStr

		if vStr == "NULL" {
			v = nil
		} else {
			switch {
			case strings.HasPrefix(t, "uniqueidentifier"):
				uStr := strings.TrimPrefix(strings.TrimSuffix(vStr, "'"), "'")
				v = util.UUIDFromStringOK(uStr)
			case strings.HasPrefix(t, "nvarchar"):
				v = strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(vStr, "N"), "'"), "'")
			case t == "bit":
				v = vStr != "0"
			case t == "int":
				var err error
				v, err = strconv.ParseInt(vStr, 10, 64)
				if err != nil {
					return "", nil, nil, err
				}
			case strings.HasPrefix(t, "decimal"):
				var err error
				v, err = strconv.ParseFloat(vStr, 64)
				if err != nil {
					return input, nil, nil, err
				}
			}
		}
		//vals[num] = fmt.Sprintf("[%s]: %v", t, v)
		vals["@"+num] = v
	}

	return query, types, vals, nil
}
