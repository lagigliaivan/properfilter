package command_test

import (
	"context"
	"errors"
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
var dataSet = model.Properties{
	{
		Name:          "6217 S Greenwood Ave",
		Price:         100,
		SquareFootage: 80,
		Rooms:         2,
		Bathrooms:     1,
		Description:   "foo",
		Ammenities:    []string{"garage"},
		Lighting:      "low",
		Location:      model.Coordinates{Lat: float32(-33.20), Long: float32(-63.430154)},
	},
	{
		Name:          "6201-03 S King",
		Price:         200,
		SquareFootage: 90,
		Rooms:         3,
		Bathrooms:     2,
		Description:   "bar",
		Lighting:      "medium",
		Ammenities:    []string{"swimmingpool"},
		Location:      model.Coordinates{Lat: float32(-33.013270), Long: float32(-63.45)},
	},
	{
		Name:          "3001-19 E 79th - 3001-19 E 79",
		Price:         300,
		SquareFootage: 100,
		Rooms:         4,
		Bathrooms:     3,
		Description:   "baz",
		Lighting:      "high",
		Ammenities:    []string{"swimmingpool", "garage"},
		Location:      model.Coordinates{Lat: float32(-33.05), Long: float32(-63.430154)},
	},
}

func TestName(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "name equals to",
			args:     []string{"--name", "6217 S Greenwood Ave"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "name contains",
			args:     []string{"--name", "S"},
			expected: []model.Property{dataSet[0], dataSet[1]},
		},
		{
			name:     "name is not present",
			args:     []string{"--name", "non existing"},
			expected: []model.Property{},
		},
	}

	run(t, uc)
}

func TestPrice(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "price equals to",
			args:     []string{"--price", "eq:200"},
			expected: []model.Property{dataSet[1]},
		},
		{
			name:     "priceless than",
			args:     []string{"--price", "lt:100"},
			expected: []model.Property{},
		},
		{
			name:     "price greater than",
			args:     []string{"--price", "gt:250"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}

func TestSquareFootage(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "square-footage equals to",
			args:     []string{"--square-footage", "eq:80"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "square-footage less than",
			args:     []string{"--square-footage", "lt:100"},
			expected: []model.Property{dataSet[0], dataSet[1]},
		},
		{
			name:     "square-footage greater than",
			args:     []string{"--square-footage", "gt:95"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}

func TestRooms(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "rooms equals to",
			args:     []string{"--rooms", "eq:3"},
			expected: []model.Property{dataSet[1]},
		},
		{
			name:     "rooms less than",
			args:     []string{"--rooms", "lt:4"},
			expected: []model.Property{dataSet[0], dataSet[1]},
		},
		{
			name:     "rooms greater than",
			args:     []string{"--rooms", "gt:3"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}
func TestBathRooms(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "Bathrooms equals to",
			args:     []string{"--bathrooms", "eq:1"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "Bathrooms less than",
			args:     []string{"--bathrooms", "lt:2"},
			expected: []model.Property{dataSet[1]},
		},
		{
			name:     "Bathrooms greater than",
			args:     []string{"--bathrooms", "gt:2"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}

func TestLighting(t *testing.T) {
	//'low' | 'medium' | 'high',
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "lighting equals to",
			args:     []string{"--lighting", "eq:low"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "lighting less than",
			args:     []string{"--lighting", "lt:high"},
			expected: []model.Property{dataSet[0], dataSet[1]},
		},
		{
			name:     "lighting greater than",
			args:     []string{"--lighting", "gt:medium"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}
func TestDescription(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name:     "name equals to",
			args:     []string{"--descrition", "foo"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "name contains",
			args:     []string{"--descrition", "fo"},
			expected: []model.Property{dataSet[0]},
		},
		{
			name:     "name is not present",
			args:     []string{"--descrition", "xxx"},
			expected: []model.Property{},
		},
	}

	run(t, uc)
}

func TestAmmenities(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name: "ammenities must include",
			args: []string{"--ammenities", "eq:garage"},

			expected: []model.Property{dataSet[0], dataSet[2]},
		},
		{
			name:     "ammenities must include",
			args:     []string{"--ammenities", "eq:swimmingpool"},
			expected: []model.Property{dataSet[1], dataSet[2]},
		},
		{
			name:     "ammenities must include more than one",
			args:     []string{"--ammenities", "eq:swimmingpool", "--ammenities", "eq:garage"},
			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}

func TestMixingOperators(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name: "ammenities must include",
			args: []string{"--ammenities", "eq:garage", "--price", "gt:100", "--rooms", "lt:5"},

			expected: []model.Property{dataSet[2]},
		},
	}

	run(t, uc)
}
func TestErrorsInParams(t *testing.T) {
	uc := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "using a token separator different from :",
			args: []string{"--ammenities", "eq=garage"},
			err:  command.ErrInvalidNumberOfArguments,
		},
		{
			name: "missing parameter operator",
			args: []string{"--ammenities"},
			err:  command.ErrorNoOperator,
		},
	}

	runOnError(t, uc)
}

func TestDistance(t *testing.T) {
	uc := []struct {
		name     string
		args     []string
		expected []model.Property
	}{
		{
			name: "less than 5km away from the reference point",
			args: []string{"--distance", "lt:5,-33.013270,-63.430154"},

			expected: []model.Property{dataSet[1], dataSet[2]},
		},
		{
			name: "more than 5km away from the reference point",
			args: []string{"--distance", "gt:5,-33.013270,-63.430154"},

			expected: []model.Property{dataSet[0]},
		},
	}

	run(t, uc)
}

func run(t *testing.T, uc []struct {
	name     string
	args     []string
	expected []model.Property
}) {
	for _, tc := range uc {
		t.Run(tc.name, func(t *testing.T) {
			cmd, err := command.New(tc.args)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotNil(t, cmd)

			result := make(model.Properties, 0)
			for _, property := range dataSet {
				p := cmd.Filter(context.Background(), property)
				if p != nil {
					result = append(result, *p)
				}
			}

			assert.Len(t, result, len(tc.expected))
			for _, property := range result {
				assert.Contains(t, result, property)
			}
		})
	}
}

func runOnError(t *testing.T, uc []struct {
	name string
	args []string
	err  error
}) {
	for _, tc := range uc {
		t.Run(tc.name, func(t *testing.T) {
			_, err := command.New(tc.args)
			if err != nil {
				assert.True(t, errors.Is(err, tc.err))
				return
			}

			t.Fatal("error expected")
		})
	}
}
