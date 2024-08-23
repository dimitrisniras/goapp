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

.PHONY: run-concurrent
run-concurrent: all
	@$(eval connections ?= 1)
	./bin/client -n $(connections)

.PHONY: clean
clean:
	go clean
	rm -f bin/*

test:
	go test -v ./...

test-bench:
	go test ./... -bench=.