.PHONY: all build clean test

EXEC_NAME=app
SRC=./...
OUTPUT=./bin

all: build test

build: test
	mkdir -p $(OUTPUT)
	go build -o $(OUTPUT)/$(EXEC_NAME) $(SRC)

clean:
	rm -f $(OUTPUT)/$(EXEC_NAME)
	go clean


test:
	go test
