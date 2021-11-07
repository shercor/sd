package main

import (
	"fmt"
	"sync"
	"math/rand"
	"os"
	"strconv"
	"time"
	"bufio"
	pb "github.com/shercor/sd/proto"
	"log"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

/************** Lock ***************/

// incrementar valor de contador
func (c *Container) inc(name string) {
    c.Lock()
    defer c.Unlock()
    c.counters[name]++
}

// reset valor de contador
func (c *Container) reset(name string) {
    c.Lock()
    defer c.Unlock()
    c.counters[name] = 0
}


/********************************** gRPC **********************************************/

type Server struct {
	pb.UnimplementedLiderServiceServer
}

// funcion response hello para debug
func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

// funcion response para unir jugador al juego
func (s *Server) Unirse(ctx context.Context, in *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	log.Printf("IP de jugador: %s, port: %s", in.IP, in.PORT)
	
	// append jugador a lista de jugadores
	var id int32
	id = int32(jugadores_conectados) + 1
	jugadores_conectados = jugadores_conectados + 1

	lista_jugadores = append(lista_jugadores, newJugador(id, in.IP, in.PORT)) // Jugador 1, 2, 3, ..., 16

	log.Printf("Jugadores conectados: %d", jugadores_conectados)
	return &pb.RespuestaSolicitud{ID: id}, nil
}

// funcion response procesar jugada de jugador
func (s *Server) ProcesarJugada(ctx context.Context, in *pb.Jugada) (*pb.Message, error) {        

	log.Printf("ID de jugador: %d, jugada:  %s", in.ID, in.Numero)
	
	// Guardar el ID de esa respuesta

	ID_rpta := in.ID
	rpta, err := strconv.Atoi(in.Numero)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	if rpta >= opt_lider {
		// Jugador eliminado
		esEliminado(*lista_jugadores[ID_rpta-1])
		muertos_por_ronda = append(muertos_por_ronda, ID_rpta) // Añadido a la lista de muertos
	} else {
		// Mandar mensaje que el jugador sigue vivo
		// Mostrar el numero que lleva acumulado (TO-DO)
		lista_jugadores[ID_rpta-1].suma_rptas_etapa1 += rpta
		if lista_jugadores[ID_rpta-1].suma_rptas_etapa1 >= 21 { // Si la suma de sus respuestas es >= 21, informar que pasa a la siguiente etapa
			// Informar que pasa a la siguiente etapa
			pasan_etapa = append(pasan_etapa, ID_rpta) // Como ya paso, no se calcula en los que faltan
		}
	}

	container.inc("cont_req")

	// Registrar su jugada en NameNode
	
	return &pb.Message{Body: "OK"}, nil 
}

// funcion response para resultado de ronda
func (s *Server) GetResultadosRonda (ctx context.Context,  in *pb.RespuestaSolicitud) (*pb.ResultadoJugada, error) {
	id_jugador := in.ID 
	
	if container.counters["cont_req"] < en_juego { // si aun no juegan todos los jugadores
		// Mandar msj a jugador
		return &pb.ResultadoJugada{Vivo: true, NEXTETAPA: false, WAIT: true}, nil 
	}
	
	vivo_bool := true
	next_etapa := false

	// Leer slices e informar quienes pasan o no
	
	// sirve en etapa 1
	for i := 0; i < len(muertos_por_ronda); i++ {
		playerID := muertos_por_ronda[i]
		if (playerID == id_jugador){
			fmt.Println("El jugador", id_jugador, "es eliminado")
			vivo_bool = false
			// Mandar msj al pozo para aumentarlo
		}

		
	}
	for i := 0; i < len(pasan_etapa); i++ {
		playerID := pasan_etapa[i]
		if (playerID == id_jugador){
			fmt.Println("El jugador", playerID, " pasa a la siguiente etapa")	
			next_etapa = true
			break
		}
	}

	container.inc("cont_res")

	// Mandar msj a jugador
	return &pb.ResultadoJugada{Vivo: vivo_bool, NEXTETAPA: next_etapa, WAIT: false}, nil 
}

// funcion response para empezar proxima etapa o notificar ganador
func  (s *Server)  EmpezarEtapa(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	if flag_next_etapa == false{ // aun no empieza la etapa
		return &pb.Message{Body: "NOK"}, nil
	
	} else if notificar_ganador == true { // notificar ganador(es)
		// TO-DO: hacer logica para multiples ganadores
		container.inc("cont_res")		
		return &pb.Message{Body: "GANADOR"}, nil
	}

	container.inc("cont_res")		
	return &pb.Message{Body: "OK"}, nil
}

/***************************************************************************************/

type jugador struct {
	ID                int32
	IP                string
	puerto            string
	estado            string
	suma_rptas_etapa1 int
	equipo_etapa2     string
	nro_pareja_etapa3 int
	rpta_etapa3       int
}

func newJugador(ID int32, IP string, puerto string) *jugador { // Crea un struct jugador con ID, IP, puerto, etc
	player := jugador{
		ID:                ID,
		IP:                IP,
		puerto:            puerto,
		estado:            "vivo",
		suma_rptas_etapa1: 0,
		nro_pareja_etapa3: 0,
		rpta_etapa3:       0}
	return &player
}

func obtenerNumJugadores() { // Parte del codigo Alvaro para obtener el num de jugadores
	argsWithoutProg := os.Args[1:]
	//var n_jugadores = 16
	if len(argsWithoutProg) == 1 {
		n_jugadores, err := strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n_jugadores)
	}
}

