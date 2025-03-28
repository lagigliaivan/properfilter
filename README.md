# I did this application as part of a recruiting process.
Unfortunatelly, even though I sent this before the expected deadline, on week later they told me that
another candidate was almost at the end of the process, so they don't even review my deliverable.

I hope this may help others to have some help on some similar exercise. 

# The properfilter command
Command Line to filter a large set of real estate properties based on their particular attributes.
This application was coded in Golang 1.22.2, so to run it you need to have golang runtime installed.

# Golang Installation
You can dowloand the golang runtime from the follwoing link: https://go.dev/dl/

# Project setup instructions
Once golang is installed, you can use the following commands to see some examples:

```
# make help 
```
The command above will perform the following steps:
  - install needed dependencies
  - build application
  - show command help and examples

In case you want to run UTs and check their coverage you may run

```
# make  run
```

# Notes
This application is distributed with a dataset which contains non real properties information, it is just for testing purposes.

# Examples:
Use -f parameter to specify a csv file or use < symbol to read from stdin

```bash
#properties where price is in the range of [10000 and 14000]
./properfilter --price gt:9999 --price lt:14001 < dataset.csv

#properties where price is greater than 10000 and square footage greater than 1000
./properfilter --price gt:10000 --squarefootage gt:1000 < dataset.csv

#properties at less than 100km from lat,long reference point
./properfilter --distance lt:100,-33.013270,-64.430154 < dataset.csv

#properties that have fireplace AND patio
./properfilter --ammenities grill,garage < dataset.csv

#properties that have fireplace OR patio
./properfilter --ammenities "grill|patio" -f dataset.csv

#properties that are described as Compact OR Classic
./properfilter --description "Compact|Classic" -f dataset.csv

#properties with medium and high lighting
./properfilter --lighting gt:low -f dataset.csv

#properties with high lighting
./properfilter --lighting eq:high -f dataset.csv
```

## Author
- [Lagiglia Ivan](https://github.com/lagigliaivan)
- Email (lagigliaivan@gmail.com)
