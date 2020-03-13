default: build

all: deps build

deps:
	@go get -d -v github.com/gorilla/mux

build:
	@go fmt
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .

clean:
	rm service

.PHONY: build
