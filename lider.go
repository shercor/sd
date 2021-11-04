package main

//import "fmt" // Para printear

/*
func plus(a int, b int) int {

	return a + b
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
*/

func recibirPeticion() { // Recibe peticion para jugar, viene desde un jugador

}

func recibirJugada() { // Recibe la jugada de un jugador en una cierta etapa

}

func informarJugadas() { // Informa a NameNode cada vez que un jugador hace una jugada

}

func informarPozo() { // Agrega 100 millones a Pozo cuando un jugador muere

}

func setupInicial() { // Setup inicial que recibe el "ok" de los jugadores cuando se están conectados

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
	// Funcion de peticiones de juego
	var lista_jugadores [16]int
	wones := 100000000           // 100 millones de wones vale cada jugador
	var muertos_por_fase [16]int // Una lista de muertos ([1,4,5,8,16] por ejemplo)

	for i := 0; i < 16; i++ {
		lista_jugadores[i] = i + 1 // Jugador 1, 2, 3, ..., 16
		// Esto se inicializa una vez que recibe y acepta 16 peticiones
	}
	setupInicial()
}
