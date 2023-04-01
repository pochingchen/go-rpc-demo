demo:
	go build -o build/client cmd/demo/client.go
	go build -o build/server cmd/demo/server.go

json:
	go build -o build/client-json cmd/demojson/client.go
	go build -o build/server-json cmd/demojson/server.go

goods:
	go build -o build/goods-server cmd/goods/goods_server.go
	go build -o build/goods-client cmd/goods/goods_client.go



all: demo json goods


