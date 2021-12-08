MODULE=passwordgenerator
MAIN_PATH=./main.go

.PHONY: default
default: build

.PHONY: clearbin
clearbin:
	rm -rf ./bin

build: main.go clearbin
	GOOS=linux GOARCH=amd64 go build -v -o ./bin/$(MODULE)-linux-amd64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -v -o ./bin/$(MODULE)-windows-amd64.exe $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -v -o ./bin/$(MODULE)-darwin-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/$(MODULE)-alpine-linux-amd64 $(MAIN_PATH)

buildzip: main.go
	rm -rf ./dist && mkdir -p ./dist
	
	cd ./dist; \
	\
	GOOS=linux GOARCH=amd64 go build -v -o ./$(MODULE) ../$(MAIN_PATH); \
	zip -r ./$(MODULE)-$(VERSION)-linux-amd64.zip ./$(MODULE); \
	rm -f ./$(MODULE); \
	\
	GOOS=windows GOARCH=amd64 go build -v -o ./$(MODULE).exe ../$(MAIN_PATH); \
	zip -r ./$(MODULE)-$(VERSION)-windows-amd64.zip ./$(MODULE).exe; \
	rm -f ./$(MODULE).exe; \
	\
	GOOS=darwin GOARCH=amd64 go build -v -o ./$(MODULE) ../$(MAIN_PATH); \
	zip -r ./$(MODULE)-$(VERSION)-darwin-amd64.zip ./$(MODULE); \
	rm -f ./$(MODULE); \
	\
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./$(MODULE) ../$(MAIN_PATH); \
	zip -r ./$(MODULE)-$(VERSION)-alpine-linux-amd64.zip ./$(MODULE); \
	rm -f ./$(MODULE);

.PHONY: cross
cross: main.go clearbin
	go build -v -o ./bin/$(MODULE) $(MAIN_PATH)

.PHONY: run
run: cross
	./bin/$(MODULE)