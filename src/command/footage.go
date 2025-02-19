package command

import "github.com/properfilter/src/model"

func EqualFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) == v
}

func LessThanFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) < v
}

func GreaterThanFootage(p model.Property, v int64) bool {
	return int64(p.SquareFootage) > v
}
