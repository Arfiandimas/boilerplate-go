.PHONY: tool clean build install

tool:
	go build -o bin/orn tools/main.go

clean:
	rm -rf ./bin

build:
	env CGO_ENABLED=0 go build -a -o bin/main src/main.go

test:
	go test ./src/app/... -coverprofile=coverage.out