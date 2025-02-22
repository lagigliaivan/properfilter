test:
		go test ./...
testsum:
		gotestsum ./...
run: build
		./properfilter --help
build:
		go build -o properfilter ./cmd/main.go
