run-api:
	go run zero-app.go
api-generate:
	goctl api go -api zero-app.api -dir . -style goZero
docs-generate:
	rm -rf zero-app.json
	goctl api plugin -plugin goctl-swagger="swagger -filename zero-app.json -host 192.168.2.34 basepath /" -api zero-app.api -dir .
run-docs:
	docker stop zero-docs
	docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/zero-app.json -v $PWD:/foo -d --name zero-docs swaggerapi/swagger-ui
generate-proto:
	goctl rpc template -o auth.proto
generate-rpc:
	goctl rpc protoc auth.proto --go_out=. --go-grpc_out=. --zrpc_out=.  -style goZero
run-etcd:
	docker run -e ALLOW_NONE_AUTHENTICATION=yes -p 2379:2379 -p 2380:2380 --rm -it --name Etcd -d bitnami/etcd:3.5.9
generate-model:
	 goctl model pg datasource -url="postgres://user:password@127.0.0.1:5454/database?sslmode=disable" --schema="public" -style go_zero --table="test_users" --dir ./model
validate:
	goctl api validate -api zero-app.api
ts-doc:
	goctl api ts --api zero-app.api --dir ./ts --unwrap axiosutil --caller axiosutil