package command

import (
	"errors"
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

var (
	ErrInvalidNumberOfArguments = errors.New("invalid number of arguments")
	ErrInvalidOperator          = errors.New("invalid operator")
	ErrorNoOperator             = errors.New("no operator provided")
)

func NewRooms(args string) (PropertyFilter, error) {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil, ErrInvalidNumberOfArguments
	}

	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil, err
	}

	switch ops[0] {
	case equal:
		return IntValue(size, EqualRoom), nil
	case lessThan:
		return IntValue(size, LessThanRoom), nil
	case greaterThan:
		return IntValue(size, GreaterThanRoom), nil
	}

	return nil, ErrInvalidOperator
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
