package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewAmmenities(args string) (PropertyFilter, error) {
	filter, err := OR(args, NewAmmenities)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		return filter, nil
	}

	ops := strings.Split(args, ",")
	if len(ops) == 0 {
		return nil, ErrInvalidNumberOfArguments
	}

	return func(p model.Property) bool {
		return equalAmmenities(p, ops)
	}, nil
}

func equalAmmenities(p model.Property, ammenities []string) bool {
	count := 0
	for _, ammenity := range ammenities {
		if ContainsElement(ammenity, p.Ammenities) {
			count++
		}
	}

	return len(ammenities) == count
}
