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

.PHONY: test
test:
	go test -v ./...

.PHONY: test-bench
test-bench:
	go test ./... -bench=.

.PHONY: install-loadtesting-deps
install-loadtesting-deps:
	npm i -g artillery

.PHONY: run-loadtesting
run-loadtesting:
	artillery run websocket_loadtesting.yml