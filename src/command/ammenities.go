package command

import (
	"github.com/properfilter/src/model"
)

func NewAmmenities(args string) (PropertyFilter, error) {
	ops, err := ParseArgsValues(args)
	if err != nil {
		return nil, err
	}

	ammenities := ops[1]

	if ops[0] != equal {
		return nil, ErrInvalidOperator
	}

	return StringValue(ammenities, equalAmmenities), nil
}

func equalAmmenities(p model.Property, v string) bool {
	for _, a := range p.Ammenities {
		if a == v {
			return true
		}
	}

	return false
}
