# The properfilter command
Command Line to filter a large set of real estate properties based on their particular attributes.
This application was coded in Golang 1.22.2, so to run it you need to have golang runtime installed.

# Golang Installation
You can dowloand the golang runtime from the follwoing link: https://go.dev/dl/

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

#properties where price greater than 10000 and square footage greater than 1000
./properfilter --price gt:10000 --squarefootage gt:1000 < dataset.csv

#properties at less than 100km from lat,long reference point
./properfilter --distance lt:100,-33.013270,-64.430154 < dataset.csv

#properties that have fireplace AND patio
./properfilter --ammenities fireplace,patio < dataset.csv

#properties that have fireplace OR patio
./properfilter --ammenities "fireplace|patio" -f dataset.csv

#properties that are described as Compact OR Classic
./properfilter --description "Compact|Classic" -f dataset.csv
```