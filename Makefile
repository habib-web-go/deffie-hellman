generate-proto:
	protoc --go_out=./gen --go_opt=paths=source_relative \
  --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
  grpc/authpb.proto

	protoc -I . --grpc-gateway_out ./gen \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
  --grpc-gateway_opt generate_unbound_methods=true \
  grpc/authpb.proto

generate-swagger-json:
	protoc -I=. --go_out=. --go-grpc_out=. --swagger_out=logtostderr=true:. grpc/authpb.proto

generate-swagger-server:
	docker pull quay.io/goswagger/swagger
	alias goswagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
	goswagger generate server -f ./authpb.swagger.json

clean-generated:
	rm gen/grpc/*.go

generate-google-annotations:
	mkdir -p google/api
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto