package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewDescription(args string) *Arguments {
	evals := make(map[string]func(model.Property) bool)
	evals[""] = func(p model.Property) bool { return ContainsDescription(p, args) }

	return &Arguments{
		evals:    evals,
		operator: "",
	}
}

func ContainsDescription(p model.Property, name string) bool {
	return strings.Contains(p.Description, name)
}
