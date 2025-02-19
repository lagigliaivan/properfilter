package model

type (
	Property struct {
		Name          string
		SquareFootage int
		Lighting      string //'low' | 'medium' | 'high',
		Price         float32
		Rooms         int
		Bathrooms     int
		Location      *Coordinates
		Description   string
		Ammenities    []string // Record<string, boolean> //yard, garage, pool, etc
	}

	Coordinates struct {
		Lat  int
		Long int
	}

	Properties []Property
)
