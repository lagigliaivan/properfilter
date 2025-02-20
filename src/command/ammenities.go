package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewAmmenities(args string) (PropertyFilter, error) {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil, ErrInvalidNumberOfArguments
	}

	ammenities := ops[1]

	if ops[0] != equal {
		return nil, ErrInvalidOperator
	}

	return StringValue(ammenities, EqualAmmenities), nil
}

func EqualAmmenities(p model.Property, v string) bool {
	for _, a := range p.Ammenities {
		if a == v {
			return true
		}
	}

	return false
}
