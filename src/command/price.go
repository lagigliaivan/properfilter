package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func EqualPrice(p model.Property, v float64) bool {
	return float64(p.Price) == v
}

func LessThanPrice(p model.Property, v float64) bool {
	return float64(p.Price) < v
}

func GreaterThanPrice(p model.Property, v float64) bool {
	return float64(p.Price) > v
}

func Contains(n string) func(model.Property) bool {
	return func(p model.Property) bool {
		return strings.Contains(n, p.Name)
	}
}
