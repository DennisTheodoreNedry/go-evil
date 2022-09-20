CC := go
OPTION := build
SRC := .
BIN := gevil

compile:
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean: clean_binary clean_docs clean_output

clean_binary:
	-rm $(BIN)

clean_docs:
	-rm -R docs

clean_output:
	-rm -R output

dependencies:
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	github.com/cloudfoundry/jibber_jabber
	go get github.com/webview/webview
	go install mvdan.cc/garble@latest

docs: clean_docs
	bash tools/generate_documentation.sh