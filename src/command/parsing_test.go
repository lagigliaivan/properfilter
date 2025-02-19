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
