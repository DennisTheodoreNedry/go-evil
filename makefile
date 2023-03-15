CC := go
OPTION := build
SRC := .
BIN := gevil
EXT := ./tools/vscode_ext/evil
VER := 2.1

compile:
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean: clean_binary clean_docs

clean_binary:
	-rm $(BIN)

clean_docs:
	-rm -R docs

dependencies:
	go get github.com/s9rA16Bf4/ArgumentParser
	go get github.com/s9rA16Bf4/notify_handler
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	go get github.com/cloudfoundry/jibber_jabber
	go get github.com/webview/webview
	go install mvdan.cc/garble@latest
	go get github.com/google/gops
	go get github.com/thanhpk/randstr

docs: clean_docs
	bash tools/generate_documentation.sh

install_ext:
	cp -R $(EXT) ~/.vscode/extensions/

uninstall_ext:
	-rm -R ~/.vscode/extensions/evil

update_ext: uninstall_ext install_ext

docker:
	docker build -t goevil/gevil-$(VER) .