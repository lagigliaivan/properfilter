package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func NewAmmenities(args string) *Arguments {
	ops := strings.Split(args, ":")
	if len(ops) != 2 {
		return nil
	}

	ammenities := ops[1]

	evals := make(map[string]func(model.Property) bool)
	evals[equal] = StringValue(ammenities, EqualAmmenities)

	return &Arguments{
		evals:    evals,
		operator: ops[0],
	}
}

func EqualAmmenities(p model.Property, v string) bool {
	for _, a := range p.Ammenities {
		if a == v {
			return true
		}
	}

	return false
}
