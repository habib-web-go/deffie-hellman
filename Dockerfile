FROM golang:latest
WORKDIR /app
RUN apt install -y protobuf-compiler
# https://stackoverflow.com/questions/66371020/how-to-use-docker-to-generate-grpc-code-based-on-go-mod-versions
COPY ./src .
RUN go build -o server .
CMD ["./server"]