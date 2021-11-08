package main

import (
	"fmt"
	"sync"
	"math/rand"
	"os"
	"strconv"
	"time"
	"bufio"
	"strings"
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

// rutina para registrar jugada en NameNode
func registrarJugada(ID int32, etapa string, rpta int, is_ronda bool){
	ronda := "0"
	if is_ronda == true {
		 ronda =strconv.Itoa(contador_rondas + 1)
	}
	jugada := strconv.Itoa(rpta)

	fmt.Print("Registrando jugada de jugador ", strconv.Itoa(int(ID) ), "...")
	// conectar con NameNode	
	var conn *grpc.ClientConn
	//conn, err := grpc.Dial(":9400", grpc.WithInsecure())
	conn, err := grpc.Dial("10.6.43.104:9400", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_lider := pb.NewNameNodeServiceClient(conn)
	
	response, err := c_lider.RegistrarJugada(context.Background(), &pb.InfoJugada{ID: ID, Etapa: etapa, Jugada: jugada, Ronda: ronda})
	if err != nil {
		log.Fatalf("Error when calling RegistrarJugada: %s", err)
	}
	fmt.Println(response.Body)

}

type Server struct {
	pb.UnimplementedLiderServiceServer
}

// funcion response hello para debug
func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) ConsultarMontoAcumulado(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	// conectar con Pozo	
	var conn *grpc.ClientConn
	//conn, err := grpc.Dial(":9500", grpc.WithInsecure())
	conn, err := grpc.Dial("10.6.43.102:9500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c_lider := pb.NewPozoServiceClient(conn)
	
	response, err := c_lider.ConsultarMontoAcumulado(context.Background(), &pb.Message{Body: "CONSULTA"})
	if err != nil {
		log.Fatalf("Error when calling ConsultarMontoAcumulado: %s", err)
	}

	monto_acumulado := response.Body
	
	return &pb.Message{Body: monto_acumulado}, nil
}

// funcion response para unir jugador al juego
func (s *Server) Unirse(ctx context.Context, in *pb.Solicitud) (*pb.RespuestaSolicitud, error) {
	log.Printf("IP de jugador: %s, port: %s", in.IP, in.PORT)
	
	// append jugador a lista de jugadores
	var id int32
	id = int32(jugadores_conectados) + 1
	//fmt.Println("ID ASIGNADA: ", id)
	jugadores_conectados = jugadores_conectados + 1

	lista_jugadores = append(lista_jugadores, newJugador(id, in.IP, in.PORT)) // Jugador 1, 2, 3, ..., 16

	log.Printf("Jugadores conectados: %d", jugadores_conectados)
	return &pb.RespuestaSolicitud{ID: id}, nil
}

// funcion response procesar jugada de jugador en etapa 1
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
		//esEliminado(*lista_jugadores[ID_rpta-1])
		esEliminado(ID_rpta-1)
		muertos_por_ronda = append(muertos_por_ronda, ID_rpta) // Añadido a la lista de muertos
	} else {
		// Mandar mensaje que el jugador sigue vivo
		lista_jugadores[ID_rpta-1].suma_rptas_etapa1 += rpta
		if lista_jugadores[ID_rpta-1].suma_rptas_etapa1 >= 21 { // Si la suma de sus respuestas es >= 21, informar que pasa a la siguiente etapa
			// Informar que pasa a la siguiente etapa
			pasan_etapa = append(pasan_etapa, ID_rpta) // Como ya paso, no se calcula en los que faltan
		}

		// eliminar jugadores que en la 4ta ronda no alcacen 21
		if (contador_rondas == 3 && lista_jugadores[ID_rpta-1].suma_rptas_etapa1 < 21 ){
			// Jugador eliminado
			//esEliminado(*lista_jugadores[ID_rpta-1])
			esEliminado(ID_rpta-1)
			muertos_por_ronda = append(muertos_por_ronda, ID_rpta) // Añadido a la lista de muertos
		}
	}

	container.inc("cont_req")
	// Registrar su jugada en NameNode
	registrarJugada(ID_rpta, "1", rpta, true)
	
	return &pb.Message{Body: "OK"}, nil 
}

