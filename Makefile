run-api:
	go run zero-app.go
api-generate:
	goctl api go -api zero-app.api -dir .
docs-generate:
	rm -rf zero-app.json
	goctl api plugin -plugin goctl-swagger="swagger -filename zero-app.json -host 192.168.2.34 basepath /" -api zero-app.api -dir .
run-docs:
	docker stop zero-docs
	docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/zero-app.json -v $PWD:/foo -d --name zero-docs swaggerapi/swagger-ui
generate-proto:
	goctl rpc template -o auth.proto
generate-rpc:
	goctl rpc protoc auth.proto --go_out=. --go-grpc_out=. --zrpc_out=.
