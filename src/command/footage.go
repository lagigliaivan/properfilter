package command

import (
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewSquareFootage(args string) *Arguments {
	ops := strings.Split(args, ":")
	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = IntValue(size, EqualFootage)
	evals[lessThan] = IntValue(size, LessThanFootage)
	evals[greaterThan] = IntValue(size, GreaterThanFootage)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
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
