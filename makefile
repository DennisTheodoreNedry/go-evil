CC := go
OPTION := build
SRC := .
BIN := Malware_Language

compile:
	mkdir output
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean:
	rm $(BIN)