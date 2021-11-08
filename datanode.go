package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	pb "github.com/shercor/sd/proto"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func leerRegistro(ID int32) string {

	str_msg := ""
	lista_txt := filtrarTxt(ID)

	if (len(lista_txt) == 0){
		return ""
	}

	for i := 0; i < len(lista_txt); i++ {
		// Lee cada archivo que involucre al jugador, uno por uno
		content, err := ioutil.ReadFile(lista_txt[i] ) 
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(string(content))
		str_msg += lista_txt[i] + ": " + string(content) + "\n" // Se envia un puro str, para enviar un puro msj

	}
	return str_msg
}

func filtrarTxt(ID int32) (lista_txt []string) { // Filtra y retorna los nombres de los txt que involucren al jugador ID

	path := "." // Imagino que este path sirve, sino ponerle el ospath de donde esta el archivo
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		str_jugador := "Jugador_" + strconv.FormatInt(int64(ID), 10)
		if strings.Contains(file.Name(), str_jugador) { // Si el archivo contiene "Jugador_ID", meterlo a la lista de los que hay que leer
			lista_txt = append(lista_txt, file.Name())
		}
	}

	return lista_txt
}

func escribirRegistro(ID_actual int32, etapa string, jugada string, ronda string) {

	nametxt := ""

	if etapa == "1" { // Si es la etapa 1, se creara un archivo para cada ronda, jugador_5__etapa_1__ronda_2.txt por ejemplo
		nametxt = "Jugador_" + strconv.FormatInt(int64(ID_actual), 10) + "__Etapa_" + etapa + "__Ronda_" + ronda + ".txt"
	} else {
		nametxt = "Jugador_" + strconv.FormatInt(int64(ID_actual), 10) + "__Etapa_" + etapa + ".txt"
	}

	f, err := os.Create(nametxt)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	// Escribir en un txt
	write_str :=  jugada
	f.WriteString(write_str + "\n")

}

/********************************** gRPC **********************************************/

type Server struct {
	pb.UnimplementedDataNodeServiceServer
}

func startServer(){
	/*  Iniciar servidor DataNode */
	fmt.Println("Iniciando servidor DataNode...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9300))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterDataNodeServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// funcion response para registrar jugada
func (s *Server) RegistrarJugada(ctx context.Context, in *pb.InfoJugada) (*pb.Message, error) {	
	escribirRegistro(in.ID, in.Etapa, in.Jugada, in.Ronda)
	return &pb.Message{Body: "OK"}, nil
}

// funcion response para consultar jugada
func  (s *Server) ConsultarJugada(ctx context.Context, in *pb.RespuestaSolicitud) (*pb.Message, error) {
	response := leerRegistro(in.ID)
	return &pb.Message{Body: response}, nil
}

/*
func main() {

	// Recibe las jugadas de cada jugador, mandadas desde el NameNode
	// Recibe el ID_Jugador, la etapa y la jugada (o jugadas en caso de la etapa 1)

	// Esperar mensajes
	// Cada vez que llega un msj, hace esto:
	peticion := "escribir"

	// HARDCODEO, obtener esto con mensajes y ponerle loop
	// Deben ser int32 todos los numeros
	ID := 1
	etapa := 1
	jugada := 5
	ronda := 1 // RONDA SOLO SE MANDA EN ETAPA 1 o de lo contrario mandar siempre con valor 0 (igual no se usará xd)

	if peticion == "escribir" {
		escribirRegistro(ID, etapa, jugada, ronda)
	} else if peticion == "leer" {
		leerRegistro(ID)
	}

}
*/

func main() {

	go startServer() // gRPC
	
	for {	
	}
}
