PROJ=hello-shapes-go
SRC=main.go
BIN=hello-shapes-go

all: run

clean:
	rm $(BIN)
	
build: $(SRC)
	go build -o $(BIN)

run: build
	go run $(BIN)
