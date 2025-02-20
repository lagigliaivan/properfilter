package command

import (
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewSquareFootage(args string) (PropertyFilter, error) {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil, ErrInvalidNumberOfArguments
	}

	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil, ErrInvalidOperator
	}

	switch ops[0] {
	case equal:
		return IntValue(size, EqualFootage), nil
	case lessThan:
		return IntValue(size, LessThanFootage), nil
	case greaterThan:
		return IntValue(size, GreaterThanFootage), nil
	}

	return nil, ErrInvalidOperator
}

func EqualFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) == v
}

func LessThanFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) < v
}

func GreaterThanFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) > v
}
