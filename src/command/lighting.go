package command

import (
	"github.com/properfilter/src/model"
)

func NewLighting(args string) (PropertyFilter, error) {
	ops, err := ParseArgsValues(args)
	if err != nil {
		return nil, err
	}

	lighting := ops[1]

	switch ops[0] {
	case equal:
		return StringValue(lighting, EqualLighting), nil
	case lessThan:
		return StringValue(lighting, LessThanLighting), nil
	case greaterThan:
		return StringValue(lighting, GreaterThanLighting), nil
	}

	return nil, ErrInvalidOperator
}
func EqualLighting(p model.Property, v string) bool {
	return p.Lighting == v
}

func LessThanLighting(p model.Property, v string) bool {
	return getLightingValue(p.Lighting) < getLightingValue(v)
}

func GreaterThanLighting(p model.Property, v string) bool {
	return getLightingValue(p.Lighting) > getLightingValue(v)
}

func getLightingValue(l string) int {
	switch l {
	case "low":
		return 0
	case "medium":
		return 1
	case "high":
		return 2
	default:
		return 0
	}
}
