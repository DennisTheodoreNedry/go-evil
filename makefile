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

generate_examples:
	python examples/generate_list_of_examples.py

test: compile generate_examples
	bash run_tests.bash

check: test

install_dependencies:
	go get github.com/thanhpk/randstr
	go get github.com/webview/webview
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go get golang.org/x/crypto
	go get golang.org/x/sys