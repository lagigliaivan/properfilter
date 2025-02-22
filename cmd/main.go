package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

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
	for s.Scan() {
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
	fmt.Print(` ______   ______     ______     ______   ______     ______     ______   __     __         ______   ______     ______    
/\  == \ /\  == \   /\  __ \   /\  == \ /\  ___\   /\  == \   /\  ___\ /\ \   /\ \       /\__  _\ /\  ___\   /\  == \   
\ \  _-/ \ \  __<   \ \ \/\ \  \ \  _-/ \ \  __\   \ \  __<   \ \  __\ \ \ \  \ \ \____  \/_/\ \/ \ \  __\   \ \  __<   
 \ \_\    \ \_\ \_\  \ \_____\  \ \_\    \ \_____\  \ \_\ \_\  \ \_\    \ \_\  \ \_____\    \ \_\  \ \_____\  \ \_\ \_\ 
  \/_/     \/_/ /_/   \/_____/   \/_/     \/_____/   \/_/ /_/   \/_/     \/_/   \/_____/     \/_/   \/_____/   \/_/ /_/ 
                                                                                                                        `)
	fmt.Print("\n\n")
	fmt.Print("Usage:\n properfilter [-f] <filename> <filter> [operator]:<value>\n\n")
	fmt.Print("Filters:\n")
	fmt.Print("--name <value>\n")
	fmt.Print("--squarefootage  <eq|lt|gt>:<value>\n")
	fmt.Print("--lighting  <eq|lt|gt>:<low|medium|high>\n")
	fmt.Print("--price  <eq|lt|gt>:<value>\n")
	fmt.Print("--rooms  <eq|lt|gt>:<value>\n")
	fmt.Print("--bathrooms  <eq|lt|gt>:<value>\n")
	fmt.Print("--description <value>\n")
	fmt.Print("--ammenities [eq]:<value>\n")
	fmt.Print("--distance <lt|gt><distance:lat:long>\n")

	fmt.Print("Examples:\n")
	fmt.Print("./properfilter --price gt:10000 --price lt:20000< dataset.csv\n")
	fmt.Print("./properfilter --price gt:10000 --squarefootage gt:1000 < dataset.csv\n")
	fmt.Print("./properfilter --distance lt:100,-33.013270,-64.430154 < dataset.csv  ")
	fmt.Print("#properties at less than 100km from lat,long reference point\n")
	fmt.Print("./properfilter --ammenities fireplace,patio < dataset.csv\n")
	fmt.Print("./properfilter --price gt:310000 --ammenities fireplace,patio < dataset.csv\n\n")

	fmt.Print("./properfilter --ammenities fireplace,patio -f dataset.csv\n")
}
