SRCS := $(shell find . -type f -name '*.go')

bin/log: $(SRCS) go.mod go.sum
	go build -o $@ ./cmd/log
