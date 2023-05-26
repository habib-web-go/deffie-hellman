FROM golang:latest
WORKDIR /app
RUN apt install -y protobuf-compiler
# https://stackoverflow.com/questions/66371020/how-to-use-docker-to-generate-grpc-code-based-on-go-mod-versions
RUN GO111MODULE=on \
        go get google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
COPY ./src .
RUN protoc --go_out=. --go-grpc_out=. authpb.proto
RUN go build -o server .
CMD ["./server"]