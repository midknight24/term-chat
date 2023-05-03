package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os/signal"

	"log"
	"os"

	chat "github.com/midknight24/term-chat/internal"
)

var serverAddr string

func init() {
	flag.StringVar(&serverAddr, "remote", "localhost:4999", "server address")
}

func scanInput(c chan string) {
	sc := bufio.NewScanner(os.Stdin)
	defer close(c)
	for {
		if sc.Scan() {
			text := sc.Text()
			if text != "" {
				c <- text
			}
		} else {
			break
		}
	}
}

func printOutput(c chan string) {
	for text := range c {
		fmt.Println(text)
	}
}

func main() {
	flag.Parse()

	fmt.Println("Enter your username: ")
	var username string = "dummy"
	fmt.Scanln(&username)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)

	input := make(chan string)
	output := make(chan string)
	go scanInput(input)
	go printOutput(output)

	client := chat.NewClient(serverAddr, username)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go gracefulShutdown(cancel, quit)

	err := client.Run(ctx, input, output)

	if err != nil {
		log.Fatal(err)
	}
}

func gracefulShutdown(cancel context.CancelFunc, c chan os.Signal) {
	<-c
	fmt.Println("Gracefully shutting down client...")
	cancel()
}
