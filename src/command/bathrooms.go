package command

import (
	"strconv"

	"github.com/properfilter/src/model"
)

func NewBathrooms(args string) (PropertyFilter, error) {
	ops, err := ParseOperator(args)
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil, err
	}

	switch ops[0] {
	case equal:
		return IntValue(price, EqualBathRooms), nil
	case lessThan:
		return IntValue(price, LessThanBathRooms), nil
	case greaterThan:
		return IntValue(price, GreaterThanBathRooms), nil
	}

	return nil, ErrInvalidOperator
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
