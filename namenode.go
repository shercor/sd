package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomNum(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min+1) + min
	return result
}

func main() {

	// Recibe las jugadas de cada jugador que lo manda al lider.go
	// Recibe el ID_Jugador, la etapa y la jugada (o jugadas en caso de la etapa 1)

	ip_datanodes := [...]string{"10.0.1.11", "10.0.1.12", "10.0.1.13"} // o algo asi

	f, _ := os.Create("namenode.txt")

	defer f.Close() // Cierra el archivo cuando termina la ejecucion

	// Esperar mensajes
	// Cada vez que llega un msj, hace esto:

	// HARDCODEO, obtener esto con mensajes
	ID_actual := 1
	etapa := 1
	ronda := 1
	opt := getRandomNum(0, 2) // Elegir un indice del slice al azar

	// Mandar mensaje a ip_datanodes[opt]
	// El mensaje contiene el ID_jugador, la etapa y la jugada

	// Escribir en un txt
	if etapa == 1 { // Si es la etapa 1 se crea con la ronda ademas
		write_str := "Jugador_" + strconv.Itoa(ID_actual) + " Etapa_" + strconv.Itoa(etapa) + " Ronda_" + strconv.Itoa(ronda) + ip_datanodes[opt]
	} else {
		write_str := "Jugador_" + strconv.Itoa(ID_actual) + " Etapa_" + strconv.Itoa(etapa) + " " + ip_datanodes[opt]
	}
	write_str := "Jugador_" + strconv.Itoa(ID_actual) + " Etapa_" + strconv.Itoa(etapa) + " " + ip_datanodes[opt]
	f.WriteString(write_str + "\n")

}
