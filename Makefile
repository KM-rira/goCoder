# Makefile

.PHONY: test

all: test

test:
	go test -v question_test.go

