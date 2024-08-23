.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: run
run: all
	./bin/server

.PHONY: clean
clean:
	go clean
	rm -f bin/*

test:
	go test -v ./...

test-bench:
	go test ./... -bench=.