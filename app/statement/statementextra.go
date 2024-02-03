package statement

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func (s *Statement) Simple() bool {
	return len(s.Values) == 0
}

func (s *Statement) PlainSQL() string {
	if s.Simple() {
		return s.SQL
	}
	ret := s.SQL
	for _, k := range lo.Reverse(s.Values.Keys()) {
		println(k)
		v := s.Values[k]
		t := s.Types.GetStringOpt(k)
		switch {
		case v == nil:
			ret = strings.ReplaceAll(ret, k, "null")
		case t == "uniqueidentifier":
			ret = strings.ReplaceAll(ret, k, "'"+fmt.Sprint(v)+"'")
		case strings.HasPrefix(t, "nvarchar(") || strings.HasPrefix(t, "varchar("):
			ret = strings.ReplaceAll(ret, k, "'"+fmt.Sprint(v)+"'")
		case strings.HasPrefix(t, "decimal("):
			ret = strings.ReplaceAll(ret, k, fmt.Sprint(v))
		case t == "int":
			ret = strings.ReplaceAll(ret, k, fmt.Sprint(v))
		case t == "bit":
			ret = strings.ReplaceAll(ret, k, fmt.Sprint(s.Values.GetIntOpt(k)))
		default:
			ret = strings.ReplaceAll(ret, k, t+":"+fmt.Sprint(v))
		}
	}
	return ret
}
