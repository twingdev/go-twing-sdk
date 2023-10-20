protoc:
	protoc --go_out=. proto/*.proto

@phony: protoc

build: protoc
	go build .

