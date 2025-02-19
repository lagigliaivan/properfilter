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
		evals    map[string]func(v interface{}) func(model.Property) bool
		value    interface{}
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

func NewPrice(args string) *Arguments {
	ops := strings.Split(args, ":")
	price, err := strconv.ParseFloat(ops[1], 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(v interface{}) func(model.Property) bool)
	evals[equal] = EqualPrice
	evals[lessThan] = LessThanPrice
	evals[greaterThan] = GreaterThanPrice

	return &Arguments{
		evals:    evals,
		value:    price,
		operator: ops[0],
	}
}

func NewSquareFootage(args string) *Arguments {
	ops := strings.Split(args, ":")
	size, err := strconv.ParseInt(ops[1], 10, 32)
	if err != nil {
		return nil
	}

	evals := make(map[string]func(v interface{}) func(model.Property) bool)
	evals[equal] = EqualFootage
	evals[lessThan] = LessThanFootage
	evals[greaterThan] = GreaterThanFootage

	return &Arguments{
		evals:    evals,
		value:    size,
		operator: ops[0],
	}
}

func (a *Arguments) GetFilter() (Filter, error) {
	operation, ok := a.evals[a.operator]
	if !ok {
		return nil, errors.New("operator not supported")
	}

	return NewFilter(operation(a.value)), nil
}
