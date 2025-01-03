build: main.go ./cmd/*.go ./db/*.go
	CGO_ENABLED=1 go build -ldflags "-w -s -X todo-go/cmd.version=$(shell git describe --tags --always --dirty) -linkmode external -extldflags -static" -trimpath -o todo-go ./

clean:
	rm -f todo-go

test:
	go test ./...

run:
	./todo-go

.PHONY: build clean test run
