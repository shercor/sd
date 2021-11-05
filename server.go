package main

import (
	pb "github.com/shercor/sd/proto"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type Server struct {
	pb.UnimplementedLiderServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}


func (s *Server) Unirse(ctx context.Context, req *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	fmt.Println("Uniendo jugador IP" + req.IP)
	return &pb.RespuestaSolicitud{
		ID: 1,
	}, nil
}


func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterLiderServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}











