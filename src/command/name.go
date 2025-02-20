package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewName(args string) (PropertyFilter, error) {
	return func(p model.Property) bool { return Contains(p, args) }, nil
}

func Contains(p model.Property, name string) bool {
	return strings.Contains(p.Name, name)
}
