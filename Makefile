MODULE=passwordgenerator

.PHONY: default
default: build

.PHONY: clearbin
clearbin:
	rm -rf ./bin

build: cmd/main.go clearbin
	GOOS=linux GOARCH=amd64 go build -v -o ./bin/$(MODULE)-linux-amd64 ./cmd/main.go
	GOOS=windows GOARCH=amd64 go build -v -o ./bin/$(MODULE)-windows-amd64.exe ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -v -o ./bin/$(MODULE)-darwin-amd64 ./cmd/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/$(MODULE)-alpine-linux-amd64 ./cmd/main.go

.PHONY: cross
cross: cmd/main.go clearbin
	go build -v -o ./bin/$(MODULE) ./main.go

.PHONY: run
run: cross
	./bin/$(MODULE)