// funcion response procesar jugada de jugador en etapa 2
func (s *Server) ProcesarJugadaDos(ctx context.Context, in *pb.Jugada) (*pb.Message, error) {
	log.Printf("ID de jugador: %d, jugada:  %s", in.ID, in.Numero)
	
	// Guardar el ID de esa respuesta
	ID_rpta := in.ID
	rpta, err := strconv.Atoi(in.Numero)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// aqui se recoje las sumas
	if lista_jugadores[ID_rpta-1].equipo_etapa2 == "A" {
		grupoA_suma += rpta
	} else if lista_jugadores[ID_rpta-1].equipo_etapa2 == "B" {
		grupoB_suma += rpta
	} else {
		// Algo paso que entro un muerto
		fmt.Println("Recibio a alguien que no tenia un grupo asignado")
	}
	
	container.inc("cont_req")
	// Registrar su jugada en NameNode
	registrarJugada(ID_rpta, "2", rpta, false)

	return &pb.Message{Body: "OK"}, nil 
}

// funcion response procesar jugada de jugador en etapa 3
func (s *Server) ProcesarJugadaTres(ctx context.Context, in *pb.Jugada) (*pb.Message, error) {
	log.Printf("ID de jugador: %d, jugada:  %s", in.ID, in.Numero)
	// Guardar el ID de esa respuesta
	ID_rpta := in.ID
	rpta, err := strconv.Atoi(in.Numero)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	lista_jugadores[ID_rpta-1].rpta_etapa3 = rpta

	container.inc("cont_req")
	// Registrar su jugada en NameNode
	registrarJugada(ID_rpta, "3", rpta, false)

	return &pb.Message{Body: "OK"}, nil 
}

// funcion response para resultado de ronda (etapa 1)
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
			//fmt.Println("El jugador", id_jugador, "es eliminado")
			lista_jugadores[id_jugador-1].estado = "muerto"
			vivo_bool = false
			aumentarPozo(lista_jugadores[index].ID, etapa_actual)
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
		container.inc("cont_res")		
		return &pb.Message{Body: "GANADOR"}, nil
	}

	container.inc("cont_res")		
	return &pb.Message{Body: "OK"}, nil
}

// funcion response para notificar si estado 
func  (s *Server) NotificarEstado (ctx context.Context, in *pb.RespuestaSolicitud) (*pb.Message, error) {
	if (flag_notificacion == false){
		return &pb.Message{Body: "NOK"}, nil
	}

	for i := 0; i < len(por_eliminar); i++  {
		//fmt.Println(por_eliminar[i], in.ID)
		if por_eliminar[i] == in.ID {
			container.inc("cont_res")
			return &pb.Message{Body: "ELIMINADO"}, nil // eliminar jugador
		}
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

func getRandomNum(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min+1) + min
	return result
}

func jugadoresVivos(cant_jugadores int, print bool) (lista_vivos []int32) { // Muestra por consola a los jugadores vivos al término de cada etapa
	
	for i := 0; i < cant_jugadores; i++ {
		//fmt.Println("-> ", lista_jugadores[i].ID, lista_jugadores[i].estado)
		if lista_jugadores[i].estado == "vivo" {
			if print {
				fmt.Println("El jugador:", lista_jugadores[i].ID, "está vivo")
			}
			lista_vivos = append(lista_vivos, lista_jugadores[i].ID)
		}
	}
	return lista_vivos
}

func mostrarGanadores(cant_jugadores int) { // Muestra por consola a los ganadores y wones ganados
	// TO DO: Añadir la logica de mostrar los wones
	for i := 0; i < cant_jugadores; i++ {
		if lista_jugadores[i].estado == "vivo" {
			fmt.Println("El jugador:", lista_jugadores[i].ID, "ha ganado el juego del calamar")
		}
	}
}


func esEliminado(index int32) { // Cuando la logica detecta que el jugador muere, avisa al jugador
	lista_jugadores[index].estado = "muerto"
	por_eliminar = append(por_eliminar, lista_jugadores[index].ID)
	aumentarPozo(lista_jugadores[index].ID, etapa_actual)
	mostrarMuerte(lista_jugadores[index].ID) // Printea por consola quien murio
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

	} else if num_vivos == 1 { 
		// Notificar ganador
		fmt.Println("Solo queda un jugador en pie")
		fmt.Println("El jugador ", lista_vivos[0], " gana el juego del Calamar")
		return 2
	} else {
		// Habilitar que jugadores que pasan a la fase 2
		fmt.Print("Los jugadores de ID: ")
		for i := 0; i < num_vivos; i++ {
			fmt.Print(lista_vivos[i], " ")
		}
		fmt.Println(" pasan a la siguiente etapa")
		return 1
	}
}

