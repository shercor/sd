package main

import (
	pb "github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	
	
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}


func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}