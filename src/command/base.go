package command

import (
	"strings"
)

func ParseOperator(args string) ([]string, error) {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil, ErrInvalidNumberOfArguments
	}

	return ops, nil
}
