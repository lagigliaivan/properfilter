test:
		go test ./...
testsum:
		gotestsum ./...
run:
	go run ./cmd/main.go --price gt=10000 < dataset.csv
build:
		go build ./cmd/main.go
