build: main.go ./cmd/*.go ./db/*.go
	CGO_ENABLED=0 go build -ldflags "-w -s -X todo-go/cmd.version=$(shell git describe --tags --always --dirty)" -trimpath -o todo-go ./

clean:
	rm -f todo-go

test:
	go test ./...

run:
	./todo-go

.PHONY: build clean test run
