// Package statement - Content managed by Project Forge, see [projectforge.md] for details.
package statement

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/dbaudit/app/util"
)

type Action struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (a Action) String() string {
	if a.Name != "" {
		return a.Name
	}
	return a.Key
}

func (a Action) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(a.Key, false), nil
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var key string
	if err := util.FromJSON(data, &key); err != nil {
		return err
	}
	*a = AllActions.Get(key, nil)
	return nil
}

func (a Action) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(a.Key, start)
}

func (a *Action) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*a = AllActions.Get(key, nil)
	return nil
}

func (a Action) Value() (driver.Value, error) {
	return a.Key, nil
}

func (a *Action) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*a = AllActions.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan Action enum from value [%v]", value)
}

func ActionParse(logger util.Logger, keys ...string) Actions {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) Action {
		return AllActions.Get(x, logger)
	})
}

type Actions []Action

func (a Actions) Keys() []string {
	return lo.Map(a, func(x Action, _ int) string {
		return x.Key
	})
}

func (a Actions) Strings() []string {
	return lo.Map(a, func(x Action, _ int) string {
		return x.String()
	})
}

func (a Actions) Help() string {
	return "Available options: [" + strings.Join(a.Strings(), ", ") + "]"
}

func (a Actions) Get(key string, logger util.Logger) Action {
	for _, value := range a {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Action] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return Action{Key: "_error", Name: "error: " + msg}
}

func (a Actions) GetByName(name string, logger util.Logger) Action {
	for _, value := range a {
		if value.Name == name {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Action] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return Action{Key: "_error", Name: "error: " + msg}
}

func (a Actions) Random() Action {
	return a[util.RandomInt(len(a))]
}

var (
	ActionProcStart    = Action{Key: "proc_start", Name: "Procedure Started"}
	ActionProcComplete = Action{Key: "proc_complete", Name: "Procedure Complete"}
	ActionHealthcheck  = Action{Key: "healthcheck", Name: "Healthcheck"}
	ActionUnknown      = Action{Key: "unknown", Name: "Unknown"}

	AllActions = Actions{ActionProcStart, ActionProcComplete, ActionHealthcheck, ActionUnknown}
)
