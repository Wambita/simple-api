.PHONY: run test format clean

run:
	go run main.go

test:
	go test -v

format:
	go fmt ./...

clean:
	rm -rf ./bin
