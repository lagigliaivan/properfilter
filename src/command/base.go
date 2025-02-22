package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func ParseArgsValues(args string) ([]string, error) {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil, ErrInvalidNumberOfArguments
	}

	return ops, nil
}

func ContainsElement[T comparable](element T, slice []T) bool {
	for _, a := range slice {
		if element == a {
			return true
		}
	}

	return false
}

func OR(args string, f func(string) (PropertyFilter, error)) (PropertyFilter, error) {
	orFilters := make([]PropertyFilter, 0)
	or := strings.Split(args, "|")

	if len(or) > 1 {
		for o := range or {
			p, err := f(or[o])
			if err != nil {
				return nil, err
			}
			orFilters = append(orFilters, p)
		}

		return func(p model.Property) bool {
			for _, f := range orFilters {
				if f(p) {
					return true
				}
			}
			return false
		}, nil
	}

	return nil, nil
}
