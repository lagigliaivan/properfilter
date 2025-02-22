package command_test

import (
	"testing"

	"github.com/properfilter/src/command"
	"github.com/stretchr/testify/assert"
)

func TestCsvToProperties(t *testing.T) {
	csvLine := "1001 W Elm St,12000.50,75,2,1,Cozy home,garage/grill,medium,-33.513270,62.930154"

	property, err := command.CsvToProperty(csvLine)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	assert.Equal(t, "1001 W Elm St", property.Address)
	assert.Equal(t, float32(12000.50), property.Price)
	assert.Equal(t, 75, property.SquareFootage)
	assert.Equal(t, 2, property.Rooms)
	assert.Equal(t, 1, property.Bathrooms)
	assert.Equal(t, "Cozy home", property.Description)
	assert.Equal(t, []string{"garage", "grill"}, property.Ammenities)
	assert.Equal(t, "medium", property.Lighting)
	assert.Equal(t, float32(-33.513270), property.Location.Lat)
	assert.Equal(t, float32(62.930154), property.Location.Long)
}

func TestCsvToPropertiesReturnsErrorWhenMissingField(t *testing.T) {
	csvLine := "6217 S Greenwood Ave,10000.30,80,2,1,foo,garage/grill/swimmingpool"

	property, err := command.CsvToProperty(csvLine)
	assert.Error(t, err)
	assert.Nil(t, property)
}
