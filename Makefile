demo:
	go build -o build/client cmd/demo/client.go
	go build -o build/server cmd/demo/server.go

json:
	go build -o build/client-json cmd/demojson/client.go
	go build -o build/server-json cmd/demojson/server.go

goods:
	go build -o build/goods-server cmd/demogoods/goods_server.go
	go build -o build/goods-client cmd/demogoods/goods_client.go

grpc:
	go build -o build/greeter-server cmd/demogrpc/greeter/server.go
	go build -o build/greeter-client cmd/demogrpc/greeter/client.go


all: demo json goods grpc


