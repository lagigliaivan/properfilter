package command

import (
	"context"
	"errors"
	"strconv"
	"strings"

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

func Parse(args []string) (*Command, error) {
	if len(args) == 0 {
		return nil, errors.New("no arguments provided")
	}

	filters := make(Filters, 0)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--price":
			if i+1 < len(args) {
				price := NewPrice(args[i+1])

				f, err := price.GetFilter()
				if err != nil {
					return nil, err
				}

				filters = append(filters, f)
			}
		case "--square-footage":
			if i+1 < len(args) {
				squareFootage := NewSquareFootage(args[i+1])
				f, err := squareFootage.GetFilter()
				if err != nil {
					return nil, err
				}

				filters = append(filters, f)
			}
		case "--rooms":
			if i+1 < len(args) {
				rooms := NewRooms(args[i+1])
				f, err := rooms.GetFilter()
				if err != nil {
					return nil, err
				}

				filters = append(filters, f)
			}
		case "--name":
			if i+1 < len(args) {
				filters = append(filters, NewFilter(Contains(args[i+1])))
			}
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

func NewPrice(args string) *Arguments {
	ops := strings.Split(args, ":")
	price, err := strconv.ParseFloat(ops[1], 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = FloatValue(price, EqualPrice)
	evals[lessThan] = FloatValue(price, LessThanPrice)
	evals[greaterThan] = FloatValue(price, GreaterThanPrice)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
}

func NewSquareFootage(args string) *Arguments {
	ops := strings.Split(args, ":")
	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = IntValue(size, EqualFootage)
	evals[lessThan] = IntValue(size, LessThanFootage)
	evals[greaterThan] = IntValue(size, GreaterThanFootage)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
}

func NewRooms(args string) *Arguments {
	ops := strings.Split(args, ":")
	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = IntValue(size, EqualRoom)
	evals[lessThan] = IntValue(size, LessThanRoom)
	evals[greaterThan] = IntValue(size, GreaterThanRoom)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
}

func (a *Arguments) GetFilter() (Filter, error) {
	operation, ok := a.evals[a.operator]
	if !ok {
		return nil, errors.New("operator not supported")
	}

	return NewFilter(operation), nil
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
