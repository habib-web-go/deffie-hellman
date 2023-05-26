## Install proto-compiler
### mac
```sh
brew install protobuf
```
### linux
```sh
apt install -y protobuf-compiler
```
## Compile Proto
```
protoc --go_out=. --go-grpc_out=.  grpc/authpb.proto
```
