package command

import (
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewPrice(args string) *Arguments {
	ops := strings.Split(args, ":")
	price, err := strconv.ParseFloat(ops[1], 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = FloatValue(price, EqualPrice)
	evals[lessThan] = FloatValue(price, LessThanPrice)
	evals[greaterThan] = FloatValue(price, GreaterThanPrice)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
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
