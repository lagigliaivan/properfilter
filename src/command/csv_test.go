package command_test

import (
	"testing"

	"github.com/properfilter/src/command"
	"github.com/stretchr/testify/assert"
)

func TestCsvToProperties(t *testing.T) {
	csvLine := "6217 S Greenwood Ave,10000.30,80,2,1,foo,garage/grill/swimmingpool,low"

	property, err := command.CsvToProperty(csvLine)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	assert.Equal(t, "6217 S Greenwood Ave", property.Name)
	assert.Equal(t, float32(10000.30), property.Price)
	assert.Equal(t, 80, property.SquareFootage)
	assert.Equal(t, 2, property.Rooms)
	assert.Equal(t, 1, property.Bathrooms)
	assert.Equal(t, "foo", property.Description)
	assert.Equal(t, []string{"garage", "grill", "swimmingpool"}, property.Ammenities)
	assert.Equal(t, "low", property.Lighting)
}

func TestCsvToPropertiesReturnsErrorWhenMissingField(t *testing.T) {
	csvLine := "6217 S Greenwood Ave,10000.30,80,2,1,foo,garage/grill/swimmingpool"

	property, err := command.CsvToProperty(csvLine)
	assert.Error(t, err)
	assert.Nil(t, property)
}
