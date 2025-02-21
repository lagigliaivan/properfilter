package command

import (
	"math"
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func NewDistance(args string) (PropertyFilter, error) {
	ops, err := ParseArgsValues(args)
	if err != nil {
		return nil, err
	}

	coords := strings.Split(ops[1], ",")
	if len(coords) != 3 {
		return nil, ErrInvalidNumberOfArguments
	}

	if len(coords) != 3 {
		return nil, ErrInvalidNumberOfArguments
	}

	c, err := model.NewCoordinatesFromString(coords[1], coords[2])
	if err != nil {
		return nil, err
	}

	ratio, err := strconv.ParseInt(coords[0], 10, 32)
	if err != nil {
		return nil, err
	}

	switch ops[0] {
	case lessThan:
		return func(p model.Property) bool { return LessThanDistance(p.Location, *c, int(ratio)) }, nil
	case greaterThan:
		return func(p model.Property) bool { return GreaterThanDistance(p.Location, *c, int(ratio)) }, nil
	}

	return nil, ErrInvalidOperator
}

func LessThanDistance(p model.Coordinates, d model.Coordinates, km int) bool {
	return distance(p, d) < float64(km)
}

func GreaterThanDistance(p model.Coordinates, d model.Coordinates, km int) bool {
	return distance(p, d) > float64(km)
}

// func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
func distance(a model.Coordinates, b model.Coordinates) float64 {
	radlat1 := float64(math.Pi * a.Lat / 180)
	radlat2 := float64(math.Pi * b.Lat / 180)

	theta := float64(a.Long - b.Long)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344

	return dist
}