func RemoveIndex(s []int32, index int32) []int32 { // Elimina la posicion index del slice, codigo de stackoverflow
	//fmt.Println(index)
	//fmt.Println(s)

	if (len(s) == 0){ // slice vacio
		return s
	}
	ret := make([]int32, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func eliminarJugadorImpar(lista_jugadores []*jugador, lista_vivos []int32) (new_lista_vivos []int32) {
	index_eliminado := getRandomNum(0, len(lista_vivos)-1)
	//fmt.Println(index_eliminado)
	//index_eliminado = 3
	ID_eliminado := lista_vivos[index_eliminado]
	//esEliminado(*lista_jugadores[ID_eliminado-1])
	esEliminado(ID_eliminado-1)
	//new_lista_vivos = RemoveIndex(lista_vivos, int32(index_eliminado) ) // Elimina de la lista de vivos al jugador del index que se debe eliminar
	new_lista_vivos = jugadoresVivos(cant_jugadores, false)
	return new_lista_vivos // slice con el jugador eliminado
}

func asignarGrupos(lista_jugadores []*jugador, lista_vivos []int32) {
	for i := 0; i < len(lista_vivos); i++ {
		currentID := lista_vivos[i]
		if i < len(lista_vivos)/2 {
			fmt.Println("Jugador ID: ", currentID, " asignado a grupo A")
			lista_jugadores[currentID-1].equipo_etapa2 = "A"
		} else {
			lista_jugadores[currentID-1].equipo_etapa2 = "B"
			fmt.Println("Jugador ID: ", currentID, " asignado a grupo B")
		}
	}
}

func eliminarGrupo(lista_jugadores []*jugador, lista_vivos []int32, letra string) (new_lista_vivos []int32) {
	//fmt.Println("Antes: ", lista_vivos)
	for i := 0; i < len(lista_vivos); i++ {
		currentID := lista_vivos[i]
		//fmt.Println(lista_jugadores[currentID-1].equipo_etapa2, letra)
		if lista_jugadores[currentID-1].equipo_etapa2 == letra {
			fmt.Println("Procesando eliminacion")
			// Si el jugador es del grupo 'letra', se elimina
			//esEliminado(*lista_jugadores[currentID-1])
			esEliminado(currentID-1)
			//lista_vivos = RemoveIndex(lista_vivos, int32(i))
		}
	}

	new_lista_vivos = jugadoresVivos(cant_jugadores, false)
	//fmt.Println("Despues: ", new_lista_vivos)

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

func consultarNameNode(){
	fmt.Println("¿Consultar jugadas de jugador? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	eleccion:= ""
	if (char == 'y'){
		fmt.Println("Elija el ID de jugador: ")
		reader := bufio.NewReader(os.Stdin)
		eleccion, _ = reader.ReadString('\n')
		eleccion = strings.TrimSuffix(eleccion, "\n")

		eleccion_int, err := strconv.Atoi(eleccion)
		
		// conectar con NameNode	
		var conn *grpc.ClientConn
		//conn, err = grpc.Dial(":9400", grpc.WithInsecure())
		conn, err = grpc.Dial("10.6.43.104:9400", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c_lider := pb.NewNameNodeServiceClient(conn)
		
		response, err := c_lider.ConsultarJugada(context.Background(), &pb.RespuestaSolicitud{ID: int32(eleccion_int) })
		if err != nil {
			log.Fatalf("Error when calling ConsultarJugada: %s", err)
		}

		fmt.Println(response.Body)

	}
}

/******* RABBITMQ *************************/
func aumentarPozo(ID int32, etapa string){ 
	fmt.Println("Aumentando pozo (asíncronamente)")
	// SETUP RABBITMQ
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@10.6.43.102:5672/")
	
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"pozo", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := "Hello World!" // MENSAJE A ENVIAR ---------------
	body := strconv.Itoa(int(in.ID)) + "_" + etapa

	err = ch.Publish( 	   // ENVIA EL MENSAJE A LA COLA
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

/******* variables globales ***************/
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
var por_eliminar []int32 // lista de quienes hay que notificar su eliminacion (etapas 2 y 3)

var opt_lider int // opcion del lider de etapa 1
var en_juego int // jugadores en juego
var flag_next_etapa bool  // flag para empezar siguiente etapa
var flag_notificacion bool // flag para notificar estado jugadores
var notificar_ganador bool // flag para notificar ganador

var contador_rondas int // contador de ronda para etapa 1

var grupoA_suma int // sumas para etapa 2
var grupoB_suma  int 

var etapa_actual int

/************** Funcion main ***************/
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
	flag_notificacion = false
	var etapa = 1
	const max_rondas = 4
	contador_rondas = 0
	en_juego = jugadores_vivos

	// esperar notificar a todos los jugadores
	flag_next_etapa = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < cant_jugadores{
		//fmt.Println(container.counters["cont_res"], cant_jugadores)
	}
	flag_next_etapa = false

	/******************** LOOP ETAPA 1 ********************************/
	fmt.Println("Inicio de Etapa 1")
	etapa_actual = "1"
	for contador_rondas < max_rondas {
		
		// Lider escoge un numero al azar entre el 6 y 10
		opt_lider = getRandomNum(6, 10)
		//opt_lider = 22
		
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
	var lista_vivos = jugadoresVivos(cant_jugadores, false)
	resultado_etapa := evaluarRestantes(lista_vivos, lista_jugadores) // Evalua si quedan jugadores suficientes para jugar
	if resultado_etapa == 0 { // no hay jugadores en pie
		fmt.Println("Terminando juego del calamar sin ganadores.")
		consultarNameNode()
		fmt.Println("Cerrando proceso lider.")
		return
	}else if resultado_etapa == 2{ // un solo jugador en pie, el ganador 
		fmt.Println("Notificando ganador...")
		notificar_ganador = true
		flag_next_etapa = true
		// esperar a que se notifique al ganador
		container.reset("cont_res")		
		for container.counters["cont_res"] < 1 {
		}
		consultarNameNode()
		fmt.Println("Cerrando proceso lider.")
		return 
	}

	
	consultarNameNode()
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
	
	// esperar notificar a todos que sigue la etapa
	flag_next_etapa = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < len(lista_vivos){
	}
	flag_next_etapa = false

	etapa_actual = "2"
	/*** Procesamiento etapa 2 ***/
	jugadores_a_notificar := jugadores_vivos
	por_eliminar = nil

	// Ver si los restantes son impar
	if len(lista_vivos)%2 == 1 {
		fmt.Println("Sobrevivientes impares, se eliminará a uno")
		lista_vivos = eliminarJugadorImpar(lista_jugadores, lista_vivos) // Se elimina uno al azar y retorna la nueva lista de vivos
		jugadores_vivos -= 1
	}
	// esperar notificar a todos los jugadores vivos
	flag_notificacion = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < jugadores_a_notificar{
	}
	flag_notificacion = false

	/*** Inicio de etapa 2 ***/
	fmt.Println("Inicio de Etapa 2")
	
	etapa = etapa + 1

	asignarGrupos(lista_jugadores, lista_vivos) // La mitad de los vivos se van al grupo A, la otra mitad al B (en enunciado no dice como escogerlos)
	grupoA_suma = 0
	grupoB_suma = 0
	
	var len_grupo = jugadores_vivos / 2
	var paridad_A = false
	var paridad_B = false
	// Lider escoge un numero al azar entre el 1 y 4
	opt_lider = getRandomNum(1, 4)
	fmt.Println("La opcion del Lider es:", opt_lider)

	// esperar a que todos los vivos jueguen		
	container.reset("cont_req")
	for container.counters["cont_req"] < jugadores_vivos{
	}

	if opt_lider%2 == grupoA_suma%2 { // Si tiene la misma paridad, se activa el flag
		// Activa flag grupo A
		paridad_A = true
	}
	if opt_lider%2 == grupoB_suma%2 {
		// Activa flag grupo B
		paridad_B = true
	}
	jugadores_a_notificar = jugadores_vivos
	por_eliminar = nil

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
	flag_notificacion = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < jugadores_a_notificar{
	}
	flag_notificacion = false

	// A estas alturas se elimino un grupo, y si ambos obtienen la misma paridad no se elimino nada ni se restaron los jugadores_vivos
	// TO-DO: lo de pozo de wones

	// Chequear si hay jugadores vivos para seguir jugando, y quienes
	lista_vivos = jugadoresVivos( cant_jugadores, false)
	resultado_etapa = evaluarRestantes(lista_vivos, lista_jugadores) // Evalua si quedan jugadores suficientes para jugar
	if resultado_etapa == 0 { // no hay jugadores en pie
		fmt.Println("Terminando juego del calamar sin ganadores.")
		fmt.Println("Cerrando proceso lider.")
		consultarNameNode()
		return
	}else if resultado_etapa == 2{ // un solo jugador en pie, el ganador 
		fmt.Println("Notificando ganador...")
		notificar_ganador = true
		// esperar a que se notifique al ganador
		container.reset("cont_res")		
		for container.counters["cont_res"] < 1 {
		}
		consultarNameNode()
		fmt.Println("Cerrando proceso lider.")
		return 
	}

	consultarNameNode()
	fmt.Println("¿Continuar a la siguiente etapa? (y/n)")
	reader = bufio.NewReader(os.Stdin)
	char, _, err = reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	if (char != 'y'){
		fmt.Println("Terminando juego del calamar inconcluso.")
		return
	}
	
	// esperar notificar a todos que sigue la etapa
	flag_next_etapa = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < len(lista_vivos){
	}
	flag_next_etapa = false

	etapa_actual = "3"
	/*** Procesamiento etapa 3 ***/
	jugadores_a_notificar = jugadores_vivos
	lista_vivos = jugadoresVivos(cant_jugadores, false)

	// Ver si los restantes son impar
	if len(lista_vivos)%2 == 1 {
		fmt.Println("Sobrevivientes impares, se eliminará a uno")
		lista_vivos = eliminarJugadorImpar(lista_jugadores, lista_vivos) // Se elimina uno al azar y retorna la nueva lista de vivos
		jugadores_vivos -= 1
	}
	// esperar notificar a todos los jugadores vivos
	flag_notificacion = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < jugadores_a_notificar{
	}
	flag_notificacion = false


	// Inicio de etapa 3
	// Avisa a los jugadores que iniciara la etapa 3
	fmt.Println("Inicio de Etapa 3")
	etapa = etapa + 1

	// Elige parejas
	asigna_parejas := 1 // Las parejas van del 1 hasta el jugadores_vivos/2
	for i := 0; i < jugadores_vivos; i++ {
		ID_player1 := lista_vivos[i]
		i++
		ID_player2 := lista_vivos[i]
		// Quedan las parejas asignadas
		lista_jugadores[ID_player1-1].nro_pareja_etapa3 = asigna_parejas
		lista_jugadores[ID_player2-1].nro_pareja_etapa3 = asigna_parejas

		asigna_parejas++
	}

	// esperar a que todos los vivos jueguen		
	container.reset("cont_req")
	for container.counters["cont_req"] < jugadores_vivos{
	}

	// Lider escoge un numero random entre 1 y 10
	opt_lider = getRandomNum(1, 10)
	fmt.Println("La opcion del Lider es:", opt_lider)

	lista_vivos = jugadoresVivos( cant_jugadores, false)

	jugadores_a_notificar = jugadores_vivos
	por_eliminar = nil

	// Procesa sus respuestas, compara resultados entre si y con el lider
	for nro_pareja := 1; nro_pareja <= jugadores_vivos/2; nro_pareja++ {
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
			//esEliminado(*lista_jugadores[vs_pareja[0]-1])
			esEliminado(vs_pareja[0]-1)
		} else if rpta1 < rpta2 { // Pierde rpta2
			// Avisar al jugador con ID vs_pareja[1] que perdio
			//esEliminado(*lista_jugadores[vs_pareja[1]-1])
			esEliminado(vs_pareja[1]-1)
		} // Si no entra a ninguno de los dos es que ambos ganaron, no se elimina a nadie
	}

	flag_notificacion = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < jugadores_a_notificar{
	}
	flag_notificacion = false

	fmt.Println("Notificando ganador...")
	lista_vivos = jugadoresVivos( cant_jugadores, false)
	// notificar ganadores
	flag_next_etapa = true
	notificar_ganador = true
	container.reset("cont_res")		
	for container.counters["cont_res"] < len(lista_vivos){
	}
	flag_notificacion = false
	
	// Entrega ganadores y la cantidad de wones
	mostrarGanadores(cant_jugadores) 
	
	consultarNameNode()
	fmt.Println("Proceso lider finalizado.")
	return

}
