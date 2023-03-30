client:
	go build -o build/client cmd/client/client.go
server:
	go build -o build/server cmd/server/server.go

all: client server

run:
	./build/cmd/client
	./build/cmd/server
	

