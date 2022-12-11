icon:
	rsrc -ico go.ico

build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go run .

compile:
	go mod tidy
	go clean
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/reverse-linux-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/reverse-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/reverse-windows-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/reverse-macos-amd64 .

all: icon build
