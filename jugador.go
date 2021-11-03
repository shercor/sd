package main
import (
    "fmt"
    "os"
	"bufio"
)
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

	for alive {
		fmt.Println("-----------Vivo----------")

		// TO-DO: logica de juegos 
		
		break // TO-DO: logica de salida
	}
	fmt.Println("Finalizando proceso jugador")
	
	
	/*
	fmt.Println("Hola, mundo.")
	fmt.Println("Bot? ", bot)

	var a string = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
	*/

	

}
