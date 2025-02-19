package command

import (
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewRooms(args string) *Arguments {
	ops := strings.Split(args, ":")
	operand := ops[1]
	operator := ops[0]

	size, err := strconv.ParseInt(operand, 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = IntValue(size, EqualRoom)
	evals[lessThan] = IntValue(size, LessThanRoom)
	evals[greaterThan] = IntValue(size, GreaterThanRoom)

	return &Arguments{
		evals:    evals,
		operator: operator,
	}
}
func EqualRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) == v
}

func LessThanRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) < v
}

func GreaterThanRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) > v
}
