package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/properfilter/src/model"
)

func CsvToProperty(csvLine string) (*model.Property, error) {
	sp := strings.Split(csvLine, ",")
	if len(sp) != 10 {
		return nil, fmt.Errorf("invalid CSV line:%s", csvLine)
	}

	price, err := strconv.ParseFloat(sp[1], 32)
	if err != nil {
		return nil, err
	}

	squareFootage, err := strconv.ParseInt(sp[2], 10, 32)
	if err != nil {
		return nil, err
	}

	rooms, err := strconv.ParseInt(sp[3], 10, 32)
	if err != nil {
		return nil, err
	}

	bathrooms, err := strconv.ParseInt(sp[4], 10, 32)
	if err != nil {
		return nil, err
	}

	c, err := model.NewCoordinatesFromString(sp[8], sp[9])
	if err != nil {
		return nil, err
	}

	return &model.Property{
		Name:          sp[0],
		Price:         float32(price),
		SquareFootage: int(squareFootage),
		Rooms:         int(rooms),
		Bathrooms:     int(bathrooms),
		Description:   sp[5],
		Ammenities:    strings.Split(sp[6], "/"),
		Lighting:      sp[7],
		Location:      *c,
	}, nil
}
