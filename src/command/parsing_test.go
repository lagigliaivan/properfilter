package command_test

import (
	"context"
	"testing"

	"github.com/properfilter/src/command"
	"github.com/properfilter/src/model"
	"github.com/stretchr/testify/assert"
)

// - Equal, lessThan and greater-than
// - Inclusion (i.e. must include garage)
// - Matching (description must include some word)
// - Distance (for location)

// Name          string
// SquareFootage int
// Lighting      string //'low' | 'medium' | 'high',
// Price         int
// Rooms         int
// Bathrooms     int
// Location      *Coordinates
// Description   string
// Ammenities    []string

// args := []string{"--name", "foo", "--price", "eq=100"}
func TestPriceEqualsTo(t *testing.T) {
	args := []string{"--price", "eq:100"}
	cmd, err := command.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cmd)

	properties := model.Properties{
		{Name: "foo", Price: 100},
		{Name: "y", Price: 200},
		{Name: "z", Price: 300},
	}

	ctx := context.Background()
	result := cmd.Execute(ctx, properties)

	assert.Len(t, result, 1)
	assert.Equal(t, properties[0].Name, result[0].Name)
}
func TestPriceEqualsToAndNameEqualsTo(t *testing.T) {
	ctx := context.Background()

	args := []string{"--price", "eq:100", "--name", "foo"}
	cmd, err := command.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cmd)

	properties := model.Properties{
		{Name: "foo", Price: 100},
		{Name: "y", Price: 200},
		{Name: "z", Price: 300},
	}

	result := cmd.Execute(ctx, properties)

	assert.Len(t, result, 1)
	assert.Equal(t, properties[0].Name, result[0].Name)

	properties = model.Properties{
		{Name: "y", Price: 100},
		{Name: "z", Price: 300},
	}

	result = cmd.Execute(ctx, properties)
	assert.Len(t, result, 0)
}

func TestPriceLessThan(t *testing.T) {
	ctx := context.Background()

	args := []string{"--price", "lt:250"}
	cmd, err := command.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cmd)

	properties := model.Properties{
		{Name: "foo", Price: 100},
		{Name: "y", Price: 200},
		{Name: "z", Price: 300},
	}

	result := cmd.Execute(ctx, properties)

	assert.Len(t, result, 2)
	assert.Contains(t, result, properties[0])
	assert.Contains(t, result, properties[1])
}
func TestPriceGreaterThan(t *testing.T) {
	ctx := context.Background()

	args := []string{"--price", "gt:250"}
	cmd, err := command.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cmd)

	properties := model.Properties{
		{Name: "foo", Price: 100},
		{Name: "y", Price: 200},
		{Name: "z", Price: 300},
	}

	result := cmd.Execute(ctx, properties)

	assert.Len(t, result, 1)
	assert.Equal(t, properties[2].Name, result[0].Name)
}
func TestSquareFootageEqualsTo(t *testing.T) {
	ctx := context.Background()

	args := []string{"--square-footage", "eq:80"}
	cmd, err := command.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cmd)

	properties := model.Properties{
		{Name: "foo", Price: 100, SquareFootage: 80},
		{Name: "y", Price: 200, SquareFootage: 90},
		{Name: "z", Price: 300, SquareFootage: 100},
	}

	result := cmd.Execute(ctx, properties)

	assert.Len(t, result, 1)
	assert.Equal(t, properties[0].SquareFootage, result[0].SquareFootage)
}

func TestRooms(t *testing.T) {
	properties := model.Properties{
		{Name: "6217 S Greenwood Ave", Price: 100, SquareFootage: 80, Rooms: 2},
		{Name: "6201-03 S King", Price: 200, SquareFootage: 90, Rooms: 3},
		{Name: "3001-19 E 79th - 3001-19 E 79th St Chicago", Price: 300, SquareFootage: 100, Rooms: 4},
	}

	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "rooms equals to",
			args:     []string{"--rooms", "eq:3"},
			expected: []model.Property{properties[1]},
		},
		{
			name:     "rooms less than",
			args:     []string{"--rooms", "lt:4"},
			expected: []model.Property{properties[0], properties[1]},
		},
		{
			name:     "rooms greater than",
			args:     []string{"--rooms", "gt:3"},
			expected: []model.Property{properties[2]},
		},
	}

	for _, tc := range uc {
		t.Run(tc.name, func(t *testing.T) {
			cmd, err := command.Parse(tc.args)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotNil(t, cmd)

			result := cmd.Execute(context.Background(), properties)

			assert.Len(t, result, len(tc.expected))
			for _, property := range result {
				assert.Contains(t, result, property)
			}
		})
	}

}
