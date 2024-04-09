// Package db - Content managed by Project Forge, see [projectforge.md] for details.
package db

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

var (
	EnginePostgres = Engine{Key: "postgres", Name: "PostgreSQL"}
	EngineMysql    = Engine{Key: "mysql", Name: "MySQL"}
	EngineMssql    = Engine{Key: "mssql", Name: "Microsoft SQL Server"}
	EngineSqlite   = Engine{Key: "sqlite", Name: "SQLite"}

	AllEngines = Engines{EnginePostgres, EngineMysql, EngineMssql, EngineSqlite}
)

type Engine struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (e Engine) String() string {
	if e.Name != "" {
		return e.Name
	}
	return e.Key
}

func (e Engine) Matches(xx Engine) bool {
	return e.Key == xx.Key
}

func (e Engine) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(e.Key, false), nil
}

func (e *Engine) UnmarshalJSON(data []byte) error {
	var key string
	if err := util.FromJSON(data, &key); err != nil {
		return err
	}
	*e = AllEngines.Get(key, nil)
	return nil
}

func (e Engine) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(e.Key, start)
}

func (e *Engine) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*e = AllEngines.Get(key, nil)
	return nil
}

func (e Engine) Value() (driver.Value, error) {
	return e.Key, nil
}

func (e *Engine) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*e = AllEngines.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan Engine enum from value [%v]", value)
}

func EngineParse(logger util.Logger, keys ...string) Engines {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) Engine {
		return AllEngines.Get(x, logger)
	})
}

type Engines []Engine

func (e Engines) Keys() []string {
	return lo.Map(e, func(x Engine, _ int) string {
		return x.Key
	})
}

func (e Engines) Strings() []string {
	return lo.Map(e, func(x Engine, _ int) string {
		return x.String()
	})
}

func (e Engines) Help() string {
	return "Available options: [" + strings.Join(e.Strings(), ", ") + "]"
}

func (e Engines) Get(key string, logger util.Logger) Engine {
	for _, value := range e {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Engine] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return Engine{Key: "_error", Name: "error: " + msg}
}

func (e Engines) GetByName(name string, logger util.Logger) Engine {
	for _, value := range e {
		if value.Name == name {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Engine] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return Engine{Key: "_error", Name: "error: " + msg}
}

func (e Engines) Random() Engine {
	return e[util.RandomInt(len(e))]
}
