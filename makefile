CC := go
OPTION := build
SRC := .
BIN := gevil

compile:
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean: clean_binary clean_docs clean_dir

clean_binary:
	-rm $(BIN)

clean_docs:
	-rm -R docs

clean_dir:
	-rm -R output

dependencies:
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

docs: clean_docs
	bash tools/generate_documentation.sh