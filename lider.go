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
	ID     int
	IP     string
	puerto string
	estado string
	// equipo?
}

func newJugador(ID int, IP string, puerto string) *jugador { // Crea un struct jugador con ID, IP, puerto, etc
	player := jugador{ID: ID, IP: IP, puerto: puerto, estado: "vivo"}
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

func llamarInterfaz() { // Llama a la interfaz principal para iniciar el juego y demases

}

func mostrarMuerte() { // Muestra por consola cuando muere un jugador

}

func jugadoresVivos(lista_jugadores []jugador, cant_jugadores int) { // Muestra por consola a los jugadores vivos al término de cada etapa
	for i := 0; i < cant_jugadores; i++ {
		if lista_jugadores[i].estado == "vivo" {
			fmt.Println("El jugador:", lista_jugadores[i].ID, "está vivo")
		}
	}
}

func mostrarGanadores() { // Muestra por consola a los ganadores

}

func revisarJugadas() { // Para preguntar sobre las jugadas historicamente de un determinado jugador

}

func esEliminado(jugador_eliminado jugador) { // Cuando la logica detecta que el jugador muere, avisa al jugador y al pozo
	// Enviar msj al jugador diciendo que murio
	// Enviar msj al pozo para aumentarlo
	jugador_eliminado.estado = "muerto"
}

func main() {
	// Definiciones iniciales

	cant_jugadores := 16
	argsWithoutProg := os.Args[1:]
	
	if len(argsWithoutProg) ==  1 {
		cant_jugadores, err :=  strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cant_jugadores)
	}

	const wones = 100000000                      // 100 millones de wones vale cada jugador
	//var lista_jugadores [cant_jugadores]*jugador // Deprecado: Lista de structs con los jugadores
	var lista_jugadores []*jugador // Slice de structs con los jugadores

	var jugadores_vivos int = cant_jugadores

	// Setup inicial recibiendo los jugadores
	for i := 0; i < cant_jugadores; i++ {

		// Aqui espera a una conexion y cuando la recibe la asigna
		// recibirPeticion()
		// Guardar la IP y puerto recibidos

		//lista_jugadores[i] = newJugador(i+1, "ip", "puerto") // Deprecado.
		lista_jugadores = append(lista_jugadores, newJugador(i+1, "ip", "puerto") ) // Jugador 1, 2, 3, ..., 16

		// Aqui responde al jugador con el ID asignado

	}

	// Inicio de etapa 1
	// Primero aqui avisa a los jugadores que iniciara la etapa 1 y que manden sus respuestas

	var etapa = 1
	const max_rondas = 4
	var contador_rondas int = 0
	// LOOP ETAPA 1
	// -------------
	for contador_rondas < max_rondas {

		var muertos_por_ronda []int // Una lista de muertos ([1,4,5,8,16] por ejemplo)
		// Lider escoge un numero al azar entre el 6 y 10
		opt_lider := getRandomNum(6, 10)
		fmt.Println("La opcion del Lider es:", opt_lider)

		for i := 0; i < jugadores_vivos; i++ {

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
				// Si la suma de sus respuestas es >= 21, informar que pasa a la siguiente etapa

				// -------------------------------------------
				// TO-DO
				// VER LO DE LA SUMA DE RESPUESTAS, AGREGAR UN CAMPO AL STRUCT
				// -------------------------------------------
			}
			// Registrar su jugada en NameNode
		}

		jugadores_vivos = jugadores_vivos - len(muertos_por_ronda)
		contador_rondas = contador_rondas + 1
	}

	etapa = etapa + 1
}
