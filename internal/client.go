package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	chat "github.com/midknight24/term-chat/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type client struct {
	username string
	remote   string
	token    string
	chat.ChatRoomClient
}

func NewClient(remote, username string) *client {
	return &client{
		remote:   remote,
		username: username,
	}
}

func (c *client) JoinChat(ctx context.Context) error {
	tkn, err := c.ChatRoomClient.JoinChat(ctx, &chat.LoginRequest{Username: c.username})
	if err != nil {
		return err
	}
	c.token = tkn.GetToken()
	fmt.Println(c.token)
	return nil
}

func (c *client) StreamChats(ctx context.Context, input, output chan string) {
	md := metadata.New(map[string]string{
		tokenHeader: c.token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)
	streamClient, err := c.ChatRoomClient.Stream(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer streamClient.CloseSend()
	go c.sendLoop(streamClient, input)
	c.recvLoop(streamClient, output)
}

func (c *client) sendLoop(sc chat.ChatRoom_StreamClient, input chan string) {
Loop:
	for {
		select {
		case msg, ok := <-input:
			if !ok {
				break Loop
			}
			sc.Send(&chat.StreamMessage{
				Timestamp: timestamppb.Now(),
				Message: &chat.StreamMessage_ClientMessage{
					ClientMessage: &chat.ChatMessage{
						Timestamp: timestamppb.Now(),
						Username:  c.username,
						Content:   msg,
					},
				},
			})
		case <-sc.Context().Done():
			break Loop
		}
	}
}

func (c *client) recvLoop(sc chat.ChatRoom_StreamClient, output chan string) {
	for {
		msg, err := sc.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			fmt.Println("stream cancelled")
			return
		} else if err == io.EOF {
			fmt.Println("stream closed by server")
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}
		output <- formatClientMessage(msg)
	}
}

func formatClientMessage(msg *chat.StreamMessage) string {
	switch m := msg.Message.(type) {
	case *chat.StreamMessage_ClientMessage:
		username := m.ClientMessage.Username
		content := m.ClientMessage.Content
		return fmt.Sprintf("[%s]: %s", username, content)
	case *chat.StreamMessage_ClientLogin:
		username := "server"
		content := fmt.Sprintf("user [%v] has entered the room", m.ClientLogin.Username)
		return fmt.Sprintf("[%s]: %s", username, content)
	case *chat.StreamMessage_ClientLogout:
		username := "server"
		content := fmt.Sprintf("user [%v] has left the room", m.ClientLogout.Username)
		return fmt.Sprintf("[%s]: %s", username, content)
	default:
		return "unsupported message"
	}
}

func (c *client) Run(ctx context.Context, input, output chan string) error {
	context, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(context, c.remote, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c.ChatRoomClient = chat.NewChatRoomClient(conn)
	err = c.JoinChat(ctx)
	if err != nil {
		return err
	}
	c.StreamChats(ctx, input, output)
	return nil
}
