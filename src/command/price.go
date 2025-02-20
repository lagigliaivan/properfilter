package command

import (
	"fmt"
	"strconv"

	"github.com/properfilter/src/model"
)

func NewPrice(args string) (PropertyFilter, error) {
	ops, err := ParseOperator(args)
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(ops[1], 32)
	if err != nil {
		return nil, fmt.Errorf("invalid price %s", ops[1])
	}

	switch ops[0] {
	case equal:
		return FloatValue(price, EqualPrice), nil
	case lessThan:
		return FloatValue(price, LessThanPrice), nil
	case greaterThan:
		return FloatValue(price, GreaterThanPrice), nil
	}

	return nil, fmt.Errorf("invalid operator %s", ops[0])
}

func EqualPrice(p model.Property, v float64) bool {
	return float64(p.Price) == v
}

func LessThanPrice(p model.Property, v float64) bool {
	return float64(p.Price) < v
}

func GreaterThanPrice(p model.Property, v float64) bool {
	return float64(p.Price) > v
}
