package model

import (
	"fmt"
	"strconv"
)

type (
	Property struct {
		Name          string
		SquareFootage int
		Lighting      string //'low' | 'medium' | 'high',
		Price         float32
		Rooms         int
		Bathrooms     int
		Location      Coordinates
		Description   string
		Ammenities    []string // Record<string, boolean> //yard, garage, pool, etc
	}

	Coordinates struct {
		Lat  float32
		Long float32
	}

	Properties []Property
)

func (p *Property) String() string {
	return fmt.Sprintf("%s,%f,%d,%d,%s", p.Name, p.Price, p.Bathrooms, p.Bathrooms, p.Ammenities)
}

func NewCoordinatesFromString(lat string, long string) (*Coordinates, error) {
	lt, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil, err
	}

	lg, err := strconv.ParseFloat(long, 32)
	if err != nil {
		return nil, err
	}

	return &Coordinates{Lat: float32(lt), Long: float32(lg)}, nil
}
