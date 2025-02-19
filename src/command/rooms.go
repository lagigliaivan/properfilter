package command

import "github.com/properfilter/src/model"

func EqualRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) == v
}

func LessThanRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) < v
}

func GreaterThanRoom(p model.Property, v int64) bool {
	return int64(p.Rooms) > v
}
