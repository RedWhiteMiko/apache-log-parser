SHELL:=/bin/bash

test:
	GOPATH=$(shell pwd) go test def/...

build: test
	mkdir -p bin/
	GOPATH=$(shell pwd) go build -o bin/apache_test src/main.go
	cp bin/apache_test .

run: build
	./bin/apache_test