func getRandomNum(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min+1) + min
	return result
}

func recibirPeticion() { // Recibe peticion para jugar, viene desde un jugador

}

func recibirJugada() { // Recibe la jugada de un jugador en una cierta etapa

}

func informarJugadas() { // Informa a NameNode cada vez que un jugador hace una jugada

}

func informarPozo() { // Agrega 100 millones a Pozo cuando un jugador muere

}

func jugadoresVivos(lista_jugadores []*jugador, cant_jugadores int, print bool) (lista_vivos []int32) { // Muestra por consola a los jugadores vivos al término de cada etapa
	for i := 0; i < cant_jugadores; i++ {
		if lista_jugadores[i].estado == "vivo" {
			if print {
				fmt.Println("El jugador:", lista_jugadores[i].ID, "está vivo")
			}
			lista_vivos = append(lista_vivos, lista_jugadores[i].ID)
		}
	}
	return lista_vivos
}

func mostrarGanadores(lista_jugadores []*jugador, cant_jugadores int) { // Muestra por consola a los ganadores y wones ganados
	// TO DO: Añadir la logica de mostrar los wones
	for i := 0; i < cant_jugadores; i++ {
		if lista_jugadores[i].estado == "vivo" {
			fmt.Println("El jugador:", lista_jugadores[i].ID, "ha ganado el juego del calamar")
		}
	}
	// TO-DO: notificar con el mensaje a ganador con el parametro GANADOR
}

func revisarJugadas() { // Para preguntar sobre las jugadas historicamente de un determinado jugador

}

func esEliminado(jugador_eliminado jugador) { // Cuando la logica detecta que el jugador muere, avisa al jugador
	jugador_eliminado.estado = "muerto"
	mostrarMuerte(jugador_eliminado.ID) // Printea por consola quien murio
}

func mostrarMuerte(ID_player int32) { // Muestra por consola cuando muere un jugador
	// enviar notificacion de que perdio
	fmt.Println("El jugador:", ID_player, "ha sido eliminado")
}


func evaluarRestantes(lista_vivos []int32, lista_jugadores []*jugador) int { // Evalua si quedan jugadores suficientes para jugar
	num_vivos := len(lista_vivos)
	if num_vivos == 0 {
		// terminar el juego
		fmt.Println("Todos los jugadores fueron eliminados")
		return 0

	} else if num_vivos == 1 { // TO-DO: hacer logica para multiples ganadores
		// Notificar ganador
		fmt.Println("Solo queda un jugador en pie")
		fmt.Println("El jugador ", lista_vivos[0], " gana el juego del Calamar")
		return 2
	} else {
		// Habilitar que jugadores que pasan a la fase 2
		fmt.Println("Los jugadores de ID:")
		for i := 0; i < num_vivos; i++ {
			fmt.Println(lista_vivos[i])
		}
		fmt.Println("pasan a la siguiente etapa")
		return 1
	}
}

