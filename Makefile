.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*

test:
	go test -v ./...

test-bench:
	go test ./... -bench=.