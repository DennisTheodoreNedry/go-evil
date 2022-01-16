CC := go
OPTION := build
SRC := .
BIN := gevil

compile: create_directory
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

generate_example_list:
	python3 examples/generate_list_of_examples.py

update_examples: compile
	python3 examples/update_compiler_version.py


test: compile generate_example_list
	bash run_tests.bash

	@-rm examples/attack_vector/encryption/extension/target_folder/*_encrypted
	@-rm examples/attack_vector/encryption/extension/target_folder/*_decrypted
	@-rm examples/attack_vector/encryption/file/*_encrypted
	@-rm examples/attack_vector/encryption/file/*_decrypted
	@-rm examples/attack_vector/encryption/folder/testfolder/*_encrypted
	@-rm examples/attack_vector/encryption/folder/testfolder/*_decrypted

check: test

install_dependencies:
	go get github.com/thanhpk/randstr
	go get github.com/webview/webview
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go get golang.org/x/crypto
	go get golang.org/x/sys
	go get gopkg.in/go-rillas/subprocess.v1