func RemoveIndex(s []int32, index int32) []int32 { // Elimina la posicion index del slice, codigo de stackoverflow
	if (len(s) == 0){ // slice vacio
		return s
	}
	ret := make([]int32, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func eliminarJugadorImpar(lista_jugadores []*jugador, lista_vivos []int32) (new_lista_vivos []int32) {
	index_eliminado := getRandomNum(0, len(lista_vivos))
	ID_eliminado := lista_vivos[index_eliminado]
	esEliminado(*lista_jugadores[ID_eliminado-1])
	new_lista_vivos = RemoveIndex(lista_vivos, int32(index_eliminado) ) // Elimina de la lista de vivos al jugador del index que se debe eliminar
	return new_lista_vivos // slice con el jugador eliminado
}

func asignarGrupos(lista_jugadores []*jugador, lista_vivos []int32) {
	for i := 0; i < len(lista_vivos); i++ {
		currentID := lista_vivos[i]
		fmt.Println("Current ID:", currentID)
		if i < len(lista_vivos)/2 {
			// Mandar mensaje a los jugadores con su equipo asignado
			lista_jugadores[currentID-1].equipo_etapa2 = "A"
		} else {
			// Mandar mensaje a los jugadores con su equipo asignado
			lista_jugadores[currentID-1].equipo_etapa2 = "B"
		}
	}
}

func eliminarGrupo(lista_jugadores []*jugador, lista_vivos []int32, letra string) (new_lista_vivos []int32) {
	for i := 0; i < len(lista_vivos); i++ {
		currentID := lista_vivos[i]
		if lista_jugadores[currentID-1].equipo_etapa2 == letra {
			// Si el jugador es del grupo 'letra', se elimina
			esEliminado(*lista_jugadores[currentID-1])
			lista_vivos = RemoveIndex(lista_vivos, currentID)
		}
	}
	new_lista_vivos = lista_vivos
	return new_lista_vivos // Retorna la lista de los vivos cuando ya se eliminó el grupo 'letra'
}

func Abs(x int) int { // GO solo tiene abs para floats, me complicaba el codigo asi que hice mi propia xd
	if x < 0 {
		return -x
	}
	return x
}

func startServer(){
	/*  Iniciar servidor Lider */
	fmt.Println("Iniciando servidor Lider...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("... listen exitoso")
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterLiderServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

/* variables globales */
var cant_jugadores int
var lista_jugadores []*jugador // Slice de structs con los jugadores
var jugadores_vivos int
var jugadores_conectados int

// Container para usar Lock
type Container struct {
    sync.Mutex
    counters map[string]int
}
var container Container 

var muertos_por_ronda []int32 // Una lista de muertos ([1,4,5,8,16] por ejemplo)
var pasan_etapa []int32       // Una lista de quienes pasan la etapa, por ronda
var opt_lider int // opcion del lider de etapa 1
var en_juego int // jugadores en juego
var flag_next_etapa bool  // flag para empezar siguiente etapa
var notificar_ganador bool // flag para notificar ganador

func main() {
	// Definiciones iniciales

	cant_jugadores = 16
	jugadores_conectados = 0

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		param, err := strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		cant_jugadores = param
		fmt.Println(cant_jugadores)
	}

	const wones = 100000000 // 100 millones de wones vale cada jugador

	jugadores_vivos = cant_jugadores

	go startServer()
	
	// Setup inicial recibiendo los jugadores
	for jugadores_conectados < cant_jugadores { // esperar hasta que se conecten los jugadores
	}

	/*** iniciar container ***/
	
	// cont_req: contador para los req de jugadores
	// cont_res: contador para respuestas de lider
	container = Container{
		counters: map[string]int{"cont_req": 0, "cont_res": 0},
	}

	/*** Inicio de etapa 1 ***/
	// Primero aqui avisa a los jugadores que iniciara la etapa 1 y que manden sus respuestas
	flag_next_etapa = false
	var etapa = 1
	const max_rondas = 4
	var contador_rondas = 0
	en_juego = jugadores_vivos

	/******************** LOOP ETAPA 1 ********************************/
	fmt.Println("Inicio de Etapa 1")
	for contador_rondas < max_rondas {
		
		// Lider escoge un numero al azar entre el 6 y 10
		opt_lider = getRandomNum(6, 10)
		
		// vaciar slices antes de empezar ronda
		muertos_por_ronda = nil  // Una lista de muertos ([1,4,5,8,16] por ejemplo)
		pasan_etapa = nil

		fmt.Println("La opcion del Lider es:", opt_lider)
		
		// esperar a que todos los vivos jueguen		
		container.reset("cont_req")
		for container.counters["cont_req"] < en_juego{
		}

		// esperar a que el lider informe el resultado a jugadores
		container.reset("cont_res")		
		for container.counters["cont_res"] < en_juego{
		}

		// update de variables
		jugadores_vivos -= len(muertos_por_ronda)
		en_juego -= len(muertos_por_ronda)
		en_juego -= len(pasan_etapa)
		contador_rondas += 1
	}
	/********************************************************************/

	// Chequear si hay jugadores vivos para seguir jugando, y quienes
	var lista_vivos = jugadoresVivos(lista_jugadores, cant_jugadores, false)
	resultado_etapa := evaluarRestantes(lista_vivos, lista_jugadores) // Evalua si quedan jugadores suficientes para jugar
	if resultado_etapa == 0 { // no hay jugadores en pie
		fmt.Println("Terminando juego del calamar sin ganadores.")
		fmt.Println("Cerrando proceso lider.")
		return
	}else if resultado_etapa == 2{ // un solo jugador en pie, el ganador 
		fmt.Println("Notificando ganador...")
		notificar_ganador = true
		// esperar a que se notifique al ganador
		// TO-DO: hacer logica para multiples ganadores (aunque no es necesario en esta etapa)
		container.reset("cont_res")		
		for container.counters["cont_res"] < 1 {
		}
		fmt.Println("Cerrando proceso lider.")
		return 
	}

	fmt.Println("¿Continuar a la siguiente etapa? (y/n)")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	if (char != 'y'){
		fmt.Println("Terminando juego del calamar inconcluso.")
		return
	}
	flag_next_etapa = true

	// esperar notificar a todos que sigue la etapa
	container.reset("cont_res")		
	for container.counters["cont_res"] < len(lista_vivos){
	}

	/*** Procesamiento etapa 2 ***/
	// Ver si los restantes son impar
	if len(lista_vivos)%2 == 1 {
		fmt.Println("Sobrevivientes impares, se eliminará a uno")
		lista_vivos = eliminarJugadorImpar(lista_jugadores, lista_vivos) // Se elimina uno al azar y retorna la nueva lista de vivos
		jugadores_vivos -= 1
	}

	/*** Inicio de etapa 2 ***/
	// Primero aqui avisa a los jugadores que iniciara la etapa 2 y que manden sus respuestas
	fmt.Println("Inicio de Etapa 2")
	flag_next_etapa = false
	etapa = etapa + 1

	asignarGrupos(lista_jugadores, lista_vivos) // La mitad de los vivos se van al grupo A, la otra mitad al B (en enunciado no dice como escogerlos)
	var grupoA_suma = 0
	var grupoB_suma = 0
	var len_grupo = jugadores_vivos / 2
	var paridad_A = false
	var paridad_B = false
	// Lider escoge un numero al azar entre el 1 y 4
	opt_lider := getRandomNum(1, 4)
	fmt.Println("La opcion del Lider es:", opt_lider)

	for i := 0; i < jugadores_vivos; i++ {

		// Recoger primera respuesta que llegue
		// Guardar el ID de esa respuesta

		// Hardcodeo
		ID_rpta := 1
		rpta := 4

		// aqui se recoje las sumas
		if lista_jugadores[ID_rpta-1].equipo_etapa2 == "A" {
			grupoA_suma += rpta
		} else if lista_jugadores[ID_rpta-1].equipo_etapa2 == "B" {
			grupoB_suma += rpta
		} else {
			// Algo paso que entro un muerto
			fmt.Println("Recibio a alguien que no tenia un grupo asignado")
		}
	}

	if opt_lider%2 == grupoA_suma%2 { // Si tiene la misma paridad, se activa el flag
		// Activa flag grupo A
		paridad_A = true
	}
	if opt_lider%2 == grupoB_suma%2 {
		// Activa flag grupo B
		paridad_B = true
	}

	// Chequear los resultados de paridad y eliminar en base a ellos
	if !paridad_A && !paridad_B { // Si ninguno de los dos tiene la paridad del lider, se elimina uno aleatoriamente
		eleccion := getRandomNum(0, 1) // Si es 0 se elimina A, si es 1 se elimina B
		if eleccion == 0 {
			lista_vivos = eliminarGrupo(lista_jugadores, lista_vivos, "A")
		} else {
			lista_vivos = eliminarGrupo(lista_jugadores, lista_vivos, "B")
		}
		jugadores_vivos -= len_grupo
	} else if !paridad_A { // La paridad no es la misma, eliminar A
		lista_vivos = eliminarGrupo(lista_jugadores, lista_vivos, "A")
		jugadores_vivos -= len_grupo

	} else if !paridad_B { // La paridad no es la misma, eliminar B
		lista_vivos = eliminarGrupo(lista_jugadores, lista_vivos, "B")
		jugadores_vivos -= len_grupo
	}

	// TO-DO: poner flag para termino de calculos para notificar quienes siguen 


	// A estas alturas se elimino un grupo, y si ambos obtienen la misma paridad no se elimino nada ni se restaron los jugadores_vivos
	// Termino de fase 2
	// Avisar por mensaje que equipos ganaron
	// Avisar cantidad y que IDs siguen jugando, ademas de pozo de wones

	// Chequear si hay jugadores vivos para seguir jugando, y quienes
	lista_vivos = jugadoresVivos(lista_jugadores, cant_jugadores, false)
	evaluarRestantes(lista_vivos, lista_jugadores) // Evalua si quedan jugadores suficientes para jugar

	// Ver si los restantes son impar
	if len(lista_vivos)%2 == 1 {
		fmt.Println("Sobrevivientes impares, se eliminará a uno")
		lista_vivos = eliminarJugadorImpar(lista_jugadores, lista_vivos) // Se elimina uno al azar y retorna la nueva lista de vivos
		jugadores_vivos -= 1
	}

	// Inicio de etapa 3
	// Avisa a los jugadores que iniciara la etapa 3
	fmt.Println("Inicio de Etapa 3")
	etapa = etapa + 1

	// Elige parejas y los informa a los jugadores
	asigna_parejas := 1 // Las parejas van del 1 hasta el jugadores_vivos/2
	for i := 0; i < jugadores_vivos; i++ {
		ID_player1 := lista_vivos[i]
		i++
		ID_player2 := lista_vivos[i]
		// Quedan las parejas asignadas
		lista_jugadores[ID_player1-1].nro_pareja_etapa3 = asigna_parejas
		lista_jugadores[ID_player2-1].nro_pareja_etapa3 = asigna_parejas

		// Avisar por mensaje a los jugadores su numero de pareja
		asigna_parejas++
	}

	// Les pide que manden sus respuestas
	for i := 0; i < jugadores_vivos; i++ {

		// Recoger primera respuesta que llegue
		// Guardar el ID de esa respuesta

		// Hardcodeo
		ID_rpta := 1
		rpta := 4

		lista_jugadores[ID_rpta-1].rpta_etapa3 = rpta

	}

	// Lider escoge un numero random entre 1 y 10
	opt_lider = getRandomNum(1, 10)

	// Procesa sus respuestas, compara resultados entre si y con el lider
	for nro_pareja := 1; nro_pareja < jugadores_vivos/2; nro_pareja++ {
		aux_break := 0
		var vs_pareja []int32 // Para guardar los ID que componen cada pareja

		// Busca el ID de los integrantes de la pareja nro 'nro_pareja' y los guarda en vs_pareja
		for i := 0; i < jugadores_vivos; i++ {
			ID_player := lista_vivos[i]
			if lista_jugadores[ID_player-1].nro_pareja_etapa3 == nro_pareja { // Ver si es el num de pareja que se está chequeando
				vs_pareja = append(vs_pareja, ID_player)
				aux_break++
			}
			if aux_break == 2 { // Por temas de optimizacion para que no recorra toda la lista xdddd
				break
			}
		}

		// Procesa y envia resultados (moriste o no)
		rpta1 := Abs(opt_lider - lista_jugadores[vs_pareja[0]-1].rpta_etapa3)
		rpta2 := Abs(opt_lider - lista_jugadores[vs_pareja[1]-1].rpta_etapa3)
		fmt.Println("players:", vs_pareja[0], "VS", vs_pareja[1])
		fmt.Println(rpta1, "VS", rpta2)

		if rpta1 > rpta2 { // Pierde rpta1
			// Avisar al jugador con ID vs_pareja[0] que perdio
			esEliminado(*lista_jugadores[vs_pareja[0]-1])
		} else if rpta1 < rpta2 { // Pierde rpta2
			// Avisar al jugador con ID vs_pareja[1] que perdio
			esEliminado(*lista_jugadores[vs_pareja[1]-1])
		} // Si no entra a ninguno de los dos es que ambos ganaron, no se elimina a nadie
	}

	// Entrega ganadores y la cantidad de wones
	mostrarGanadores(lista_jugadores, cant_jugadores)

}
