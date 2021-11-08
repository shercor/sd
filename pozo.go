package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	pb "github.com/shercor/sd/proto"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) { // Para errores
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

/********************************** gRPC **********************************************/

type Server struct {
	pb.UnimplementedPozoServiceServer
}

func startServer(){
	/*  Iniciar servidor Pozo */
	fmt.Println("Iniciando servidor Pozo...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9500))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterPozoServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func  (s *Server)  ConsultarMontoAcumulado(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from Lider: %s", in.Body)
	monto_acumulado := strconv.Itoa(pozo)
	return &pb.Message{Body: monto_acumulado}, nil
}

/*************************************************************************************/

var pozo int // pozo acumulado

func main() {

	go startServer() // gRPC

	// Recibe un msj con el ID de un jugador muerto y la etapa en la que fue eliminado
	// Este se debe enviar en la funcion esEliminado de lider
	// Debe quedarse escuchando e ir escribiendo en el txt
	
	//cant_jugadores := 16 // Esto se asume pero para testear CAMBIARLOOOOOOOOOOOO
	pozo = 0
	const wones = 100000000

	f, _ := os.Create("pozo.txt")

	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	// RabbitMQ
	conn, err := amqp.Dial("amqp://pozo:pozo@10.6.43.102:5672/") 
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declara una queue
	q, err := ch.QueueDeclare(
		"wones", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Println(q)

	// Asíncrono - RabbitMQ
	msgs, err := ch.Consume( // Consumo mensajes
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool) // Crear un canal para recibir mensajes en loop infinito

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body) // recibe mensaje

			// d.Body (type []byte) 
			s := string(d.Body)
			
			ID_actual := strings.Split(s, "_")[0]
			etapa := strings.Split(s, "_")[1]
			
			// Por cada mensaje, aumenta los wones y escribe en el txt el jugador que murio
			pozo += wones
			write_str := "Jugador_" + ID_actual + " Etapa_" + etapa + " " + strconv.Itoa(pozo)
			f.WriteString(write_str + "\n")
		}
	}()
	<-forever

	for {	
	}

}
