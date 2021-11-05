package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
func plus(a int, b int) int {

	return a + b
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
*/

type jugador struct {
	ID                int
	IP                string
	puerto            string
	estado            string
	suma_rptas_etapa1 int
	equipo_etapa2     string
	// pareja??
}

func newJugador(ID int, IP string, puerto string) *jugador { // Crea un struct jugador con ID, IP, puerto, etc
	player := jugador{
		ID:                ID,
		IP:                IP,
		puerto:            puerto,
		estado:            "vivo",
		suma_rptas_etapa1: 0}
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

func jugadoresVivos(lista_jugadores []*jugador, cant_jugadores int, print bool) (lista_vivos []int) { // Muestra por consola a los jugadores vivos al término de cada etapa
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

func mostrarGanadores() { // Muestra por consola a los ganadores

}

func revisarJugadas() { // Para preguntar sobre las jugadas historicamente de un determinado jugador

}

func esEliminado(jugador_eliminado jugador) { // Cuando la logica detecta que el jugador muere, avisa al jugador
	jugador_eliminado.estado = "muerto"
	mostrarMuerte(jugador_eliminado.ID) // Printea por consola quien murio
}

func mostrarMuerte(ID_player int) { // Muestra por consola cuando muere un jugador
	fmt.Println("El jugador:", ID_player, "ha sido eliminado")
}

func informarResultadoRonda(muertos_por_ronda []int, pasan_etapa []int) { // Lee los ID de quienes murieron/pasaron y les manda mensaje, tambien al pozo
	for i := 0; i < len(muertos_por_ronda); i++ {
		playerID := muertos_por_ronda[i]
		// Mandar msj a jugador de ID playerID
		// Mensaje de ejemplo:
		fmt.Println("El jugador", playerID, "(tú) está eliminado")
		// Mandar msj al pozo para aumentarlo
	}
	for i := 0; i < len(pasan_etapa); i++ {
		playerID := pasan_etapa[i]
		// Mandar msj a jugador de ID playerID
		// Mensajes de ejemplo
		fmt.Println("El jugador", playerID, "(tú) pasa a la siguiente etapa")
		fmt.Println("Esperando término de etapa...")
	}
}

func evaluarRestantes(lista_vivos []int, lista_jugadores []*jugador) { // Evalua si quedan jugadores suficientes para jugar
	num_vivos := len(lista_vivos)
	if num_vivos == 0 {
		// Mandar mensaje a ¿todos?, no hay ganadores
		fmt.Println("Todos los jugadores fueron eliminados")
		fmt.Println("No hay ganadores")
		// Dar la orden para terminar el juego
	} else if num_vivos == 1 {
		// Mandar mensaje avisando quien gano
		fmt.Println("Solo queda un jugador en pie")
		fmt.Println("El jugador", lista_vivos[0], "ganó el juego del Calamar")
		// Dar la orden para terminar el juego
	} else {
		// Mandar mensaje a los jugadores que pasan a la fase 2
		fmt.Println("Los jugadores de ID:")
		for i := 0; i < num_vivos; i++ {
			// Mandar mensaje a los jugadores
			fmt.Println(lista_vivos[i])
		}
		fmt.Println("pasan a la siguiente ronda")
	}
}

func RemoveIndex(s []int, index int) []int { // Elimina la posicion index del slice, codigo de stackoverflow
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func eliminarJugadorImpar(lista_jugadores []*jugador, lista_vivos []int) (new_lista_vivos []int) {
	index_eliminado := getRandomNum(0, len(lista_vivos))
	ID_eliminado := lista_vivos[index_eliminado]
	esEliminado(*lista_jugadores[ID_eliminado])
	new_lista_vivos = RemoveIndex(lista_vivos, index_eliminado) // Elimina de la lista de vivos al jugador del index que se debe eliminar
	return new_lista_vivos
}

func asignarGrupos(lista_jugadores []*jugador, lista_vivos []int) {
	for i := 0; i < len(lista_vivos); i++ {
		currentID := lista_vivos[i]
		if i < len(lista_vivos)/2 {
			// Mandar mensaje a los jugadores con su equipo asignado
			lista_jugadores[currentID].equipo_etapa2 = "A"
		} else {
			// Mandar mensaje a los jugadores con su equipo asignado
			lista_jugadores[currentID].equipo_etapa2 = "B"
		}
	}
}

func main() {
	// Definiciones iniciales

	cant_jugadores := 16
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		cant_jugadores, err := strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cant_jugadores)
	}

	const wones = 100000000 // 100 millones de wones vale cada jugador
	//var lista_jugadores [cant_jugadores]*jugador // Deprecado: Lista de structs con los jugadores
	var lista_jugadores []*jugador // Slice de structs con los jugadores

	var jugadores_vivos int = cant_jugadores

	// Setup inicial recibiendo los jugadores
	for i := 0; i < cant_jugadores; i++ {

		// Aqui espera a una conexion y cuando la recibe la asigna
		// recibirPeticion()
		// Guardar la IP y puerto recibidos

		//lista_jugadores[i] = newJugador(i+1, "ip", "puerto") // Deprecado.
		lista_jugadores = append(lista_jugadores, newJugador(i+1, "ip", "puerto")) // Jugador 1, 2, 3, ..., 16

		// Aqui responde al jugador con el ID asignado

	}

	// Inicio de etapa 1
	// Primero aqui avisa a los jugadores que iniciara la etapa 1 y que manden sus respuestas

	var etapa = 1
	const max_rondas = 4
	var contador_rondas = 0
	var en_juego = jugadores_vivos

	// LOOP ETAPA 1
	// -------------
	for contador_rondas < max_rondas {

		var muertos_por_ronda []int // Una lista de muertos ([1,4,5,8,16] por ejemplo)
		var pasan_etapa []int       // Una lista de quienes pasan la etapa, por ronda
		// Lider escoge un numero al azar entre el 6 y 10
		opt_lider := getRandomNum(6, 10)
		fmt.Println("La opcion del Lider es:", opt_lider)

		for i := 0; i < en_juego; i++ {

			// Recoger primera respuesta que llegue
			// Guardar el ID de esa respuesta

			// Hardcodeo
			ID_rpta := 1
			rpta := 4

			if rpta >= opt_lider {
				// Jugador eliminado
				esEliminado(*lista_jugadores[ID_rpta])
				muertos_por_ronda = append(muertos_por_ronda, ID_rpta) // Añadido a la lista de muertos
			} else {
				// Mandar mensaje que el jugador sigue vivo
				// Mostrar el numero que lleva acumulado
				lista_jugadores[ID_rpta-1].suma_rptas_etapa1 += rpta
				if lista_jugadores[ID_rpta-1].suma_rptas_etapa1 >= 21 { // Si la suma de sus respuestas es >= 21, informar que pasa a la siguiente etapa
					// Informar que pasa a la siguiente etapa
					pasan_etapa = append(pasan_etapa, ID_rpta) // Como ya paso, no se calcula en los que faltan
				}
			}
			// Registrar su jugada en NameNode
		}

		jugadores_vivos -= len(muertos_por_ronda)
		en_juego -= len(muertos_por_ronda)
		en_juego -= len(pasan_etapa)
		contador_rondas += 1

		// Leer slices e informar quienes pasan o no
		informarResultadoRonda(muertos_por_ronda, pasan_etapa)
	}
	// Chequear si hay jugadores vivos para seguir jugando, y quienes
	var lista_vivos = jugadoresVivos(lista_jugadores, cant_jugadores, false)
	evaluarRestantes(lista_vivos, lista_jugadores) // Evalua si quedan jugadores suficientes para jugar

	// Ver si los restantes son impar
	if len(lista_vivos)%2 == 1 {
		fmt.Println("Sobrevivientes impares, se eliminará a uno")
		lista_vivos = eliminarJugadorImpar(lista_jugadores, lista_vivos) // Se elimina uno al azar y retorna la nueva lista de vivos
	}

	// Inicio de etapa 2
	// Primero aqui avisa a los jugadores que iniciara la etapa 2 y que manden sus respuestas
	etapa = etapa + 1

	asignarGrupos(lista_jugadores, lista_vivos)
	//var grupoA_suma = 0
	//var grupoB_suma = 0
}
