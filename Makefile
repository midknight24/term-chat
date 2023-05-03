.PHONY: all
all: client server

.PHONY: client
client: protos
	go install ./cmd/term-client

.PHONY: server
server: protos
	go install ./cmd/term-server


.PHONY: protos
protos: protos/chatroom.proto
	protoc --go_opt=paths=source_relative --go_out="." --go-grpc_out="." \
	--go-grpc_opt=paths=source_relative .\protos\chatroom.proto


.PHONY: uninstall
uninstall:
	rm -f ${GOPATH}/bin/term-client*
	rm -f ${GOPATH}/bin/term-server*