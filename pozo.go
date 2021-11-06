package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Recibe un msj con el ID de un jugador muerto y la etapa en la que fue eliminado
	// Este se debe enviar en la funcion esEliminado de lider
	// Debe quedarse escuchando e ir escribiendo en el txt

	cant_jugadores := 16 // Esto se asume pero para testear CAMBIARLOOOOOOOOOOOO
	pozo := 0
	const wones = 100000000

	f, _ := os.Create("pozo.txt")

	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	for i := 0; i < cant_jugadores; i++ {
		// Espera 16 mensajes, o en realidad deberian ser 15?
		// o libre? (puede haber mas de un ganador)
		// AAAAA modificar en caso de

		// HARDCODEO, obtener esto con mensajes
		ID_actual := 1
		etapa := 1

		// Propuesta de "paralelismo" (agregar lo de sincrono/asincrono con rabbitMQ)
		if ID_actual == 0 { // Mandar este ID para decir "estoy consultando cuantos wones hay"
			// Mandar msj con los wones actuales
			fmt.Println("El pozo actual acumulado es de:", pozo) // <-- mandar esto
		} else {
			// Funcionamiento normal de agregar al pozo
			pozo += wones
			write_str := "Jugador_" + strconv.Itoa(ID_actual) + " Etapa_" + strconv.Itoa(etapa) + " " + strconv.Itoa(pozo)
			f.WriteString(write_str + "\n")
		}

	}

}
