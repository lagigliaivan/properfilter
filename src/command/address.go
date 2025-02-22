package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewAddress(args string) (PropertyFilter, error) {
	filters, err := OR(args, NewAddress)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		return filters, nil
	}

	return func(p model.Property) bool { return Contains(p, args) }, nil
}

func Contains(p model.Property, name string) bool {
	return strings.Contains(p.Address, name)
}
