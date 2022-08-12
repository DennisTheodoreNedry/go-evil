CC := go
OPTION := build
SRC := ./go/*.go
BIN := gevil

compile:
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean:
	-rm $(BIN)

dependencies:
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler