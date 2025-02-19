package command

import (
	"strings"

	"github.com/properfilter/src/model"
)

func Contains(n string) func(model.Property) bool {
	return func(p model.Property) bool {
		return strings.Contains(n, p.Name)
	}
}
