package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewDescription(args string) (PropertyFilter, error) {
	filters, err := OR(args, NewDescription)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		return filters, nil
	}

	return func(p model.Property) bool { return ContainsDescription(p, args) }, nil
}

func ContainsDescription(p model.Property, name string) bool {
	return strings.Contains(p.Description, name)
}
