package main
import (
    "fmt"
    "os"
)
func main() {
	argsWithoutProg := os.Args[1:]
	// checker si es bot o no
	var bot = false
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "bot"{
		bot = true
	}
	
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

}
