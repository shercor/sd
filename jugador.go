package main
import (
    "fmt"
    "os"
	"bufio"
	"net"
	"strings"
	"strconv"
	"math/rand"
	"time"
	pb "github.com/shercor/sd/proto"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// TO-DO: func ver pozo

func etapa_uno(bot bool, my_ID int32, c pb.LiderServiceClient) int {

	// elegir numero entre el 1 y 10 
	eleccion := 0 

	if bot == false{
		fmt.Println("Elija numero del 1 al 10:")

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		eleccion = int(char - '0')

	}else{
		min := 1
		max := 10
		eleccion = rand.Intn(max-min) + min
	}

	fmt.Println(strconv.Itoa(eleccion))

	// procesar jugada
	response, err := c.ProcesarJugada(context.Background(), &pb.Jugada{ID: my_ID, Numero: strconv.Itoa(eleccion)})
	if err != nil {
		log.Fatalf("Error when calling ProcesarJugada: %s", err)
	}
	log.Printf(response.Body)

	contador_etapa_1 = contador_etapa_1 + eleccion
	fmt.Println("Sumas el valor de ", contador_etapa_1)

	// esperar Resultados ronda (que todos los jugadores vivos respondan)
	fmt.Println("Esperando jugada de todos los jugadores vivos...")
	wait_jugadores := true
	flag_next_etapa := false
	flag_vivo := true

	for wait_jugadores == true{
		response_ronda, err := c.GetResultadosRonda(context.Background(), &pb.RespuestaSolicitud{ID: my_ID})
		if err != nil {
			log.Fatalf("Error when calling GetResultadosRonda: %s", err)
		}
		wait_jugadores = response_ronda.WAIT
		flag_next_etapa = response_ronda.NEXTETAPA
		flag_vivo = response_ronda.Vivo
		time.Sleep(1*time.Second)
	}

	//fmt.Println(flag_vivo, flag_next_etapa)

	if flag_next_etapa == true{ // pasa a la siguiente etapa
		return 2
	}

	if flag_vivo == false{ // jugador muere
		return 0
	}
	return 1 // jugador pasa de ronda
}

func etapa_dos(bot bool) int {
	return 0
}


func etapa_tres(bot bool) int {
	return 0
}

// rutina para esperar a que empiece la proxima etapa
func esperar_etapa(c pb.LiderServiceClient) int {
	fmt.Println("Esperando que empiece la proxima etapa...")
	wait_jugadores := true
	flag_ganador := false

	for wait_jugadores == true{ 
		response_ronda, err := c.EmpezarEtapa(context.Background(), &pb.Message{Body: "wait"})
		if err != nil {
			log.Fatalf("Error when calling EmpezarEtapa: %s", err)
		}
		if (response_ronda.Body == "OK"){
			wait_jugadores = false
		} else if (response_ronda.Body == "GANADOR"){
			wait_jugadores = false
			flag_ganador = true
		}
		time.Sleep(1*time.Second)
	}

	if flag_ganador == true { // si gana el juego del calamar response_ronda.Body == "GANADOR"
		return 1
	}

	// si no
	return 0
}

var my_id int32 // ID jugador
var contador_etapa_1 int // contador para interfaz en etapa 1

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
	
	// solicitar unirse al juego del calamar 
	response, err := c.Unirse(context.Background(), &pb.Solicitud{IP: my_IP, PORT: my_PORT})
	if err != nil {
		log.Fatalf("Error when calling Unirse: %s", err)
	}
	my_id = response.ID
	log.Printf("ID asignado por Lider: %d", response.ID)
	
	/*******************************************************************/

	var etapa = 1
	contador_etapa_1 = 0
	for alive {
		var estado = 0
		var ronda = 0
		var ganador = 0
		switch etapa {
		case 1:
			fmt.Println("Etapa 1")
			for estado != 2 { // estado == 2 es que pasa a la siguiente etapa (o es ganador)
				fmt.Println("Ronda: ", ronda+1)
				estado = etapa_uno(bot, my_id, c)
				if estado == 0 {
					alive = false
					break
				}
				ronda = ronda + 1
			}		
		case 2:
			fmt.Println("Etapa 2")
		case 3:
			fmt.Println("Etapa 2")
		default: 
			fmt.Println("Etapa no valida")
		}

		if (alive  == false){
			fmt.Println("Has muerto")
			break
		}
		if (estado == 2) {
			fmt.Println("Pasas a la siguiente ronda")
			ganador = esperar_etapa(c)
			etapa = etapa + 1
			fmt.Println("-------------------------")
		}		

		if (ganador == 1){
			fmt.Println("Ganaste el juego del Calamar")
			fmt.Println("Ganaste X wones")

			// consultar wones con jugador->lider->pozo y de vuelta
		}
	}

	

	fmt.Println("Finalizando proceso jugador")
}
