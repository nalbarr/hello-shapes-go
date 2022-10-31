PROJ=hello-shapes-go
SRC=main.go
BIN=hello-shapes-go

all: run

clean:
	rm $(BIN)
	
build: $(SRC)
	# go build -o $(BIN)
	go build

test:
	go test -v ./...

run: build
	go run $(BIN)
