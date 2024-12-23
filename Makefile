build: main.go ./cmd/*.go ./db/*.go
	CGO_ENABLED=0 go build -trimpath -o todo-go ./

clean:
	rm -f todo-go

test:
	go test ./...

run:
	./todo-go

.PHONY: build clean test run
