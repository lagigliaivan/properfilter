package command

import (
	"context"
	"errors"

	"github.com/properfilter/src/model"
)

const (
	equal       = "eq"
	lessThan    = "lt"
	greaterThan = "gt"
)

type (
	Command struct {
		filters Filters
	}

	Filter  func(model.Properties) model.Properties
	Filters []Filter

	Arguments struct {
		evals    map[string]func(model.Property) bool
		operator string
	}
)

var (
	params = map[string]func([]string, int, Filters) (Filters, error){
		"--price": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewPrice, f)
		},
		"--square-footage": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewSquareFootage, f)
		},
		"--rooms": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewRooms, f)
		},
		"--bathrooms": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewBathrooms, f)
		},
		"--name": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewName, f)
		},
		"--descrition": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewDescription, f)
		},
		"--ammenities": func(args []string, index int, f Filters) (Filters, error) {
			return parseParam(args, index, NewAmmenities, f)
		},
	}
)

func Parse(args []string) (*Command, error) {
	if len(args) == 0 {
		return nil, errors.New("no arguments provided")
	}

	filters := make(Filters, 0)
	for i := 0; i < len(args); i++ {
		p, ok := params[args[i]]
		if !ok {
			continue
		}
		var err error
		filters, err = p(args, i, filters)
		if err != nil {
			return nil, err
		}
	}

	if len(filters) <= 0 {
		return nil, errors.New("no valid arguments provided")
	}

	return &Command{filters: filters}, nil
}

func (c *Command) Execute(ctx context.Context, ps model.Properties) model.Properties {
	for _, f := range c.filters {
		ps = f(ps)
	}

	return ps
}

func (a *Arguments) GetFilter() (Filter, error) {
	operation, ok := a.evals[a.operator]
	if !ok {
		return nil, errors.New("operator not supported")
	}

	return NewFilter(operation), nil
}

func NewFilter(eval func(model.Property) bool) Filter {
	return func(ps model.Properties) model.Properties {
		result := make(model.Properties, 0)
		for _, p := range ps {
			if eval(p) {
				result = append(result, p)
			}
		}

		return result
	}
}

func IntValue(v interface{}, predicate func(model.Property, int64) bool) func(model.Property) bool {
	return func(p model.Property) bool {
		footage, ok := v.(int64)
		if !ok {
			return false
		}

		return predicate(p, footage)
	}
}

func FloatValue(v interface{}, predicate func(model.Property, float64) bool) func(model.Property) bool {
	return func(p model.Property) bool {
		price, ok := v.(float64)
		if !ok {
			return false
		}

		return predicate(p, price)
	}
}

func StringValue(v interface{}, predicate func(model.Property, string) bool) func(model.Property) bool {
	return func(p model.Property) bool {
		a, ok := v.(string)
		if !ok {
			return false
		}

		return predicate(p, a)
	}
}

func parseParam(args []string, i int, constructor func(string) *Arguments, filters Filters) (Filters, error) {
	if i+1 < len(args) {
		price := constructor(args[i+1])

		f, err := price.GetFilter()
		if err != nil {
			return nil, err
		}

		filters = append(filters, f)
	}

	return filters, nil
}
