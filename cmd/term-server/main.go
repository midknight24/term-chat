package main

import (
	"flag"
	"fmt"

	chat "github.com/midknight24/term-chat/internal"
)

var host string

func init() {
	flag.StringVar(&host, "host", "localhost:4999", "server address")
}

func main() {
	flag.Parse()

	server := chat.NewServer(host)

	fmt.Println("starting grpc server...")
	server.Run()
}
