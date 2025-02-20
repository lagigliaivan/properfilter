package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewDescription(args string) (PropertyFilter, error) {
	return func(p model.Property) bool { return ContainsDescription(p, args) }, nil
}

func ContainsDescription(p model.Property, name string) bool {
	return strings.Contains(p.Description, name)
}
