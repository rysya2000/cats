.PHONY: all run test

all: test run

run:
	go run main.go

test:
	go test -v ./... -count=1
