package main

import (
	"os"
	"strconv"
)

func main() {

	// Recibe las jugadas de cada jugador, mandadas desde el NameNode
	// Recibe el ID_Jugador, la etapa y la jugada (o jugadas en caso de la etapa 1)

	// Esperar mensajes
	// Cada vez que llega un msj, hace esto:

	// HARDCODEO, obtener esto con mensajes y ponerle loop
	ID_actual := 1
	etapa := 1
	jugada := 5
	ronda := 1 // RONDA SOLO SE MANDA EN ETAPA 1 o de lo contrario mandar siempre con valor 0 (igual no se usará xd)

	nametxt := ""

	if etapa == 1 { // Si es la etapa 1, se creara un archivo para cada ronda, jugador_5__etapa_1__ronda_2.txt por ejemplo
		nametxt = "Jugador_" + strconv.Itoa(ID_actual) + "__Etapa_" + strconv.Itoa(etapa) + "__Ronda_" + strconv.Itoa(ronda) + ".txt"
	} else {
		nametxt = "Jugador_" + strconv.Itoa(ID_actual) + "__Etapa_" + strconv.Itoa(etapa) + ".txt"
	}

	f, _ := os.Create(nametxt)

	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	// Escribir en un txt
	write_str := strconv.Itoa(jugada)
	f.WriteString(write_str + "\n")

	// Falta implementar el caso para la etapa 1, que es la unica multirondas

}
