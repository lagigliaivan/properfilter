package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewLighting(args string) *Arguments {
	ops := strings.Split(args, ":")
	lighting := ops[1]

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = StringValue(lighting, EqualLighting)
	evals[lessThan] = StringValue(lighting, LessThanLighting)
	evals[greaterThan] = StringValue(lighting, GreaterThanLighting)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
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
