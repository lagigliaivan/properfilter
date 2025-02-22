# properfilter command
Command Line to filter a large set of real estate properties based on their particular attributes.
This application was coded in Golang 1.22.2, so to run it you need to have it installed.

# Installation
You can dowloand the golang image from the follwoing link: https://go.dev/dl/

Once golang is installed, you can use the following commands to see some examples:

```
# make run
```
The command above will show you the command help with some examples that are self explanatory.

## Notes:

OR operation is not supported. All filters will be act as an AND operation.

# Examples:
Examples:
```bash
/properfilter --price gt:10000 --price lt:20000< dataset.csv
```
```bash
#properties which price is greater than 10000
./properfilter --price gt:10000 --squarefootage gt:1000 < dataset.csv
```
```bash
#properties at less than 100km from lat,long reference point
./properfilter --distance lt:100,-33.013270,-64.430154 < dataset.csv  
```
```bash
#properties which have ammenities such as fireplace and pation
./properfilter --ammenities fireplace,patio < dataset.csv
```