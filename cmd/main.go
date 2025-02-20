package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/properfilter/src/command"
	"github.com/properfilter/src/model"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("Usage: properfilter [command] [arguments]\n")
		return
	}

	cmd, err := command.Parse(args)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	properties := make(model.Properties, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		property, err := command.CsvToProperty(s.Text())
		if err != nil {
			log.Printf("Error: %s\n", err)
			continue
		}

		properties = append(properties, *property)
	}

	props := cmd.Execute(context.Background(), properties)

	for _, p := range props {
		fmt.Printf("%s,%f,%d,%d,%s\n", p.Name, p.Price, p.Bathrooms, p.Bathrooms, p.Ammenities)
	}
}
