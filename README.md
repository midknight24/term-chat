# term-chat
simple grpc based terminal chatroom app

## requirement
- go1.19
- protoc, version3
- google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

## usage
```shell
make install
```

client command:

```shell
term-client -remote ${server_address}
```

server command:

```shell
term-server -host ${bind_address}
```
