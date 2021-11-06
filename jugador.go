package main
import (
    "fmt"
    "os"
	"bufio"
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

	var etapa = 1 // TO-DO: recibir esto desde Lider

	// conectar con lider

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewLiderServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	/*******************************************************************/

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
