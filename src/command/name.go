package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewName(args string) *Arguments {
	evals := make(map[string]func(model.Property) bool)
	evals[""] = func(p model.Property) bool { return Contains(p, args) }

	return &Arguments{
		evals:    evals,
		operator: "",
	}
}
func Contains(p model.Property, name string) bool {
	return strings.Contains(p.Name, name)
}
