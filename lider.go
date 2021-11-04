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

func jugadoresVivos() { // Muestra por consola a los jugadores vivos al término de cada etapa

}

func mostrarGanadores() { // Muestra por consola a los ganadores

}

func comenzarEtapa() { // Comienza la siguiente etapa

}

func revisarJugadas() { // Para preguntar sobre las jugadas historicamente de un determinado jugador

}

func esEliminado() { // Cuando la logica detecta que el jugador muere, avisa al jugador y al pozo

}

func main() {
	// Definiciones iniciales
	const cant_jugadores = 16
	const wones = 100000000                      // 100 millones de wones vale cada jugador
	var lista_jugadores [cant_jugadores]*jugador // Lista de structs con los jugadores
	var muertos_por_ronda [cant_jugadores]int    // Una lista de muertos ([1,4,5,8,16] por ejemplo)
	var jugadores_vivos = cant_jugadores

	// Setup inicial recibiendo los jugadores
	for i := 0; i < 16; i++ { // Tal vez cambiarlo a un while, dependiendo de la mecanica de esperar

		// Aqui espera a una conexion y cuando la recibe la asigna
		// recibirPeticion()
		// Guardar la IP y puerto recibidos

		lista_jugadores[i] = newJugador(i+1, "ip", "puerto") // Jugador 1, 2, 3, ..., 16

		// Aqui responde al jugador con el ID asignado

	}

	// Inicio de etapa 1
	// Primero aqui avisa a los jugadores que iniciara la etapa 1 y que manden sus respuestas

	// Lider escoge un numero al azar entre el 6 y 10
	fmt.Println(getRandomNum(5, 10))
}
