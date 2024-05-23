CC := go
OPTION := build
SRC := .
BIN := gevil
EXT := ./tools/vscode_ext/evil

default: compile

submodules:
	git submodule init
	git submodule update

compile: compile_domains
	$(CC) $(OPTION) -o $(BIN) $(SRC)

compile_domains:
	bash tools/compile_domains.sh

update_domains:
	cd domains && git pull

update_examples:
	cd examples && git pull

clean: clean_binary

clean_binary:
	-rm $(BIN)

dependencies:
	go get github.com/DennisTheodoreNedry/ArgumentParser
	go get github.com/DennisTheodoreNedry/notify_handler
	go get github.com/cloudfoundry/jibber_jabber
	go get github.com/webview/webview
	go install mvdan.cc/garble@latest
	go get github.com/google/gops
	go get github.com/thanhpk/randstr
	go get github.com/tatsushid/go-fastping
	go get github.com/ARaChn3/gfb
	go get github.com/ARaChn3/puffgo
	go get github.com/redcode-labs/Coldfire
	go get github.com/MarinX/keylogger
	go get github.com/KindlyFire/go-keylogger
	go get github.com/DennisTheodoreNedry/Go-tools
	go get github.com/DennisTheodoreNedry/lorca

install_ext:
	cp -R $(EXT) ~/.vscode/extensions/

uninstall_ext:
	-rm -R ~/.vscode/extensions/evil

update_ext: uninstall_ext install_ext

docker:
	docker build -t goevil/gevil .