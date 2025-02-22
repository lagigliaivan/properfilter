package command

import (
	"strings"
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
