package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	chat "github.com/midknight24/term-chat/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const tokenHeader = "x-user-token"

type ChatServer struct {
	Addr string
	// Messages sent to this channel will be broadcasted to all clients
	Broadcast            chan *chat.StreamMessage
	ClientNames          map[string]string
	ClientStreams        map[string]chan *chat.StreamMessage
	namesMtx, streamsMtx sync.RWMutex
	chat.UnimplementedChatRoomServer
}

func NewServer(host string) *ChatServer {
	return &ChatServer{
		Addr:          host,
		Broadcast:     make(chan *chat.StreamMessage, 5000),
		ClientNames:   make(map[string]string),
		ClientStreams: make(map[string]chan *chat.StreamMessage),
	}
}

func (srv *ChatServer) Run() {
	lis, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chat.RegisterChatRoomServer(grpcServer, srv)

	go srv.broadcast()

	fmt.Println("serving")
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}

func (srv *ChatServer) getUsername(id string) (string, bool) {
	srv.namesMtx.RLock()
	defer srv.namesMtx.RUnlock()
	if name, ok := srv.ClientNames[id]; ok {
		return name, true
	}
	return "", false
}

// sendBroadcasts sends broadcast messages received from indivdiual stream channel to remote client
func (srv *ChatServer) sendBroadcasts(ssrv chat.ChatRoom_StreamServer, id string) {
	username, _ := srv.getUsername(id)
	for msg := range srv.ClientStreams[id] {
		if s, ok := status.FromError(ssrv.Send(msg)); ok {
			switch s.Code() {
			case codes.OK:
			case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
				log.Printf("client [%s] terminated\n", username)
			default:
				log.Printf("failed to send to client [%s], error: %v", username, s.Err())
			}
		}
	}

}

// broadcast receiveds broadcast messages from Broadcast channel and sends to client stream channels
func (srv *ChatServer) broadcast() {
	for msg := range srv.Broadcast {
		srv.streamsMtx.RLock()
		for _, stream := range srv.ClientStreams {
			select {
			case stream <- msg:
			default:
			}
		}
		srv.streamsMtx.RUnlock()
	}
}

// JoinChat receives a *protos.LoginRequest and returns a *protos.ChatToken for later chat streaming
func (srv *ChatServer) JoinChat(_ context.Context, req *chat.LoginRequest) (*chat.ChatToken, error) {
	id := uuid.New().String()
	fmt.Println(id)
	srv.namesMtx.Lock()
	srv.ClientNames[id] = req.Username
	srv.namesMtx.Unlock()
	return &chat.ChatToken{Token: id}, nil
}

func (srv *ChatServer) openStream(id string) chan *chat.StreamMessage {
	stream := make(chan *chat.StreamMessage, 1000)
	srv.streamsMtx.Lock()
	srv.ClientStreams[id] = stream
	srv.streamsMtx.Unlock()
	return stream
}

func (srv *ChatServer) closeStream(id string) {
	srv.streamsMtx.Lock()
	if stream, ok := srv.ClientStreams[id]; ok {
		delete(srv.ClientStreams, id)
		close(stream)
	}
	srv.streamsMtx.Unlock()
}

func (srv *ChatServer) broadcastExit(id string) {
	username, _ := srv.getUsername(id)
	srv.Broadcast <- &chat.StreamMessage{
		Timestamp: timestamppb.Now(),
		Message: &chat.StreamMessage_ClientLogout{
			ClientLogout: &chat.Logout{Username: username},
		},
	}
}

func (srv *ChatServer) broadcastEnter(id string) {
	username, _ := srv.getUsername(id)
	srv.Broadcast <- &chat.StreamMessage{
		Timestamp: timestamppb.Now(),
		Message: &chat.StreamMessage_ClientLogin{
			ClientLogin: &chat.Login{Username: username},
		},
	}
}

func (srv *ChatServer) Stream(ssrv chat.ChatRoom_StreamServer) error {
	tkn, ok := srv.extractToken(ssrv.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing token header")
	}
	if _, ok := srv.getUsername(tkn); !ok {
		return status.Error(codes.Unauthenticated, "missing name")
	}
	srv.openStream(tkn)
	srv.broadcastEnter(tkn)
	defer srv.closeStream(tkn)
	defer srv.broadcastExit(tkn)

	go srv.sendBroadcasts(ssrv, tkn)

	for {
		msg, err := ssrv.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		srv.Broadcast <- srv.handleStreamMsg(msg)
	}
}

func (srv *ChatServer) handleStreamMsg(msg *chat.StreamMessage) *chat.StreamMessage {
	return msg
}

func (srv *ChatServer) extractToken(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[tokenHeader]) == 0 {
		return "", false
	}
	return md[tokenHeader][0], true
}
