PROJ=hello-shapes-go
SRC=main.go
SPECS_DIR=./shapes
BIN=hello-shapes-go

help:
	@@echo make clean
	@@echo make build
	@@echo make test
	@@echo make specs	
	@@echo make run
	@@echo make all

clean:
	rm $(BIN)
	
build: $(SRC)
	go build -o $(BIN)

test:
	go test -v ./...

specs:
	ginkgo $(SPECS_DIR)

run: build
	go run $(BIN)

all: run
