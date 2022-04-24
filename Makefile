# GOBASE=$(shell pwd)
# GOBIN=$(GOBASE)/bin

.PHONY: start
start: bin/modd
	bin/modd

.PHONY: restart
restart: clean build

.PHONY: clean
clean:
	rm -f bin/notes

.PHONY: build
build: bin/notes
bin/notes:
	go build -o bin/notes ./cmd/notes

bin/modd:
	go build -o bin/modd github.com/cortesi/modd/cmd/modd