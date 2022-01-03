CC := go
OPTION := build
SRC := .
BIN := gevil

compile: create_directory
	$(CC) $(OPTION) -o $(BIN) $(SRC)

create_directory:
	mkdir -p output

clean: clean_output clean_binary

clean_output:
	rm -R output

clean_binary:
	rm $(BIN)
