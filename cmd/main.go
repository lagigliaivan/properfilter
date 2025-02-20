package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/properfilter/src/command"
	"github.com/properfilter/src/model"
)

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
	},
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("Usage: properfilter [command] [arguments]\n")
		return
	}

	cmd, err := command.Parse(args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	props := cmd.Execute(context.Background(), dataSet)

	fmt.Print(props)
}
