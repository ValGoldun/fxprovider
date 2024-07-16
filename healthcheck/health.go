package healthcheck

import (
	"fmt"
	"strings"
)

type Health struct {
	IsOK       bool       `json:"isOK"`
	FailPolicy FailPolicy `json:"failPolicy"`
	Checks     []Check    `json:"checks"`
}

func (h Health) String() string {
	var errors []string

	for index, check := range h.Checks {
		if check.Error == nil {
			continue
		}

		errors[index] = fmt.Sprintf("%s: %v", check.Source, check.Error)
	}

	return fmt.Sprintf("health check fail (%s): %s", h.FailPolicy, strings.Join(errors, ";"))
}

type Check struct {
	Source string `json:"source"`
	Error  error  `json:"error,omitempty"`
}
