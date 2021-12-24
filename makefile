CC := go
OPTION := build
SRC := .
BIN := gevil

compile:
	mkdir -p output

	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean:
	rm -R output
	rm $(BIN)