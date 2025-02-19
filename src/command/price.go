package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewFilter(eval func(model.Property) bool) Filter {
	return func(ps model.Properties) model.Properties {
		result := make(model.Properties, 0)
		for _, p := range ps {
			if eval(p) {
				result = append(result, p)
			}
		}

		return result
	}
}

func EqualPrice(v interface{}) func(model.Property) bool {
	price, ok := v.(float64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return float64(p.Price) == price
	}
}

func LessThanPrice(v interface{}) func(model.Property) bool {
	price, ok := v.(float64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return float64(p.Price) < price
	}
}

func GreaterThanPrice(v interface{}) func(model.Property) bool {
	price, ok := v.(float64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return float64(p.Price) > price
	}
}

func Contains(n string) func(model.Property) bool {
	return func(p model.Property) bool {
		return strings.Contains(n, p.Name)
	}
}
