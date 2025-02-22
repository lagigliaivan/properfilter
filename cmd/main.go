package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/properfilter/src/command"
)

func main() {
	args := os.Args[1:] //removing the first argument which is the program name

	if len(args) == 0 {
		help()
		return
	}

	if exist("--help", args) {
		help()
		return
	}

	cmd, err := command.New(args)
	if err != nil {
		log.Printf("Error: %s\n", err)
		help()
		return
	}

	var r io.Reader
	if exist("-f", args) {
		fileName := find("-f", args)
		f, e := os.Open(fileName)
		if e != nil {
			log.Printf("Error: %s\n", err)
			help()
			return
		}

		defer f.Close()
		r = f

	} else {
		r = os.Stdin
	}

	scanInput(func(line string) {
		property, err := command.CsvToProperty(line)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		p := cmd.Filter(context.Background(), *property)
		if p != nil {
			fmt.Printf("%s \n", p.String())
		}
	}, r)
}

func scanInput(run func(string), r io.Reader) error {
	s := bufio.NewScanner(r)
	csvHeader := 0
	for s.Scan() {
		if csvHeader == 0 {
			csvHeader++
			continue
		}

		run(s.Text())
	}

	return nil
}

func exist(value string, args []string) bool {
	for _, arg := range args {
		if arg == value {
			return true
		}
	}

	return false
}

func find(value string, args []string) string {
	for i, arg := range args {
		if arg == value {
			if len(args) <= i+1 {
				return ""
			}
			return args[i+1]
		}
	}

	return ""
}
func help() {
	color.Green(` ______   ______     ______     ______   ______     ______     ______   __     __         ______   ______     ______    
/\  == \ /\  == \   /\  __ \   /\  == \ /\  ___\   /\  == \   /\  ___\ /\ \   /\ \       /\__  _\ /\  ___\   /\  == \   
\ \  _-/ \ \  __<   \ \ \/\ \  \ \  _-/ \ \  __\   \ \  __<   \ \  __\ \ \ \  \ \ \____  \/_/\ \/ \ \  __\   \ \  __<   
 \ \_\    \ \_\ \_\  \ \_____\  \ \_\    \ \_____\  \ \_\ \_\  \ \_\    \ \_\  \ \_____\    \ \_\  \ \_____\  \ \_\ \_\ 
  \/_/     \/_/ /_/   \/_____/   \/_/     \/_____/   \/_/ /_/   \/_/     \/_/   \/_____/     \/_/   \/_____/   \/_/ /_/ 
                                                                                                                        `)
	fmt.Print("Usage:\n properfilter [-f] <filename> <filter> [operator]:<value>\n\n")
	fmt.Print("Filters:\n")
	fmt.Print("--address <value[,|]...>\n")
	fmt.Print("--squarefootage  <eq|lt|gt>:<value>\n")
	fmt.Print("--lighting  <eq|lt|gt>:<low|medium|high>\n")
	fmt.Print("--price  <eq|lt|gt>:<value>\n")
	fmt.Print("--rooms  <eq|lt|gt>:<value>\n")
	fmt.Print("--bathrooms  <eq|lt|gt>:<value>\n")
	fmt.Print("--description <value[,|]...>\n")
	fmt.Print("--ammenities <value[,|]...>\n")
	fmt.Print("--distance <lt|gt><distance:lat:long>\n\n")

	fmt.Print("!!!Some filters support OR/AND operations, where | means OR and , means AND\n\n")

	fmt.Print("Examples:\n")
	color.Cyan("#properties where price is in the range of [10000 and 14000]\n")
	fmt.Print("./properfilter --price gt:9999 --price lt:14001< dataset.csv\n\n")

	color.Cyan("#properties where price is greater than 10000 and square footage greater than 90 \n")
	fmt.Print("./properfilter --price gt:10000 --squarefootage gt:90 < dataset.csv\n\n")

	color.Cyan("#properties at less than 10km from lat,long reference point\n")
	fmt.Print("./properfilter --distance lt:10,-33.90,71.700000 < dataset.csv\n\n")

	color.Cyan("#properties at more than 10km from lat,long reference point\n")
	fmt.Print("./properfilter --distance gt:10,-33.90,71.700000 < dataset.csv\n\n")

	color.Cyan("#properties that have fireplace AND garage\n")
	fmt.Print("./properfilter --ammenities grill,garage < dataset.csv\n\n")

	color.Cyan("#properties that have grill OR patio\n")
	fmt.Print("./properfilter --ammenities \"grill|patio\" -f dataset.csv\n\n")

	color.Cyan("#properties that contains the word Spruce in their address\n")
	fmt.Print("./properfilter --address Spruce -f dataset.csv\n\n")

	color.Cyan("#properties that are described as Compact OR Classic\n")
	fmt.Print("./properfilter --description \"Compact|Classic\" -f dataset.csv\n\n")

	color.Cyan("#properties with medium and high lighting\n")
	fmt.Print("./properfilter --lighting gt:low -f dataset.csv\n\n")

	color.Cyan("#properties with high lighting\n")
	fmt.Print("./properfilter --lighting eq:high -f dataset.csv\n")
}
