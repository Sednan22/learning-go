.DEFAULT_GOAL := build

BINARY_FILE=binary-go

.PHONY:fmt vet build clean
fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o $(BINARY_FILE)

run:
	@./$(BINARY_FILE)

all: build run
	
clean:
	go clean