package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/properfilter/src/command"
)

func main() {
	args := os.Args[1:] //removing the first argument which is the program name

	if len(args) == 0 {
		fmt.Print("Usage: properfilter [command] [arguments]\n")
		return
	}

	cmd, err := command.New(args)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	scanStdin(func(line string) {
		property, err := command.CsvToProperty(line)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		p := cmd.Filter(context.Background(), *property)
		if p != nil {
			fmt.Println(p.String())
		}
	})
}

func scanStdin(run func(string)) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		run(s.Text())
	}
}
