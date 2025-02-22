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
	Filters []PropertyFilter
	Command struct {
		filters Filters
	}

	PropertyFilter            func(model.Property) bool
	PropertyFilterConstructor func(string) (PropertyFilter, error)
)

var (
	params = map[string]func(string, Filters) (Filters, error){
		"--price": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewPrice, f)
		},
		"--squarefootage": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewSquareFootage, f)
		},
		"--rooms": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewRooms, f)
		},
		"--bathrooms": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewBathrooms, f)
		},
		"--address": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewAddress, f)
		},
		"--description": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewDescription, f)
		},
		"--lighting": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewLighting, f)
		},
		"--ammenities": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewAmmenities, f)
		},
		"--distance": func(args string, f Filters) (Filters, error) {
			return parseParam(args, NewDistance, f)
		},
	}
)

func New(args []string) (*Command, error) {
	if len(args) == 0 {
		return nil, errors.New("no arguments provided")
	}

	filters := make(Filters, 0)
	for i := 0; i < len(args); i = i + 2 {
		param, ok := params[args[i]]
		if !ok {
			continue
		}
		var err error

		if i+1 >= len(args) {
			return nil, ErrorNoOperator
		}

		filters, err = param(args[i+1], filters)
		if err != nil {
			return nil, err
		}
	}

	if len(filters) <= 0 {
		return nil, errors.New("no valid arguments provided")
	}

	return &Command{filters: filters}, nil
}

func (c *Command) Filter(ctx context.Context, prop model.Property) *model.Property {
	for _, filter := range c.filters {
		if !filter(prop) {
			return nil
		}
	}

	return &prop
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

func parseParam(args string, constructor PropertyFilterConstructor, filters Filters) (Filters, error) {
	c, error := constructor(args)
	if error != nil {
		return nil, error
	}

	return append(filters, c), nil
}
