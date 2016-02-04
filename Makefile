.PHONY: clean install uninstall

all:
	go build hello-server.go

install: all
	mkdir -p ~/bin
	cp hello-server ~/bin

clean:
	rm hello-server

uninstall:
	rm ~/bin/hello-server
