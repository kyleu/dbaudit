package parse

import "github.com/kyleu/dbaudit/app/statement"

type Result struct {
	Events     Events               `json:"events,omitempty"`
	Statements statement.Statements `json:"statements,omitempty"`
	X          any                  `json:"x,omitempty"`
}
