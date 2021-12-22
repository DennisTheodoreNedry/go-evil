CC := go
OPTION := build
SRC := .
BIN := Malware_Language

compile:
	mkdir -p output
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean:
	rm $(BIN)