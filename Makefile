.PHONY: check

default: check

check:
	golangci-lint run