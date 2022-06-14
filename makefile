CC := go
OPTION := build
SRC := .
BIN := gevil

compile: create_directory
	@-echo "## Compiling"
	$(CC) $(OPTION) -o $(BIN) $(SRC)

create_directory:
	mkdir -p output

clean: clean_output clean_binary clean_examples_list

clean_output:
	-rm -R output/*

clean_binary:
	-rm $(BIN)

clean_examples_list:
	-rm examples/examples.txt

update_examples: compile
	python3 examples/update_compiler_version.py

test:
	@-echo "## Checking if builtin functionallity is working as expected"
	cd ./utility/contains/ && go test
	cd ./utility/converter/ && go test
	cd ./utility/algorithm/encryption/aes/ && go test
	cd ./utility/algorithm/encryption/rsa/ && go test
	cd ./utility/algorithm/path/ && go test

install_dependencies:
	go get github.com/thanhpk/randstr
	go get github.com/webview/webview
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go get golang.org/x/crypto
	go get golang.org/x/sys
	go get gopkg.in/go-rillas/subprocess.v1
	go get github.com/TwiN/go-pastebin
	go get github.com/cloudfoundry/jibber_jabber
	go mod download golang.org/x/net
