# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=admin
BINARY_MAC=$(BINARY_NAME)_mac
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME)_win.exe


all: run

default:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-mac:
	$(GOBUILD) -o $(BINARY_MAC) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WIN)
	rm -f $(BINARY_MAC)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v 
	./$(BINARY_NAME)



