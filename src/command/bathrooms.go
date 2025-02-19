package command

import (
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewBathrooms(args string) *Arguments {
	ops := strings.Split(args, ":")
	price, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = IntValue(price, EqualBathRooms)
	evals[lessThan] = IntValue(price, LessThanBathRooms)
	evals[greaterThan] = IntValue(price, GreaterThanBathRooms)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
}
func EqualBathRooms(p model.Property, v int64) bool {
	return int64(p.Bathrooms) == v
}

func LessThanBathRooms(p model.Property, v int64) bool {
	return int64(p.Bathrooms) < v
}

func GreaterThanBathRooms(p model.Property, v int64) bool {
	return int64(p.Bathrooms) > v
}
