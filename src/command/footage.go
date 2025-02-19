package command

import "github.com/properfilter/src/model"

func EqualFootage(v interface{}) func(model.Property) bool {
	footage, ok := v.(int64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return int64(p.SquareFootage) == footage
	}
}

func LessThanFootage(v interface{}) func(model.Property) bool {
	footage, ok := v.(int64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return int64(p.SquareFootage) < footage
	}
}

func GreaterThanFootage(v interface{}) func(model.Property) bool {
	footage, ok := v.(int64)
	if !ok {
		return func(p model.Property) bool {
			return false
		}
	}

	return func(p model.Property) bool {
		return int64(p.SquareFootage) > footage
	}
}
