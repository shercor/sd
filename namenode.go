package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	pb "github.com/shercor/sd/proto"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func getRandomNum(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min+1) + min
	return result
}

/********************************** gRPC **********************************************/

type Server struct {
	pb.UnimplementedNameNodeServiceServer
}

// funcion response para consultar jugada
func  (s *Server) ConsultarJugada(ctx context.Context, in *pb.RespuestaSolicitud) (*pb.Message, error) {
	full_response := ""
	for i := 0; i < 3; i++ { // consultar cada NameNode
		// conectar con NameNode	
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(ip_datanodes[i] + ":9300", grpc.WithInsecure() )
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := pb.NewDataNodeServiceClient(conn)
		
		response, err := c.ConsultarJugada(context.Background(), &pb.RespuestaSolicitud{ID: in.ID})
		if err != nil {
			log.Fatalf("Error when calling ConsultarJugada: %s", err)
		}

		full_response = full_response + response.Body
	
	}
	
	return &pb.Message{Body: full_response}, nil
}

// funcion response para registrar jugada
func (s *Server) RegistrarJugada(ctx context.Context, in *pb.InfoJugada) (*pb.Message, error) {
	
	// seleccionar IP random
	opt := getRandomNum(0, 2) // Elegir un indice del slice al azar

	// Escribir en un txt
	f, err := os.OpenFile("namenode.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	write_str := ""

	if in.Etapa == "1" { // Si es la etapa 1 se crea con la ronda ademas
        write_str = "Jugador" + strconv.Itoa(int(in.ID))  + " Etapa" + in.Etapa + " Ronda_" + in.Ronda + " " + ip_datanodes[opt] + "\n"
    } else {
		write_str = "Jugador_" + strconv.Itoa(int(in.ID)) + " Etapa_" + in.Etapa + " " + ip_datanodes[opt] + "\n"
    }
	
	// f.WriteString(write_str + "\n")
	if _, err = f.WriteString(write_str); err != nil {
		panic(err)
	}

	// conectar con NameNode	
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ip_datanodes[opt] + ":9300", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewDataNodeServiceClient(conn)
	
	response, err := c.RegistrarJugada(context.Background(), &pb.InfoJugada{ID: in.ID, Etapa: in.Etapa, Jugada: in.Jugada, Ronda: in.Ronda})
	if err != nil {
		log.Fatalf("Error when calling RegistrarJugada: %s", err)
	}
	fmt.Println(response.Body)


	return &pb.Message{Body: "OK"}, nil
}

func startServer(){
	/*  Iniciar servidor NameNode */
	fmt.Println("Iniciando servidor NameNode...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9400))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterNameNodeServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

var ip_datanodes []string
func main() {
	
	
	/*
	ip_datanodes = append(ip_datanodes, "") // local 
	ip_datanodes = append(ip_datanodes, "")
	ip_datanodes = append(ip_datanodes, "")
	*/
	ip_datanodes = append(ip_datanodes, "10.6.43.101") // VM 
	ip_datanodes = append(ip_datanodes, "10.6.43.102")
	ip_datanodes = append(ip_datanodes, "10.6.43.103")



	go startServer() // gRPC
	
	f, _ := os.Create("namenode.txt")
	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	for {		
	}

}
