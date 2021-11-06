package main
import (
    "fmt"
    "os"
	"bufio"
	"net"
	"strings"
	//"net"
	//"strconv"
	//"math/rand"
	pb "github.com/shercor/sd/proto"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//func etapa_uno(x int, y int) int {
func etapa_uno(bot bool) int {

	// elegir numero entre el 1 y 10 


	return 0
}

func etapa_dos(bot bool) int {
	return 0
}


func etapa_tres(bot bool) int {
	return 0
}

var my_id int32 // ID jugador

func main() {
	argsWithoutProg := os.Args[1:]
	// checker si es bot o no
	var bot = false
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "bot"{
		bot = true
	}

	if bot {
		fmt.Println("Bienvenid@ bot")
	} else{
		fmt.Println("Bienvenid@ jugador")
	}

	var continue_flag = true
	var alive = false

	
	fmt.Println("Elija opción:")
	fmt.Println("1. Unirse al juego del calamar\n2. Terminar todo")

	for continue_flag {	
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
	
		if err != nil {
			fmt.Println(err)
		}

		switch char {
			case '1':
				alive = true
				continue_flag = false
			case '2':
				continue_flag = false
			default: 
				fmt.Println("Respuesta no valida")
		}
	}

	// obtener puerto disponible para el jugador
	// rango para el puerto aleatorio de la partida
	/*
	min := 49154
	max := 65535

	PUERTO_JUGADOR := ""
	PUERTO_JUGADOR = strconv.Itoa(rand.Intn(max-min) + min)

	for{
		s, err := net.ResolveUDPAddr("udp4", PUERTO_JUGADOR) 
		if err != nil {
			fmt.Println(err)
			PUERTO_JUGADOR := strconv.Itoa(rand.Intn(max-min) + min)
			continue
		}
		s.Close()
		break
	}				
	fmt.Println("PUERTO_JUGADOR", PUERTO_JUGADOR)
	*/

	// solicitar entrar al juego del calamar

	// conectar con lider

	// obtener mi IP, se hace una conexion UDP a DNS para obtener mi IP
	c_temp, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
		log.Fatalf("Error %s", err)
	}

    localAddr := c_temp.LocalAddr().(*net.UDPAddr).String()
	c_temp.Close()

	fmt.Println(strings.Split(localAddr, ":")[0])

	my_IP := strings.Split(localAddr, ":")[0]
	my_PORT := strings.Split(localAddr, ":")[1]

	
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewLiderServiceClient(conn)

	/*
	// request: mensaje de saludo
	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	*/
	
	// solicitar unirse al juego del calamara 
	response, err := c.Unirse(context.Background(), &pb.Solicitud{IP: my_IP, PORT: my_PORT})
	if err != nil {
		log.Fatalf("Error when calling Unirse: %s", err)
	}
	my_id = response.ID
	log.Printf("ID asignado por Lider: %s", response.ID)
	
	/*******************************************************************/

	var etapa = 1 
	for alive {
		fmt.Println("-----------Vivo----------")
		switch etapa {
		case 1:
			fmt.Println("Etapa 1")
		case 2:
			fmt.Println("Etapa 2")
		case 3:
			fmt.Println("Etapa 2")
		default: 
			fmt.Println("Etapa no valida")
	}

		// TO-DO: logica de juegos 

		break // TO-DO: logica de salida
	}
	fmt.Println("Finalizando proceso jugador")
	
	
	/*
	fmt.Println("Hola, mundo.")
	fmt.Println("Bot? ", bot)

	var a string = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
	*/

	

}
